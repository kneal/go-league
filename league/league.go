package league

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"reflect"
	"strings"

	"github.com/google/go-querystring/query"
)

const (
	userAgent = "go-league"
)

// Client is a client that manages communication with the LOL API.
type (
	Client struct {
		// HTTP client used to communicate with the LOL API.
		client *http.Client

		// Base URL for LOL API requests.
		baseURL *url.URL

		// User agent used when communicating with the LOL API.
		UserAgent string

		// LOL service for authentication.
		Authentication *AuthenticationService
		Champion       *ChampionService
		Summoner       *SummonerService
		Match          *MatchService
	}

	service struct {
		client *Client
	}
)

// NewClient returns a new LOL API client.
// baseURL has to be the HTTP endpoint of the LOL API.
// If no httpClient is provided, then the http.DefaultClient will be used.
func NewClient(baseURL string, httpClient *http.Client) (*Client, error) {
	// use http.DefaultClient if no client is provided
	if httpClient == nil {
		httpClient = http.DefaultClient
	}

	// we must have a url provided to create the client
	if len(baseURL) == 0 {
		return nil, fmt.Errorf("No LOL baseURL provided")
	}

	// parse url provided for the client
	url, err := url.Parse(baseURL)
	if err != nil {
		return nil, err
	}

	// create initial client fields
	c := &Client{
		client:    httpClient,
		baseURL:   url,
		UserAgent: userAgent,
	}

	// instantiate all client services
	c.Authentication = &AuthenticationService{client: c}
	c.Champion = &ChampionService{client: c}
	c.Summoner = &SummonerService{client: c}
	c.Match = &MatchService{client: c}

	return c, nil
}

// buildURLForRequest will build the URL (as a string) that will be called.
// It does several cleaning tasks for us.
func (c *Client) buildURLForRequest(urlStr string) (string, error) {
	// capture base url from client for string
	u := c.baseURL.String()

	// If there is no / at the end, add one.
	if strings.HasSuffix(u, "/") == false {
		u += "/"
	}

	// If there is a "/" at the start, remove it.
	if strings.HasPrefix(urlStr, "/") == true {
		urlStr = urlStr[1:]
	}

	// parse trimmed url string
	rel, err := url.Parse(urlStr)
	if err != nil {
		return "", err
	}
	u += rel.String()

	return u, nil
}

// addAuthentication adds the necessary authentication to the request.
func (c *Client) addAuthentication(req *http.Request) {
	// Apply Token Authentication.
	if c.Authentication.HasTokenAuth() {
		req.Header.Add("X-Riot-Token", fmt.Sprintf("%s", *c.Authentication.secret))
	}
}

// addOptions adds the parameters in opt as url query parameters to s.
// opt must be a struct whose fields may contain "url" tags.
func addOptions(s string, opt interface{}) (string, error) {
	// return url if option is a pointer but is also nil
	v := reflect.ValueOf(opt)
	if v.Kind() == reflect.Ptr && v.IsNil() {
		return s, nil
	}

	// parse url provided for the options
	u, err := url.Parse(s)
	if err != nil {
		return s, err
	}

	// add query values to url
	qs, err := query.Values(opt)
	if err != nil {
		return s, err
	}

	// safely encode url with query values
	u.RawQuery = qs.Encode()
	return u.String(), nil
}

// NewRequest creates an API request.
// A relative URL can be provided in urlStr,
// in which case it is resolved relative to the baseURL of the Client.
// Relative URLs should always be specified without a preceding slash.
// If specified, the value pointed to by body is JSON encoded and included as the request body.
func (c *Client) NewRequest(method, urlStr string, body interface{}) (*http.Request, error) {
	// build url for request
	u, err := c.buildURLForRequest(urlStr)
	if err != nil {
		return nil, err
	}

	var buf io.ReadWriter
	if body != nil {
		// buffer to store request body
		buf = new(bytes.Buffer)

		// encode request body into buffer for request
		err := json.NewEncoder(buf).Encode(body)
		if err != nil {
			return nil, err
		}
	}

	// create new http request from built url and body
	req, err := http.NewRequest(method, u, buf)
	if err != nil {
		return nil, err
	}

	// apply authentication to request if client is set
	if c.Authentication.HasAuth() {
		c.addAuthentication(req)
	}

	// apply default header for request
	req.Header.Add("Content-Type", "application/json")

	return req, nil
}

// Response represents an LOL API response.
// This wraps the standard http.Response returned from LOL.
type Response struct {
	*http.Response
}

// Call is a combine function for Client.NewRequest and Client.Do.
//
// Most API methods are quite the same.
// Get the URL, apply options, make a request, and get the response.
// Without adding special headers or something.
// To avoid a big amount of code duplication you can Client.Call.
//
// method is the HTTP method you want to call.
// u is the URL you want to call.
// body is the HTTP body.
// v is the HTTP response.
//
// For more information read https://github.com/google/go-github/issues/234
func (c *Client) Call(method, u string, body interface{}, v interface{}) (*Response, error) {
	// create new request from parameters
	req, err := c.NewRequest(method, u, body)
	if err != nil {
		return nil, err
	}

	// send request with client
	resp, err := c.Do(req, v)
	if err != nil {
		return resp, err
	}

	return resp, err
}

// Do sends an API request and returns the API response.
// The API response is JSON decoded and stored in the value pointed to by v,
// or returned as an error if an API error has occurred.
// If v implements the io.Writer interface, the raw response body will be written to v,
// without attempting to first decode it.
func (c *Client) Do(req *http.Request, v interface{}) (*Response, error) {
	// send request with client
	resp, err := c.client.Do(req)
	if err != nil {
		return nil, err
	}
	// defer closing response body
	defer resp.Body.Close()

	// wrap response
	response := &Response{Response: resp}

	// check response for errors
	err = CheckResponse(resp)
	if err != nil {
		// if error is present, we still return the response so the caller
		// may inspect it further for debugging and troubleshooting
		return response, err
	}

	// if return object is provided
	if v != nil {
		// copy response body if object implements io.Writer interface
		if w, ok := v.(io.Writer); ok {
			io.Copy(w, resp.Body)
		} else {
			// copy all bytes from response body
			body, err := ioutil.ReadAll(resp.Body)
			// ensure response body is not empty so the user may inspect
			// it further for debugging and troubleshooting
			resp.Body = ioutil.NopCloser(bytes.NewBuffer(body))
			if err != nil {
				// if error is present, we still return the response so the caller
				// may inspect it further for debugging and troubleshooting
				return response, err
			}
			// unmarshal the body to the return object
			_ = json.Unmarshal(body, v)
		}
	}
	return response, err
}

// CheckResponse checks the API response for errors, and returns them if present.
// A response is considered an error if it has a status code outside the 200 range.
func CheckResponse(r *http.Response) error {
	if c := r.StatusCode; 200 <= c && c <= 299 {
		return nil
	}

	return fmt.Errorf("API call to %s failed: %s", r.Request.URL.String(), r.Status)
}

// Bool is a helper routine that allocates a new boolean
// value to store v and returns a pointer to it.
func Bool(v bool) *bool { return &v }

// Bytes is a helper routine that allocates a new byte
// array value to store v and returns a pointer to it.
func Bytes(v []byte) *[]byte { return &v }

// Int is a helper routine that allocates a new integer
// value to store v and returns a pointer to it.
func Int(v int) *int { return &v }

// Int64 is a helper routine that allocates a new 64 bit
// integer value to store v and returns a pointer to it.
func Int64(v int64) *int64 { return &v }

// String is a helper routine that allocates a new string
// value to store v and returns a pointer to it.
func String(v string) *string { return &v }

// Strings is a helper routine that allocates a new string
// array value to store v and returns a pointer to it.
func Strings(v []string) *[]string { return &v }

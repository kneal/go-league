package league

import (
	"fmt"
)

// SummonerService handles manging Summoners.
type (
	SummonerService service

	// Summoner represents Summoner rotations, including free-to-play and low-level free-to-play rotations
	Summoner struct {
		ID            string `json:"id"`
		AccountID     string `json:"accountId"`
		Puuid         string `json:"puuid"`
		Name          string `json:"name"`
		ProfileIconID int    `json:"profileIconId"`
		RevisionDate  int64  `json:"revisionDate"`
		SummonerLevel int    `json:"summonerLevel"`
	}
)

// GetByEncryptedAccountID returns a summoner.
//
// https://developer.riotgames.com/apis#summoner-v4
func (svc *SummonerService) GetByEncryptedAccountID(encryptedAccountID string) (*Summoner, *Response, error) {
	// set the API endpoint path we send the request to
	u := fmt.Sprintf("/lol/summoner/v4/summoners/by-account/%s", encryptedAccountID)

	// slice Summoner type we want to return
	v := new(Summoner)

	// send request using client
	resp, err := svc.client.Call("GET", u, nil, v)

	return v, resp, err
}

// GetByName returns a summoner.
//
// https://developer.riotgames.com/apis#summoner-v4
func (svc *SummonerService) GetByName(name string) (*Summoner, *Response, error) {
	// set the API endpoint path we send the request to
	u := fmt.Sprintf("/lol/summoner/v4/summoners/by-name/%s", name)

	// slice Summoner type we want to return
	v := new(Summoner)

	// send request using client
	resp, err := svc.client.Call("GET", u, nil, v)

	return v, resp, err
}

// GetByPUUID returns a summoner.
//
// https://developer.riotgames.com/apis#summoner-v4
func (svc *SummonerService) GetByPUUID(puuid string) (*Summoner, *Response, error) {
	// set the API endpoint path we send the request to
	u := fmt.Sprintf("/lol/summoner/v4/summoners/by-puuid/%s", puuid)

	// slice Summoner type we want to return
	v := new(Summoner)

	// send request using client
	resp, err := svc.client.Call("GET", u, nil, v)

	return v, resp, err
}

// GetBySummonerID returns a summoner.
//
// https://developer.riotgames.com/apis#summoner-v4
func (svc *SummonerService) GetBySummonerID(summonerID string) (*Summoner, *Response, error) {
	// set the API endpoint path we send the request to
	u := fmt.Sprintf("/lol/summoner/v4/summoners/%s", summonerID)

	// slice Summoner type we want to return
	v := new(Summoner)

	// send request using client
	resp, err := svc.client.Call("GET", u, nil, v)

	return v, resp, err
}

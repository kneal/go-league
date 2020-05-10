package league

import (
	"fmt"
	"log"
	"os"
)

func ExampleChampionService_GetRotations() {
	// Set environment variables before running example
	// export LEAGUE_URL=https://<region>.api.riotgames.com/
	// export LEAGUE_API_KEY=<token_value>
	url := os.Getenv("LEAGUE_URL")
	token := os.Getenv("LEAGUE_API_KEY")

	// Setup league client with token authentication
	l, err := NewClient(url, nil)
	if err != nil {
		log.Fatal("unable to create client")
	}
	// Add the user token to the clients
	l.Authentication.SetTokenAuth(token)

	// get the rotation champions
	champions, resp, err := l.Champion.GetRotations()
	if err != nil {
		log.Fatalf("unable to get champions: %v", err)
	}
	log.Printf("Received response code %d", resp.StatusCode)

	fmt.Printf("%+v \n", champions)
}

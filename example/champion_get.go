package main

import (
	"fmt"
	"log"
	"os"

	"github.com/kneal/go-league/league"
)

func main() {

	// Set environment variables in your terminal before running example
	// export LEAGUE_URL=https://<region>.api.riotgames.com/lol/platform/v3/champion-rotations
	// export LEAGUE_API_KEY=<token_value>
	url := os.Getenv("LEAGUE_URL")
	token := os.Getenv("LEAGUE_API_KEY")

	// Setup league client with token authentication
	l, err := league.NewClient(url, nil)
	if err != nil {
		log.Fatal("unable to create client")
	}
	// Add the user token to the clients
	l.Authentication.SetTokenAuth(token)

	// get the rotation champions
	champions, resp, err := l.Champion.GetAll()
	if err != nil {
		log.Fatal("unable to get champions %s: %w", err)
	}
	log.Printf("Received response code %d", resp.StatusCode)

	fmt.Printf("%+v \n", champions)
}
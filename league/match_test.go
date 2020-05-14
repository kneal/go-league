package league

import (
	"fmt"
	"log"
	"os"
)

func ExampleMatchService_GetMatchByID() {
	// Set environment variables before running example
	// export LEAGUE_URL=https://<region>.api.riotgames.com/
	// export LEAGUE_API_KEY=<token_value>
	url := os.Getenv("LEAGUE_URL")
	token := os.Getenv("LEAGUE_API_KEY")
	matchID := os.Getenv("MATCH_ID")

	// Setup league client with token authentication
	l, err := NewClient(url, nil)
	if err != nil {
		log.Fatal("unable to create client")
	}
	// Add the user token to the clients
	l.Authentication.SetTokenAuth(token)

	match, resp, err := l.Match.GetMatchByID(matchID)
	if err != nil {
		log.Fatalf("unable to get match %s: %v", matchID, err)
	}
	log.Printf("Recieved response code %d", resp.StatusCode)

	fmt.Printf("%+v \n", match)

}

func ExampleMatchService_GetMatchsByAccountID() {
	// Set environment variables before running example
	// export LEAGUE_URL=https://<region>.api.riotgames.com/
	// export LEAGUE_API_KEY=<token_value>
	url := os.Getenv("LEAGUE_URL")
	token := os.Getenv("LEAGUE_API_KEY")
	id := os.Getenv("SUMMONER_ENCRYPTED_ACCOUNT_ID")

	// Setup league client with token authentication
	l, err := NewClient(url, nil)
	if err != nil {
		log.Fatal("unable to create client")
	}
	// Add the user token to the clients
	l.Authentication.SetTokenAuth(token)

	match, resp, err := l.Match.GetMatchsByAccountID(id)
	if err != nil {
		log.Fatalf("unable to get matchs for account %s: %v", id, err)
	}

	log.Printf("Recieved response code %d", resp.StatusCode)

	fmt.Printf("%+v \n", match)

}

func ExampleMatchService_GetMatchTimelineByID() {
	// Set environment variables before running example
	// export LEAGUE_URL=https://<region>.api.riotgames.com/
	// export LEAGUE_API_KEY=<token_value>
	url := os.Getenv("LEAGUE_URL")
	token := os.Getenv("LEAGUE_API_KEY")
	matchID := os.Getenv("MATCH_ID")

	// Setup league client with token authentication
	l, err := NewClient(url, nil)
	if err != nil {
		log.Fatal("unable to create client")
	}
	// Add the user token to the clients
	l.Authentication.SetTokenAuth(token)

	match, resp, err := l.Match.GetMatchsByAccountID(matchID)
	if err != nil {
		log.Fatalf("unable to get match timeline for match %s: %v", matchID, err)
	}

	log.Printf("Recieved response code %d", resp.StatusCode)

	fmt.Printf("%+v \n", match)

}

func ExampleMatchService_GetMatchIDsOfTournament() {
	// Set environment variables before running example
	// export LEAGUE_URL=https://<region>.api.riotgames.com/
	// export LEAGUE_API_KEY=<token_value>
	url := os.Getenv("LEAGUE_URL")
	token := os.Getenv("LEAGUE_API_KEY")
	tournamentCode := os.Getenv("TOURNAMENT_CODE")

	// Setup league client with token authentication
	l, err := NewClient(url, nil)
	if err != nil {
		log.Fatal("unable to create client")
	}
	// Add the user token to the clients
	l.Authentication.SetTokenAuth(token)

	match, resp, err := l.Match.GetMatchIDsOfTournament(tournamentCode)
	if err != nil {
		log.Fatalf("unable to get match ids for tournament %s: %v", tournamentCode, err)
	}

	log.Printf("Recieved response code %d", resp.StatusCode)

	fmt.Printf("%+v \n", match)

}

func ExampleMatchService_GetMatchByIDOfTournament() {
	// Set environment variables before running example
	// export LEAGUE_URL=https://<region>.api.riotgames.com/
	// export LEAGUE_API_KEY=<token_value>
	url := os.Getenv("LEAGUE_URL")
	token := os.Getenv("LEAGUE_API_KEY")
	tournamentCode := os.Getenv("TOURNAMENT_CODE")
	matchID := os.Getenv("TOURNAMENT_MATCH_ID")

	// Setup league client with token authentication
	l, err := NewClient(url, nil)
	if err != nil {
		log.Fatal("unable to create client")
	}
	// Add the user token to the clients
	l.Authentication.SetTokenAuth(token)

	match, resp, err := l.Match.GetMatchByIDOfTournament(matchID, tournamentCode)
	if err != nil {
		log.Fatalf("unable to get match ids for tournament %s: %v", tournamentCode, err)
	}

	log.Printf("Recieved response code %d", resp.StatusCode)

	fmt.Printf("%+v \n", match)

}

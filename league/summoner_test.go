package league

import (
	"fmt"
	"log"
	"os"
)

func ExampleSummonerService_GetByEncryptedAccountID() {
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

	// get a summoner by encrypted summoner id
	summoner, resp, err := l.Summoner.GetByEncryptedAccountID(id)
	if err != nil {
		log.Fatalf("unable to get summoners %s: %v", id, err)
	}
	log.Printf("Received response code %d", resp.StatusCode)

	fmt.Printf("%+v \n", summoner)
}

func ExampleSummonerService_GetByName() {
	// Set environment variables before running example
	// export LEAGUE_URL=https://<region>.api.riotgames.com/
	// export LEAGUE_API_KEY=<token_value>
	url := os.Getenv("LEAGUE_URL")
	token := os.Getenv("LEAGUE_API_KEY")
	name := os.Getenv("SUMMONER_NAME")

	// Setup league client with token authentication
	l, err := NewClient(url, nil)
	if err != nil {
		log.Fatal("unable to create client")
	}
	// Add the user token to the clients
	l.Authentication.SetTokenAuth(token)

	// get a summoner by name
	summoner, resp, err := l.Summoner.GetByName(name)
	if err != nil {
		log.Fatalf("unable to get summoners %s: %v", name, err)
	}
	log.Printf("Received response code %d", resp.StatusCode)

	fmt.Printf("%+v \n", summoner)
}

func ExampleSummonerService_GetByPUUID() {
	// Set environment variables before running example
	// export LEAGUE_URL=https://<region>.api.riotgames.com/
	// export LEAGUE_API_KEY=<token_value>
	url := os.Getenv("LEAGUE_URL")
	token := os.Getenv("LEAGUE_API_KEY")
	puuid := os.Getenv("SUMMONER_PUUID")

	// Setup league client with token authentication
	l, err := NewClient(url, nil)
	if err != nil {
		log.Fatal("unable to create client")
	}
	// Add the user token to the clients
	l.Authentication.SetTokenAuth(token)

	// get a summoner by puuid
	summoner, resp, err := l.Summoner.GetByPUUID(puuid)
	if err != nil {
		log.Fatalf("unable to get summoners %s: %v", puuid, err)
	}
	log.Printf("Received response code %d", resp.StatusCode)

	fmt.Printf("%+v \n", summoner)
}

func ExampleSummonerService_GetBySummonerID() {
	// Set environment variables before running example
	// export LEAGUE_URL=https://<region>.api.riotgames.com/
	// export LEAGUE_API_KEY=<token_value>
	url := os.Getenv("LEAGUE_URL")
	token := os.Getenv("LEAGUE_API_KEY")
	id := os.Getenv("SUMMONER_ID")

	// Setup league client with token authentication
	l, err := NewClient(url, nil)
	if err != nil {
		log.Fatal("unable to create client")
	}
	// Add the user token to the clients
	l.Authentication.SetTokenAuth(token)

	// get a summoner by summoner id
	summoner, resp, err := l.Summoner.GetBySummonerID(id)
	if err != nil {
		log.Fatalf("unable to get summoners %s: %v", id, err)
	}
	log.Printf("Received response code %d", resp.StatusCode)

	fmt.Printf("%+v \n", summoner)
}

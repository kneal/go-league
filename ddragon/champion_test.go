package ddragon

import (
	"fmt"
	"log"
	"os"
)

func ExampleChampionService_Get() {
	// Set environment variables before running example
	// export DDRAGON_LEAGUE_VERSION=10.9.1
	// export DDRAGON_LANGUAGE_CODE=en_US
	version := os.Getenv("DDRAGON_LEAGUE_VERSION")
	language := os.Getenv("DDRAGON_LANGUAGE_CODE")

	// Setup league client with token authentication
	l, err := NewClient(version, language, nil)
	if err != nil {
		log.Fatal("unable to create client")
	}

	// get the rotation champions
	champions, resp, err := l.Champion.Get("Aatrox")
	if err != nil {
		log.Fatalf("unable to get champions: %v", err)
	}
	log.Printf("Received response code %d", resp.StatusCode)

	fmt.Printf("%+v \n", champions)
}

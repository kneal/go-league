package league

import (
	"fmt"
)

// ChampionService handles manging champions.
type ChampionService service

// Champion represents champion rotations, including free-to-play and low-level free-to-play rotations
type Champion struct {
	FreeChampionIDs              []int `json:"freeChampionIds"`
	FreeChampionIDsForNewPlayers []int `json:"freeChampionIdsForNewPlayers"`
	MaxNewPlayerLevel            int   `json:"maxNewPlayerLevel"`
}

// GetAll returns a list of rotation champions.
//
// https://developer.riotgames.com/apis#champion-v3
func (svc *ChampionService) GetRotations() (*Champion, *Response, error) {
	// set the API endpoint path we send the request to
	u := fmt.Sprintf("/lol/platform/v3/champion-rotations")

	// slice Champion type we want to return
	v := new(Champion)

	// send request using client
	resp, err := svc.client.Call("GET", u, nil, v)

	return v, resp, err
}

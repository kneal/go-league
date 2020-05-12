package ddragon

import (
	"fmt"
)

type (

	// ChampionService handles manging champions.
	ChampionService service

	// ChampionResponse represents the json document for champions
	ChampionResponse struct {
		Type    string              `json:"type"`
		Format  string              `json:"format"`
		Version string              `json:"version"`
		Data    map[string]Champion `json:"data"`
	}

	// Child data structures used within champion

	// Champion represents the data used to define a champion
	Champion struct {
		ID    string `json:"id"`
		Key   string `json:"key"`
		Name  string `json:"name"`
		Title string `json:"title"`
		Image struct {
			Full   string `json:"full"`
			Sprite string `json:"sprite"`
			Group  string `json:"group"`
			X      int    `json:"x"`
			Y      int    `json:"y"`
			W      int    `json:"w"`
			H      int    `json:"h"`
		} `json:"image"`
		Skins []struct {
			ID      string `json:"id"`
			Num     int    `json:"num"`
			Name    string `json:"name"`
			Chromas bool   `json:"chromas"`
		} `json:"skins"`
		Lore      string   `json:"lore"`
		Blurb     string   `json:"blurb"`
		Allytips  []string `json:"allytips"`
		Enemytips []string `json:"enemytips"`
		Tags      []string `json:"tags"`
		Partype   string   `json:"partype"`
		Info      struct {
			Attack     int `json:"attack"`
			Defense    int `json:"defense"`
			Magic      int `json:"magic"`
			Difficulty int `json:"difficulty"`
		} `json:"info"`
		Stats struct {
			Hp                   int     `json:"hp"`
			Hpperlevel           int     `json:"hpperlevel"`
			Mp                   int     `json:"mp"`
			Mpperlevel           int     `json:"mpperlevel"`
			Movespeed            int     `json:"movespeed"`
			Armor                int     `json:"armor"`
			Armorperlevel        float64 `json:"armorperlevel"`
			Spellblock           float64 `json:"spellblock"`
			Spellblockperlevel   float64 `json:"spellblockperlevel"`
			Attackrange          int     `json:"attackrange"`
			Hpregen              int     `json:"hpregen"`
			Hpregenperlevel      int     `json:"hpregenperlevel"`
			Mpregen              int     `json:"mpregen"`
			Mpregenperlevel      int     `json:"mpregenperlevel"`
			Crit                 int     `json:"crit"`
			Critperlevel         int     `json:"critperlevel"`
			Attackdamage         int     `json:"attackdamage"`
			Attackdamageperlevel int     `json:"attackdamageperlevel"`
			Attackspeedperlevel  float64 `json:"attackspeedperlevel"`
			Attackspeed          float64 `json:"attackspeed"`
		} `json:"stats"`
		Spells []struct {
			ID          string `json:"id"`
			Name        string `json:"name"`
			Description string `json:"description"`
			Tooltip     string `json:"tooltip"`
			Leveltip    struct {
				Label  []string `json:"label"`
				Effect []string `json:"effect"`
			} `json:"leveltip"`
			Maxrank      int    `json:"maxrank"`
			Cooldown     []int  `json:"cooldown"`
			CooldownBurn string `json:"cooldownBurn"`
			Cost         []int  `json:"cost"`
			CostBurn     string `json:"costBurn"`
			Datavalues   struct {
			} `json:"datavalues"`
			Effect     []interface{} `json:"effect"`
			EffectBurn []interface{} `json:"effectBurn"`
			Vars       []interface{} `json:"vars"`
			CostType   string        `json:"costType"`
			Maxammo    string        `json:"maxammo"`
			Range      []int         `json:"range"`
			RangeBurn  string        `json:"rangeBurn"`
			Image      struct {
				Full   string `json:"full"`
				Sprite string `json:"sprite"`
				Group  string `json:"group"`
				X      int    `json:"x"`
				Y      int    `json:"y"`
				W      int    `json:"w"`
				H      int    `json:"h"`
			} `json:"image"`
			Resource string `json:"resource"`
		} `json:"spells"`
		Passive struct {
			Name        string `json:"name"`
			Description string `json:"description"`
			Image       struct {
				Full   string `json:"full"`
				Sprite string `json:"sprite"`
				Group  string `json:"group"`
				X      int    `json:"x"`
				Y      int    `json:"y"`
				W      int    `json:"w"`
				H      int    `json:"h"`
			} `json:"image"`
		} `json:"passive"`
		Recommended []struct {
			Champion            string      `json:"champion"`
			Title               string      `json:"title"`
			Map                 string      `json:"map"`
			Mode                string      `json:"mode"`
			Type                string      `json:"type"`
			CustomTag           string      `json:"customTag"`
			Sortrank            int         `json:"sortrank,omitempty"`
			ExtensionPage       bool        `json:"extensionPage"`
			UseObviousCheckmark bool        `json:"useObviousCheckmark,omitempty"`
			CustomPanel         interface{} `json:"customPanel"`
			Blocks              []struct {
				Type                string   `json:"type"`
				RecMath             bool     `json:"recMath"`
				RecSteps            bool     `json:"recSteps"`
				MinSummonerLevel    int      `json:"minSummonerLevel"`
				MaxSummonerLevel    int      `json:"maxSummonerLevel"`
				ShowIfSummonerSpell string   `json:"showIfSummonerSpell"`
				HideIfSummonerSpell string   `json:"hideIfSummonerSpell"`
				AppendAfterSection  string   `json:"appendAfterSection"`
				VisibleWithAllOf    []string `json:"visibleWithAllOf"`
				HiddenWithAnyOf     []string `json:"hiddenWithAnyOf"`
				Items               []struct {
					ID        string `json:"id"`
					Count     int    `json:"count"`
					HideCount bool   `json:"hideCount"`
				} `json:"items"`
			} `json:"blocks"`
		} `json:"recommended"`
	}
)

// GetAll returns a list of all champions.
//
// https://developer.riotgames.com/docs/lol#data-dragon_champions
func (svc *ChampionService) GetAll(name string) (*ChampionResponse, *Response, error) {
	// set the API endpoint path we send the request to
	u := fmt.Sprintf("/champion.json", name)

	// slice Champion type we want to return
	v := new(ChampionResponse)

	// send request using client
	resp, err := svc.client.Call("GET", u, nil, v)

	return v, resp, err
}

// Get returns a champion.
//
// https://developer.riotgames.com/docs/lol#data-dragon_champions
func (svc *ChampionService) Get(name string) (*ChampionResponse, *Response, error) {
	// set the API endpoint path we send the request to
	u := fmt.Sprintf("/champion/%s.json", name)

	// slice Champion type we want to return
	v := new(ChampionResponse)

	// send request using client
	resp, err := svc.client.Call("GET", u, nil, v)

	return v, resp, err
}

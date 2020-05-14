package league

import "fmt"

//MatchService handles getting match data
type MatchService service

//Match is a struct to hold all the data from a given match
type Match struct {
	GameID       int64  `json:"gameId"`
	PlatformID   string `json:"platformId"`
	GameCreation int64  `json:"gameCreation"`
	GameDuration int    `json:"gameDuration"`
	QueueID      int    `json:"queueId"`
	MapID        int    `json:"mapId"`
	SeasonID     int    `json:"seasonId"`
	GameVersion  string `json:"gameVersion"`
	GameMode     string `json:"gameMode"`
	GameType     string `json:"gameType"`
	Teams        []struct {
		TeamID               int    `json:"teamId"`
		Win                  string `json:"win"`
		FirstBlood           bool   `json:"firstBlood"`
		FirstTower           bool   `json:"firstTower"`
		FirstInhibitor       bool   `json:"firstInhibitor"`
		FirstBaron           bool   `json:"firstBaron"`
		FirstDragon          bool   `json:"firstDragon"`
		FirstRiftHerald      bool   `json:"firstRiftHerald"`
		TowerKills           int    `json:"towerKills"`
		InhibitorKills       int    `json:"inhibitorKills"`
		BaronKills           int    `json:"baronKills"`
		DragonKills          int    `json:"dragonKills"`
		VilemawKills         int    `json:"vilemawKills"`
		RiftHeraldKills      int    `json:"riftHeraldKills"`
		DominionVictoryScore int    `json:"dominionVictoryScore"`
		Bans                 []struct {
			ChampionID int `json:"championId"`
			PickTurn   int `json:"pickTurn"`
		} `json:"bans"`
	} `json:"teams"`
	Participants []struct {
		ParticipantID int `json:"participantId"`
		TeamID        int `json:"teamId"`
		ChampionID    int `json:"championId"`
		Spell1ID      int `json:"spell1Id"`
		Spell2ID      int `json:"spell2Id"`
		Stats         struct {
			ParticipantID                   int  `json:"participantId"`
			Win                             bool `json:"win"`
			Item0                           int  `json:"item0"`
			Item1                           int  `json:"item1"`
			Item2                           int  `json:"item2"`
			Item3                           int  `json:"item3"`
			Item4                           int  `json:"item4"`
			Item5                           int  `json:"item5"`
			Item6                           int  `json:"item6"`
			Kills                           int  `json:"kills"`
			Deaths                          int  `json:"deaths"`
			Assists                         int  `json:"assists"`
			LargestKillingSpree             int  `json:"largestKillingSpree"`
			LargestMultiKill                int  `json:"largestMultiKill"`
			KillingSprees                   int  `json:"killingSprees"`
			LongestTimeSpentLiving          int  `json:"longestTimeSpentLiving"`
			DoubleKills                     int  `json:"doubleKills"`
			TripleKills                     int  `json:"tripleKills"`
			QuadraKills                     int  `json:"quadraKills"`
			PentaKills                      int  `json:"pentaKills"`
			UnrealKills                     int  `json:"unrealKills"`
			TotalDamageDealt                int  `json:"totalDamageDealt"`
			MagicDamageDealt                int  `json:"magicDamageDealt"`
			PhysicalDamageDealt             int  `json:"physicalDamageDealt"`
			TrueDamageDealt                 int  `json:"trueDamageDealt"`
			LargestCriticalStrike           int  `json:"largestCriticalStrike"`
			TotalDamageDealtToChampions     int  `json:"totalDamageDealtToChampions"`
			MagicDamageDealtToChampions     int  `json:"magicDamageDealtToChampions"`
			PhysicalDamageDealtToChampions  int  `json:"physicalDamageDealtToChampions"`
			TrueDamageDealtToChampions      int  `json:"trueDamageDealtToChampions"`
			TotalHeal                       int  `json:"totalHeal"`
			TotalUnitsHealed                int  `json:"totalUnitsHealed"`
			DamageSelfMitigated             int  `json:"damageSelfMitigated"`
			DamageDealtToObjectives         int  `json:"damageDealtToObjectives"`
			DamageDealtToTurrets            int  `json:"damageDealtToTurrets"`
			VisionScore                     int  `json:"visionScore"`
			TimeCCingOthers                 int  `json:"timeCCingOthers"`
			TotalDamageTaken                int  `json:"totalDamageTaken"`
			MagicalDamageTaken              int  `json:"magicalDamageTaken"`
			PhysicalDamageTaken             int  `json:"physicalDamageTaken"`
			TrueDamageTaken                 int  `json:"trueDamageTaken"`
			GoldEarned                      int  `json:"goldEarned"`
			GoldSpent                       int  `json:"goldSpent"`
			TurretKills                     int  `json:"turretKills"`
			InhibitorKills                  int  `json:"inhibitorKills"`
			TotalMinionsKilled              int  `json:"totalMinionsKilled"`
			NeutralMinionsKilled            int  `json:"neutralMinionsKilled"`
			NeutralMinionsKilledTeamJungle  int  `json:"neutralMinionsKilledTeamJungle"`
			NeutralMinionsKilledEnemyJungle int  `json:"neutralMinionsKilledEnemyJungle"`
			TotalTimeCrowdControlDealt      int  `json:"totalTimeCrowdControlDealt"`
			ChampLevel                      int  `json:"champLevel"`
			VisionWardsBoughtInGame         int  `json:"visionWardsBoughtInGame"`
			SightWardsBoughtInGame          int  `json:"sightWardsBoughtInGame"`
			WardsPlaced                     int  `json:"wardsPlaced"`
			WardsKilled                     int  `json:"wardsKilled"`
			FirstBloodKill                  bool `json:"firstBloodKill"`
			FirstBloodAssist                bool `json:"firstBloodAssist"`
			FirstTowerKill                  bool `json:"firstTowerKill"`
			FirstTowerAssist                bool `json:"firstTowerAssist"`
			CombatPlayerScore               int  `json:"combatPlayerScore"`
			ObjectivePlayerScore            int  `json:"objectivePlayerScore"`
			TotalPlayerScore                int  `json:"totalPlayerScore"`
			TotalScoreRank                  int  `json:"totalScoreRank"`
			PlayerScore0                    int  `json:"playerScore0"`
			PlayerScore1                    int  `json:"playerScore1"`
			PlayerScore2                    int  `json:"playerScore2"`
			PlayerScore3                    int  `json:"playerScore3"`
			PlayerScore4                    int  `json:"playerScore4"`
			PlayerScore5                    int  `json:"playerScore5"`
			PlayerScore6                    int  `json:"playerScore6"`
			PlayerScore7                    int  `json:"playerScore7"`
			PlayerScore8                    int  `json:"playerScore8"`
			PlayerScore9                    int  `json:"playerScore9"`
			Perk0                           int  `json:"perk0"`
			Perk0Var1                       int  `json:"perk0Var1"`
			Perk0Var2                       int  `json:"perk0Var2"`
			Perk0Var3                       int  `json:"perk0Var3"`
			Perk1                           int  `json:"perk1"`
			Perk1Var1                       int  `json:"perk1Var1"`
			Perk1Var2                       int  `json:"perk1Var2"`
			Perk1Var3                       int  `json:"perk1Var3"`
			Perk2                           int  `json:"perk2"`
			Perk2Var1                       int  `json:"perk2Var1"`
			Perk2Var2                       int  `json:"perk2Var2"`
			Perk2Var3                       int  `json:"perk2Var3"`
			Perk3                           int  `json:"perk3"`
			Perk3Var1                       int  `json:"perk3Var1"`
			Perk3Var2                       int  `json:"perk3Var2"`
			Perk3Var3                       int  `json:"perk3Var3"`
			Perk4                           int  `json:"perk4"`
			Perk4Var1                       int  `json:"perk4Var1"`
			Perk4Var2                       int  `json:"perk4Var2"`
			Perk4Var3                       int  `json:"perk4Var3"`
			Perk5                           int  `json:"perk5"`
			Perk5Var1                       int  `json:"perk5Var1"`
			Perk5Var2                       int  `json:"perk5Var2"`
			Perk5Var3                       int  `json:"perk5Var3"`
			PerkPrimaryStyle                int  `json:"perkPrimaryStyle"`
			PerkSubStyle                    int  `json:"perkSubStyle"`
			StatPerk0                       int  `json:"statPerk0"`
			StatPerk1                       int  `json:"statPerk1"`
			StatPerk2                       int  `json:"statPerk2"`
		} `json:"stats"`
		Timeline struct {
			ParticipantID      int `json:"participantId"`
			CreepsPerMinDeltas struct {
				One020 float64 `json:"10-20"`
				Zero10 float64 `json:"0-10"`
			} `json:"creepsPerMinDeltas"`
			XpPerMinDeltas struct {
				One020 float64 `json:"10-20"`
				Zero10 float64 `json:"0-10"`
			} `json:"xpPerMinDeltas"`
			GoldPerMinDeltas struct {
				One020 float64 `json:"10-20"`
				Zero10 float64 `json:"0-10"`
			} `json:"goldPerMinDeltas"`
			CsDiffPerMinDeltas struct {
				One020 float64 `json:"10-20"`
				Zero10 float64 `json:"0-10"`
			} `json:"csDiffPerMinDeltas"`
			XpDiffPerMinDeltas struct {
				One020 float64 `json:"10-20"`
				Zero10 float64 `json:"0-10"`
			} `json:"xpDiffPerMinDeltas"`
			DamageTakenPerMinDeltas struct {
				One020 float64 `json:"10-20"`
				Zero10 float64 `json:"0-10"`
			} `json:"damageTakenPerMinDeltas"`
			DamageTakenDiffPerMinDeltas struct {
				One020 float64 `json:"10-20"`
				Zero10 float64 `json:"0-10"`
			} `json:"damageTakenDiffPerMinDeltas"`
			Role string `json:"role"`
			Lane string `json:"lane"`
		} `json:"timeline"`
	} `json:"participants"`
	ParticipantIdentities []struct {
		ParticipantID int `json:"participantId"`
		Player        struct {
			PlatformID        string `json:"platformId"`
			AccountID         string `json:"accountId"`
			SummonerName      string `json:"summonerName"`
			SummonerID        string `json:"summonerId"`
			CurrentPlatformID string `json:"currentPlatformId"`
			CurrentAccountID  string `json:"currentAccountId"`
			MatchHistoryURI   string `json:"matchHistoryUri"`
			ProfileIcon       int    `json:"profileIcon"`
		} `json:"player"`
	} `json:"participantIdentities"`
}

//MatchByAccountID is a struct to hold the data
//of matches from an account
type MatchByAccountID struct {
	Matches []struct {
		PlatformID string `json:"platformId"`
		GameID     int64  `json:"gameId"`
		Champion   int    `json:"champion"`
		Queue      int    `json:"queue"`
		Season     int    `json:"season"`
		Timestamp  int64  `json:"timestamp"`
		Role       string `json:"role"`
		Lane       string `json:"lane"`
	} `json:"matches"`
	StartIndex int `json:"startIndex"`
	EndIndex   int `json:"endIndex"`
	TotalGames int `json:"totalGames"`
}

//MatchByTimeline is a struct to hold the data
//from the timeline endpoint
type MatchByTimeline struct {
	Frames []struct {
		ParticipantFrames struct {
			Num1 struct {
				ParticipantID int `json:"participantId"`
				Position      struct {
					X int `json:"x"`
					Y int `json:"y"`
				} `json:"position"`
				CurrentGold         int `json:"currentGold"`
				TotalGold           int `json:"totalGold"`
				Level               int `json:"level"`
				Xp                  int `json:"xp"`
				MinionsKilled       int `json:"minionsKilled"`
				JungleMinionsKilled int `json:"jungleMinionsKilled"`
				DominionScore       int `json:"dominionScore"`
				TeamScore           int `json:"teamScore"`
			} `json:"1"`
			Num2 struct {
				ParticipantID int `json:"participantId"`
				Position      struct {
					X int `json:"x"`
					Y int `json:"y"`
				} `json:"position"`
				CurrentGold         int `json:"currentGold"`
				TotalGold           int `json:"totalGold"`
				Level               int `json:"level"`
				Xp                  int `json:"xp"`
				MinionsKilled       int `json:"minionsKilled"`
				JungleMinionsKilled int `json:"jungleMinionsKilled"`
				DominionScore       int `json:"dominionScore"`
				TeamScore           int `json:"teamScore"`
			} `json:"2"`
			Num3 struct {
				ParticipantID int `json:"participantId"`
				Position      struct {
					X int `json:"x"`
					Y int `json:"y"`
				} `json:"position"`
				CurrentGold         int `json:"currentGold"`
				TotalGold           int `json:"totalGold"`
				Level               int `json:"level"`
				Xp                  int `json:"xp"`
				MinionsKilled       int `json:"minionsKilled"`
				JungleMinionsKilled int `json:"jungleMinionsKilled"`
				DominionScore       int `json:"dominionScore"`
				TeamScore           int `json:"teamScore"`
			} `json:"3"`
			Num4 struct {
				ParticipantID int `json:"participantId"`
				Position      struct {
					X int `json:"x"`
					Y int `json:"y"`
				} `json:"position"`
				CurrentGold         int `json:"currentGold"`
				TotalGold           int `json:"totalGold"`
				Level               int `json:"level"`
				Xp                  int `json:"xp"`
				MinionsKilled       int `json:"minionsKilled"`
				JungleMinionsKilled int `json:"jungleMinionsKilled"`
				DominionScore       int `json:"dominionScore"`
				TeamScore           int `json:"teamScore"`
			} `json:"4"`
			Num5 struct {
				ParticipantID int `json:"participantId"`
				Position      struct {
					X int `json:"x"`
					Y int `json:"y"`
				} `json:"position"`
				CurrentGold         int `json:"currentGold"`
				TotalGold           int `json:"totalGold"`
				Level               int `json:"level"`
				Xp                  int `json:"xp"`
				MinionsKilled       int `json:"minionsKilled"`
				JungleMinionsKilled int `json:"jungleMinionsKilled"`
				DominionScore       int `json:"dominionScore"`
				TeamScore           int `json:"teamScore"`
			} `json:"5"`
			Num6 struct {
				ParticipantID int `json:"participantId"`
				Position      struct {
					X int `json:"x"`
					Y int `json:"y"`
				} `json:"position"`
				CurrentGold         int `json:"currentGold"`
				TotalGold           int `json:"totalGold"`
				Level               int `json:"level"`
				Xp                  int `json:"xp"`
				MinionsKilled       int `json:"minionsKilled"`
				JungleMinionsKilled int `json:"jungleMinionsKilled"`
				DominionScore       int `json:"dominionScore"`
				TeamScore           int `json:"teamScore"`
			} `json:"6"`
			Num7 struct {
				ParticipantID int `json:"participantId"`
				Position      struct {
					X int `json:"x"`
					Y int `json:"y"`
				} `json:"position"`
				CurrentGold         int `json:"currentGold"`
				TotalGold           int `json:"totalGold"`
				Level               int `json:"level"`
				Xp                  int `json:"xp"`
				MinionsKilled       int `json:"minionsKilled"`
				JungleMinionsKilled int `json:"jungleMinionsKilled"`
				DominionScore       int `json:"dominionScore"`
				TeamScore           int `json:"teamScore"`
			} `json:"7"`
			Num8 struct {
				ParticipantID int `json:"participantId"`
				Position      struct {
					X int `json:"x"`
					Y int `json:"y"`
				} `json:"position"`
				CurrentGold         int `json:"currentGold"`
				TotalGold           int `json:"totalGold"`
				Level               int `json:"level"`
				Xp                  int `json:"xp"`
				MinionsKilled       int `json:"minionsKilled"`
				JungleMinionsKilled int `json:"jungleMinionsKilled"`
				DominionScore       int `json:"dominionScore"`
				TeamScore           int `json:"teamScore"`
			} `json:"8"`
			Num9 struct {
				ParticipantID int `json:"participantId"`
				Position      struct {
					X int `json:"x"`
					Y int `json:"y"`
				} `json:"position"`
				CurrentGold         int `json:"currentGold"`
				TotalGold           int `json:"totalGold"`
				Level               int `json:"level"`
				Xp                  int `json:"xp"`
				MinionsKilled       int `json:"minionsKilled"`
				JungleMinionsKilled int `json:"jungleMinionsKilled"`
				DominionScore       int `json:"dominionScore"`
				TeamScore           int `json:"teamScore"`
			} `json:"9"`
			Num10 struct {
				ParticipantID int `json:"participantId"`
				Position      struct {
					X int `json:"x"`
					Y int `json:"y"`
				} `json:"position"`
				CurrentGold         int `json:"currentGold"`
				TotalGold           int `json:"totalGold"`
				Level               int `json:"level"`
				Xp                  int `json:"xp"`
				MinionsKilled       int `json:"minionsKilled"`
				JungleMinionsKilled int `json:"jungleMinionsKilled"`
				DominionScore       int `json:"dominionScore"`
				TeamScore           int `json:"teamScore"`
			} `json:"10"`
		} `json:"participantFrames"`
		Events    []interface{} `json:"events"`
		Timestamp int           `json:"timestamp"`
	} `json:"frames"`
	FrameInterval int `json:"frameInterval"`
}

//MatchIDs holds all the id's
type MatchIDs struct {
	MatchIDS []int64 `json:"matchIds"`
}

//GetMatchByID returns a given match from the match ID
//
//https://developer.riotgames.com/apis#match-v4
func (svc *MatchService) GetMatchByID(matchID string) (*Match, *Response, error) {
	//Set the API endpoint path we send to the request
	u := fmt.Sprintf("lol/match/v4/matches/%s", matchID)

	// slice Match type we want to return
	m := new(Match)

	// send request using client
	resp, err := svc.client.Call("GET", u, nil, m)

	return m, resp, err
}

//GetMatchsByAccount returns a match by an account ID
//
//https://developer.riotgames.com/apis#match-v4
func (svc *MatchService) GetMatchsByAccount(encryptedAccountID string) (*MatchByAccountID, *Response, error) {
	// set the API endpoint path we send to the request
	u := fmt.Sprintf("lol/match/v4/matchlists/by-account/%s", encryptedAccountID)

	// slice MatchByID we want to return
	m := new(MatchByAccountID)

	resp, err := svc.client.Call("GET", u, nil, m)

	return m, resp, err
}

//GetMatchTimelineByID returns a match with time stats
//from the match ID
//https://developer.riotgames.com/apis#match-v4
func (svc *MatchService) GetMatchTimelineByID(matchID string) (*MatchByTimeline, *Response, error) {
	// set the API endpoint path we send to the request

	u := fmt.Sprintf("lol/match/v4/timelines/by-match/%s", matchID)

	m := new(MatchByTimeline)

	resp, err := svc.client.Call("GET", u, nil, m)

	return m, resp, err
}

//GetMatchIDsTournamentCode return all matchIds from a given tournament
//
//https://developer.riotgames.com/apis#match-v4
func (svc *MatchService) GetMatchIDsTournamentCode(tournamentCode string) (*MatchIDs, *Response, error) {
	// set the API endpoint path we send to the request
	u := fmt.Sprintf("lol/match/v4/matches/by-tournament-code/%s/ids", tournamentCode)

	m := new(MatchIDs)

	resp, err := svc.client.Call("GET", u, nil, m)

	return m, resp, err
}

//GetMatchByIDOfTournament returns a specific match from a tournament
//
//https://developer.riotgames.com/apis#match-v4
func (svc *MatchService) GetMatchByIDOfTournament(matchID, tournamentCode string) (*Match, *Response, error) {
	// set the API endpoint path we send to the request

	u := fmt.Sprintf("lol/match/v4/matches/%s/by-tournament-code/%s", matchID, tournamentCode)

	m := new(Match)

	resp, err := svc.client.Call("GET", u, nil, m)

	return m, resp, err
}

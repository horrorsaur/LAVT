package local

import (
	"context"
	"encoding/json"
	"fmt"
)

type (
	SessionResponse struct {
		GameName string `json:"game_name"`
		GameTag  string `json:"game_tag"`
		PID      string `json:"pid"`
		PlayerId string `json:"puuid"`
		Region   string `json:"region"`
		State    string `json:"state"`
	}

	EntitlementResponse struct {
		AccessToken string `json:"accessToken"` // token
		Token       string `json:"token"`       // entitlement
	}

	Presence struct {
		PID      string `json:"pid"`
		Actor    string `json:"actor"`
		GameName string `json:"game_name"`
		GameTag  string `json:"game_tag"`
		State    string `json:"state"`
		Product  string `json:"product"`
		Region   string `json:"region"`
	}

	PresencesResponse struct {
		Presences []Presence
	}

	PreGameMatchResponse struct {
		MatchID string `json:"MatchID"`
		State   string `json:"State"`
		ModeID  string `json:"ModeID"`

		Players []struct {
			Subject         string `json:"Subject"` // player uuid
			TeamID          string `json:"TeamID"`
			PlayerIdentitiy struct {
				Subject          string `json:"Subject"` // player uuid
				Incognito        bool   `json:"Incognito"`
				HideAccountLevel bool   `json:"HideAccountLevel"`
			} `json:"PlayerIdentity"`
		} `json:"Players"`

		MatchmakingData interface{} `json:"MatchmakingData"`
	}
)

func (c *ValorantClient) GetSession(ctx context.Context) (SessionResponse, error) {
	url := fmt.Sprintf("https://127.0.0.1:%v/chat/v1/session", c.lockfile.Port)
	body, err := c.handleRequest(ctx, "GET", url, true, false)
	if err != nil {
		return SessionResponse{}, err
	}

	var response SessionResponse
	jsonErr := json.Unmarshal(body, &response)
	if jsonErr != nil {
		return response, jsonErr
	}

	return response, nil
}

func (c *ValorantClient) GetHelp(ctx context.Context) ([]byte, error) {
	url := fmt.Sprintf("https://127.0.0.1:%v/help", c.lockfile.Port)

	body, err := c.handleRequest(ctx, "GET", url, true, false)
	if err != nil {
		return nil, err
	}

	return body, nil
}

func (c *ValorantClient) GetPresences(ctx context.Context) (PresencesResponse, error) {
	url := fmt.Sprintf("https://127.0.0.1:%v/chat/v4/presences", c.lockfile.Port)

	body, err := c.handleRequest(ctx, "GET", url, true, false)
	if err != nil {
		return PresencesResponse{}, err
	}

	var d PresencesResponse
	jsonErr := json.Unmarshal(body, &d)
	if err != nil {
		return d, jsonErr
	}

	return d, nil
}

func (c ValorantClient) GetEntitlementsToken(ctx context.Context) (EntitlementResponse, error) {
	url := fmt.Sprintf("https://127.0.0.1:%v/entitlements/v1/token", c.lockfile.Port)
	body, err := c.handleRequest(ctx, "GET", url, true, false)
	if err != nil {
		return EntitlementResponse{}, err
	}

	var d EntitlementResponse
	jsonErr := json.Unmarshal(body, &d)
	if jsonErr != nil {
		return d, jsonErr
	}
	return d, nil
}

func (c ValorantClient) GetPreGameMatchDetails(ctx context.Context, pregameId string) (PreGameMatchResponse, error) {
	url := fmt.Sprintf(
		"https://glz-%v-1.%v.a.pvp.net/pregame/v1/matches/%v",
		"na",      // region
		"na",      // shard
		pregameId, // pregame match id
	) // region/shard hardcoded for now

	body, err := c.handleRequest(ctx, "GET", url, false, true)
	if err != nil {
		return PreGameMatchResponse{}, err
	}

	var d PreGameMatchResponse
	jsonErr := json.Unmarshal(body, &d)
	if jsonErr != nil {
		return PreGameMatchResponse{}, jsonErr
	}

	return d, nil
}

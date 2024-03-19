package local

import (
	"context"
	"encoding/json"
	"fmt"
)

type (
	LocalSession struct {
		GameName string `json:"game_name"`
		GameTag  string `json:"game_tag"`
		PID      string `json:"pid"`
		PlayerId string `json:"puuid"`
		Region   string `json:"region"`
		State    string `json:"state"`
	}

	Entitlement struct {
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

	Presences []Presence

	PreGameMatch struct {
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

func (c *ValorantClient) GetLocalSession(ctx context.Context) (LocalSession, error) {
	url := fmt.Sprintf("https://127.0.0.1:%v/chat/v1/session", c.lockfile.Port)
	var response LocalSession

	p := Params{
		ctx:        ctx,
		httpMethod: "GET",
		url:        url,
	}
	body, err := c.handleRequest(p)
	if err != nil {
		return response, err
	}

	jsonErr := json.Unmarshal(body, &response)
	if jsonErr != nil {
		return response, jsonErr
	}

	return response, nil
}

func (c *ValorantClient) GetHelp(ctx context.Context) ([]byte, error) {
	url := fmt.Sprintf("https://127.0.0.1:%v/help", c.lockfile.Port)

	p := Params{
		ctx:        ctx,
		httpMethod: "GET",
		url:        url,
	}
	body, err := c.handleRequest(p)
	if err != nil {
		return nil, err
	}

	return body, nil
}

func (c *ValorantClient) GetPresences(ctx context.Context) (Presences, error) {
	url := fmt.Sprintf("https://127.0.0.1:%v/chat/v4/presences", c.lockfile.Port)
	var d Presences

	p := Params{
		ctx:        ctx,
		httpMethod: "GET",
		url:        url,
	}
	body, err := c.handleRequest(p)
	if err != nil {
		return d, err
	}

	jsonErr := json.Unmarshal(body, &d)
	if jsonErr != nil {
		return d, jsonErr
	}

	return d, nil
}

func (c ValorantClient) GetEntitlementsToken(ctx context.Context) (Entitlement, error) {
	url := fmt.Sprintf("https://127.0.0.1:%v/entitlements/v1/token", c.lockfile.Port)
	var d Entitlement

	p := Params{
		ctx:        ctx,
		httpMethod: "GET",
		url:        url,
	}
	body, err := c.handleRequest(p)
	if err != nil {
		return d, err
	}

	jsonErr := json.Unmarshal(body, &d)
	if jsonErr != nil {
		return d, jsonErr
	}
	return d, nil
}

func (c ValorantClient) GetPreGameMatchDetails(ctx context.Context, pregameId string, entitlements Entitlement) (PreGameMatch, error) {
	url := fmt.Sprintf(
		"https://glz-%v-1.%v.a.pvp.net/pregame/v1/matches/%v",
		"na",      // region
		"na",      // shard
		pregameId, // pregame match id
	) // region/shard hardcoded for now
	var d PreGameMatch

	p := Params{
		ctx:          ctx,
		httpMethod:   "GET",
		url:          url,
		entitlements: entitlements,
	}
	body, err := c.handleRequest(p)
	if err != nil {
		return d, err
	}

	jsonErr := json.Unmarshal(body, &d)
	if jsonErr != nil {
		return d, jsonErr
	}

	return d, nil
}

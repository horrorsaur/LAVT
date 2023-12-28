package local

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
)

type (
	// general player info
	SessionResponse struct {
		Name     string `json:"name"`
		PID      string `json:"pid"`
		PlayerId string `json:"puuid"`
		State    string `json:"state"`
	}

	// credentials used for some endpoints
	EntitlementResponse struct {
		Token       string `json:"accessToken"`
		Entitlement string `json:"token"`
	}

	// friend data
	PresencesResponse struct {
		Presences []Presence
	}
)

func (c *ValorantClient) GetSession(ctx context.Context) (SessionResponse, error) {
	url := fmt.Sprintf("https://127.0.0.1:%v/chat/v1/session", c.lockfile.Port)
	body, err := c.makeRequest(ctx, url, "GET")
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

	body, err := c.makeRequest(ctx, url, "GET")
	if err != nil {
		return nil, err
	}
	return body, nil
}

func (c *ValorantClient) GetPresences(ctx context.Context) (PresencesResponse, error) {
	url := fmt.Sprintf("https://127.0.0.1:%v/chat/v4/presences", c.lockfile.Port)
	body, err := c.makeRequest(ctx, url, "GET")
	if err != nil {
		return PresencesResponse{}, err
	}

	var d PresencesResponse
	jsonErr := json.Unmarshal(body, &d)
	if err != nil {
		log.Print(jsonErr)
		return d, jsonErr
	}

	return d, nil
}

func (c ValorantClient) GetEntitlementsToken(ctx context.Context) (EntitlementResponse, error) {
	url := fmt.Sprintf("https://127.0.0.1:%v/entitlements/v1/token", c.lockfile.Port)
	body, err := c.makeRequest(ctx, url, "GET")
	if err != nil {
		return EntitlementResponse{}, err
	}

	var d EntitlementResponse
	jsonErr := json.Unmarshal(body, &d)
	if jsonErr != nil {
		log.Print(jsonErr)
		return d, jsonErr
	}
	return d, nil
}

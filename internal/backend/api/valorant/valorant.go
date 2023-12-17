package valorant

import (
	"context"
	"encoding/json"
	"fmt"
)

type (
	API struct {
		lockfile *LockfileWatcher
	}

	// holds session info from Riot client
	SessionInfo struct {
		PID      string
		PlayerId string
	}

	SessionResponse struct {
		Name     string `json:"name"`
		PID      string `json:"pid"`
		PlayerId string `json:"puuid"`
		State    string `json:"state"`
	}
)

func (c *ValorantClient) GetSession(ctx context.Context) (*SessionResponse, error) {
	url := fmt.Sprintf("https://127.0.0.1:%v/chat/v1/session", c.lockfile.Port)
	body, err := c.makeRequest(ctx, url, "GET")
	if err != nil {
		return nil, err
	}

	var response SessionResponse
	jsonErr := json.Unmarshal(body, &response)
	if jsonErr != nil {
		return nil, jsonErr
	}

	// cache session info on client
	c.session.PlayerId = response.PlayerId

	return &response, nil
}

func (w *ValorantClient) GetHelp(ctx context.Context) ([]byte, error) {
	url := fmt.Sprintf("https://127.0.0.1:%v/help", w.lockfile.Port)

	body, err := w.makeRequest(ctx, url, "GET")
	if err != nil {
		return nil, err
	}
	return body, nil
}

func (w *ValorantClient) GetPresences(ctx context.Context) ([]byte, error) {
	url := fmt.Sprintf("https://127.0.0.1:%v/chat/v4/presences", w.lockfile.Port)
	body, err := w.makeRequest(ctx, url, "GET")
	if err != nil {
		return nil, err
	}
	return body, nil
}

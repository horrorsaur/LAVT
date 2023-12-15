package valorant

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/wailsapp/wails/v2/pkg/runtime"
)

type (
	API struct {
		lockfile *LockfileWatcher
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

func (w *ValorantClient) GetPresences(ctx context.Context) {
	runtime.LogInfo(ctx, "get presences")
}

package local

// Package local provides network IO with the local instance of the Riot Client running on the host PC
// http is a standard HTTP transport
// socket is a websocket transport

import (
	"context"
	"crypto/tls"
	"io"
	"net/http"
	"strings"
	"time"

	"github.com/wailsapp/wails/v2/pkg/runtime"
	"nhooyr.io/websocket"
)

const REQUEST_TIMEOUT = 30 * time.Second
const WEBSOCKET_TIMEOUT = 30 * time.Second

type (
	ValorantClient struct {
		lockfile *RiotClientLockfileInfo
		http     *http.Client
		socket   *websocket.Conn

		registeredEvents []string
	}

	// General request data for the Valorant local APIs
	Params struct {
		ctx          context.Context
		httpMethod   string
		url          string
		entitlements Entitlement
	}
)

// Returns a new RiotClient from reading the lockfile at LOCALAPPDIR
func NewClient(lockfile *RiotClientLockfileInfo) *ValorantClient {
	return &ValorantClient{
		lockfile: lockfile,
		http: &http.Client{
			Transport: &http.Transport{
				TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
			},
		},
	}
}

// Returns bytes from HTTP body
func (c *ValorantClient) handleRequest(params Params) ([]byte, error) {
	ctx, cancelFunc := context.WithTimeout(params.ctx, REQUEST_TIMEOUT)
	defer cancelFunc()

	req, err := http.NewRequestWithContext(ctx, params.httpMethod, params.url, nil)
	if err != nil {
		return nil, err
	}

	if addBasicAuth := strings.Contains(params.url, "127.0.0.1"); addBasicAuth {
		runtime.LogDebug(ctx, "adding basic auth header")
		req.SetBasicAuth("riot", c.lockfile.Password)
	}

	if params.entitlements.AccessToken != "" && params.entitlements.Token != "" {
		runtime.LogDebug(ctx, "adding X-Riot-Entitlements-JWT and Auth Header")
		req.Header.Add("X-Riot-Entitlements-JWT", params.entitlements.Token)
		req.Header.Add("Authorization", "Bearer "+params.entitlements.AccessToken)
	}

	resp, err := c.http.Do(req)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	runtime.LogDebugf(ctx, "RESPONSE: %+v\n", string(body))
	return body, nil
}

package local

import (
	"context"
	"crypto/tls"
	"io"
	"log"
	"net/http"
	"time"

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

// move auth headers over to some kind of config struct or variadic parameter
func (c *ValorantClient) handleRequest(ctx context.Context, httpMethod, url string, includeBasicAuth bool, entitlementsHeader bool) ([]byte, error) {
	ctx, cancelFunc := context.WithTimeout(ctx, REQUEST_TIMEOUT)
	defer cancelFunc()

	req, err := http.NewRequestWithContext(ctx, httpMethod, url, nil)
	if err != nil {
		return nil, err
	}

	if includeBasicAuth {
		req.SetBasicAuth("riot", c.lockfile.Password)
	}

	if entitlementsHeader {
		e := ctx.Value("entitlements").(EntitlementResponse)
		req.Header.Add("X-Riot-Entitlements-JWT", e.Token)
		req.Header.Add("Authorization", "Bearer "+e.AccessToken)
	}

	log.Printf("REQUEST: %+v\n", req)
	resp, err := c.http.Do(req)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	log.Printf("RESPONSE: %+v\n", string(body))
	return body, nil
}

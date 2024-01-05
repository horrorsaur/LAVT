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

// builds and makes a request to the valorant local client
func (c *ValorantClient) makeRequest(ctx context.Context, url, httpMethod string) ([]byte, error) {
	ctx, cancelFunc := context.WithTimeout(ctx, REQUEST_TIMEOUT)
	defer cancelFunc() // cancel based on timeout

	req, err := http.NewRequestWithContext(ctx, httpMethod, url, nil)
	if err != nil {
		return nil, err
	}

	// auth
	e, ok := ctx.Value("entitlements").(EntitlementResponse)
	if ok {
		log.Print("using entitlements response auth headers")
		req.Header.Add("Authorization", "Bearer "+e.AccessToken)
		req.Header.Add("X-Riot-Entitlements-JWT", e.Token)
	} else {
		log.Printf("using lockfile basic auth creds")
		req.SetBasicAuth("riot", c.lockfile.Password)
	}

	log.Printf("REQ: %+v\nHEADERS: %+v\n", req, req.Header)

	resp, err := c.http.Do(req)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	log.Printf("RESP BODY: %+v\n", string(body))
	return body, nil
}

package valorant

import (
	"context"
	"crypto/tls"
	"fmt"
	"io"
	"net/http"
	"strconv"
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

		baseUrl string
	}
)

// Returns a new RiotClient from reading the lockfile at LOCALAPPDIR
func NewClient(lockfile *RiotClientLockfileInfo, baseUrl string) *ValorantClient {
	return &ValorantClient{
		lockfile: lockfile,
		http: &http.Client{
			Transport: &http.Transport{
				TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
			},
		},
		baseUrl: baseUrl,
	}
}

// connect to the local valorant client websocket
func (w *ValorantClient) ConnectToWS(ctx context.Context) {
	conn, err := w.connectToRiotWS(ctx)
	if err != nil {
		runtime.LogError(ctx, err.Error())
	}
	w.socket = conn
}

func (w *ValorantClient) connectToRiotWS(ctx context.Context) (*websocket.Conn, error) {
	ctx, cancel := context.WithTimeout(ctx, REQUEST_TIMEOUT)
	defer cancel()

	url := fmt.Sprintf(
		"ws://riot:%s@localhost:%s/",
		w.lockfile.Password,
		strconv.Itoa(w.lockfile.Port),
	)

	runtime.LogInfof(ctx, "connecting to %v", url)
	conn, _, err := websocket.Dial(ctx, url, nil)
	if err != nil {
		runtime.LogError(ctx, err.Error())
		return nil, err
	}

	return conn, nil
}

func (c *ValorantClient) makeRequest(ctx context.Context, url, httpMethod string) ([]byte, error) {
	ctx, cancelFunc := context.WithTimeout(ctx, REQUEST_TIMEOUT)
	defer cancelFunc() // cancel based on timeout

	req, err := http.NewRequestWithContext(ctx, httpMethod, url, nil)
	if err != nil {
		return nil, err
	}

	req.SetBasicAuth("riot", c.lockfile.Password)

	resp, err := c.http.Do(req)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return body, nil
}

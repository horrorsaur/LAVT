package valorant

import (
	"context"
	"crypto/tls"
	"fmt"
	"io"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/wailsapp/wails/v2/pkg/runtime"
	"nhooyr.io/websocket"
	"nhooyr.io/websocket/wsjson"
)

const REQUEST_TIMEOUT = 30 * time.Second
const WEBSOCKET_TIMEOUT = 10 * time.Second

type (
	ValorantClient struct {
		lockfile *RiotClientLockfileInfo
		http     *http.Client
		socket   *websocket.Conn

		baseUrl string
		session SessionInfo
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
func (c *ValorantClient) ConnectToWS(ctx context.Context) {
	conn, err := c.connectToRiotWS(ctx)
	log.Printf("res: %+v", conn)
	if err != nil {
		runtime.LogError(ctx, err.Error())
	}
	c.socket = conn
}

func (c ValorantClient) SubscribeToAll(ctx context.Context) {
	ctx, cancel := context.WithTimeout(ctx, WEBSOCKET_TIMEOUT)
	defer cancel()

	// [5, "OnJsonApiEvent"]
	var msg = []interface{}{5, "OnJsonApiEvent"}
	err := wsjson.Write(ctx, c.socket, msg)
	if err != nil {
		log.Print(err)
	}
}

func (w *ValorantClient) connectToRiotWS(ctx context.Context) (*websocket.Conn, error) {
	ctx, cancel := context.WithTimeout(ctx, WEBSOCKET_TIMEOUT)
	defer cancel()

	url := fmt.Sprintf(
		"wss://riot:%s@127.0.0.1:%s",
		w.lockfile.Password,
		strconv.Itoa(w.lockfile.Port),
	)

	opts := &websocket.DialOptions{HTTPClient: w.http}
	log.Printf("WS attempting connection to %v...", url)
	conn, _, err := websocket.Dial(ctx, url, opts)
	if err != nil {
		runtime.LogError(ctx, err.Error())
		return nil, err
	}

	// todo: do something w/ close func since we typically defer it
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

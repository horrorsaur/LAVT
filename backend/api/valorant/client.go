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

	"nhooyr.io/websocket"
)

const REQUEST_TIMEOUT = 30 * time.Second

type (
	ValorantClient struct {
		lockfile *RiotClientLockfileInfo
		http     *http.Client
		socket   *websocket.Conn
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

func (w *ValorantClient) Connect(ctx context.Context) {
	w.connectToRiotWS(ctx)
}

func (w *ValorantClient) connectToRiotWS(ctx context.Context) *websocket.Conn {
	ctx, cancel := context.WithTimeout(ctx, REQUEST_TIMEOUT)
	defer cancel()

	url := fmt.Sprintf(
		"ws://riot:%s@localhost:%s/",
		w.lockfile.Password,
		strconv.Itoa(w.lockfile.Port),
	)

	log.Printf("connecting to %v", url)
	conn, resp, err := websocket.Dial(ctx, url, nil)

	log.Printf("resp: \n\n %v+", resp)
	log.Printf("conn: \n\n %v+", conn)

	if err != nil {
		panic(err)
	}

	// dialOptions := &websocket.DialOptions{
	// 	HTTPHeader: headers,
	// }

	return conn
}

// TODO: refactor request creation behind pkg
func (w *ValorantClient) GetSession(ctx context.Context) {
	ctx, cancel := context.WithTimeout(ctx, REQUEST_TIMEOUT)
	defer cancel()

	url := fmt.Sprintf("https://127.0.0.1:%v/chat/v1/session", w.lockfile.Port)
	req, err := http.NewRequestWithContext(ctx, "GET", url, nil)
	req.SetBasicAuth("riot", w.lockfile.Password) // this is neat
	if err != nil {
		panic(err)
	}

	resp, err := w.http.Do(req)
	if err != nil {
		panic(err)
	}

	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}

	log.Printf("resp: \n\n %+v", string(body))
}

func (w *ValorantClient) GetHelp(ctx context.Context) {
	ctx, cancel := context.WithTimeout(ctx, REQUEST_TIMEOUT)
	defer cancel()

	url := fmt.Sprintf("https://127.0.0.1:%v/help", w.lockfile.Port)
	req, err := http.NewRequestWithContext(ctx, "GET", url, nil)
	req.SetBasicAuth("riot", w.lockfile.Password) // this is neat
	if err != nil {
		panic(err)
	}

	resp, err := w.http.Do(req)
	if err != nil {
		panic(err)
	}

	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}

	log.Printf("resp: \n\n %+v", string(body))
}

func (w *ValorantClient) GetPresences(ctx context.Context) {
	log.Printf("get presences")
}

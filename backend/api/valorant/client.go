package valorant

import (
	"context"
	"fmt"
	"log"
	"strconv"
	"time"

	"nhooyr.io/websocket"
)

type (
	ValorantClient struct {
		lockfile *RiotClientLockfileInfo
		conn     *websocket.Conn
	}
)

// Returns a new RiotClient from reading the lockfile at LOCALAPPDIR
func NewClient(lockfile *RiotClientLockfileInfo) *ValorantClient {
	return &ValorantClient{
		lockfile: lockfile,
	}
}

func (w *ValorantClient) Connect(ctx context.Context) {
	w.connectToRiotWS(ctx)
}

func (w *ValorantClient) connectToRiotWS(ctx context.Context) *websocket.Conn {
	ctx, cancel := context.WithTimeout(ctx, time.Minute)
	defer cancel()

	// password := base64.StdEncoding.EncodeToString([]byte(w.lockfile.Password))
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

	// headers := http.Header{}
	// headers.Add("Authorization", fmt.Sprintf("Basic %v", w.lockfile.Password))

	// dialOptions := &websocket.DialOptions{
	// 	HTTPHeader: headers,
	// }

	return conn
}

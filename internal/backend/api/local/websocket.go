package local

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"strconv"

	"nhooyr.io/websocket"
	"nhooyr.io/websocket/wsjson"
)

type (
	// generic type for sending websocket data to the Riot LCU WS socket
	Msg []interface{}

	WSResponse struct {
		Opcode    int
		EventName string
		Msg       struct {
			Data      interface{} `json:"data"`
			EventType string      `json:"eventType"`
			Uri       string      `json:"uri"`
		}
	}
)

// the general format from the WS messages is an array containing opcode, eventName, {"data", "eventType", "uri"}
//
// Example:
//
// [8, "OnJsonApiEvent", {"data":[], "eventType":"Update", "uri":"/lol-ranked/v1/notifications"}]
func (r *WSResponse) UnmarshalJSON(buf []byte) error {
	// https://eagain.net/articles/go-json-array-to-struct/
	tmp := []interface{}{&r.Opcode, &r.EventName, &r.Msg}
	expLen := len(tmp)

	err := json.Unmarshal(buf, &tmp)
	if err != nil {
		return err
	}

	if g, e := len(tmp), expLen; g != e {
		return fmt.Errorf("unexpected msg length: %d != %d", g, e)
	}

	return nil
}

// connect to the local valorant client websocket
//
// sets the socket field on the client
func (c *ValorantClient) ConnectToSocket(ctx context.Context) error {
	return c.connect(ctx)
}

func (c *ValorantClient) CloseSocket(ctx context.Context) error {
	return c.socket.CloseNow()
}

// sends opcode 5 (subscribe) w/ event to the websocket
func (c *ValorantClient) Subscribe(ctx context.Context, event string) error {
	return c.sendMsg(ctx, Msg{5, event})
}

// sends opcode 6 (unsubscribe) w/ event to the websocket
func (c *ValorantClient) Unsubscribe(ctx context.Context, event string) error {
	return c.sendMsg(ctx, Msg{6, event})
}

// opcode -1 is an error
func (c *ValorantClient) ReceiveSocketMsgs(ctx context.Context) WSResponse {
	b, err := c.receiveMsg(ctx)
	if err != nil {
		log.Print(err)
		return WSResponse{Opcode: -1}
	}

	var r WSResponse

	jsonErr := json.Unmarshal(b, &r)
	if jsonErr != nil {
		log.Print(jsonErr)
		return WSResponse{Opcode: -1}
	}

	// log.Printf("ws msg: %+v\n", r)
	return r
}

func (c *ValorantClient) connect(ctx context.Context) error {
	ctx, cancel := context.WithTimeout(ctx, WEBSOCKET_TIMEOUT)
	defer cancel()

	url := fmt.Sprintf(
		"wss://riot:%s@127.0.0.1:%s",
		c.lockfile.Password,
		strconv.Itoa(c.lockfile.Port),
	)

	conn, _, err := websocket.Dial(ctx, url, &websocket.DialOptions{HTTPClient: c.http})
	if err != nil {
		return err
	}

	c.socket = conn
	return nil
}

func (c *ValorantClient) sendMsg(ctx context.Context, data interface{}) error {
	ctx, cancel := context.WithTimeout(ctx, WEBSOCKET_TIMEOUT)
	defer cancel()
	return wsjson.Write(ctx, c.socket, data)
}

// 8 will be the opcode, the second item will be the name of the event and
// the third item will be a JSON blob with 3 entries: data, eventType and uri.
func (c *ValorantClient) receiveMsg(ctx context.Context) ([]byte, error) {
	_, r, err := c.socket.Reader(ctx)
	if err != nil {
		return nil, err
	}

	body, err := io.ReadAll(r)
	if err != nil {
		return nil, err
	}

	return body, nil
}

package ynison

import (
	"crypto/rand"
	"encoding/json"
	"net/http"
	"net/url"

	"github.com/gorilla/websocket"
)

type Client struct {
	token    string
	deviceID string
	header   http.Header
	conn     *Conn
}

func NewClient(token string) *Client {
	h := make(http.Header)
	h.Set("Origin", "https://music.yandex.ru")
	h.Set("Authorization", "OAuth "+token)
	return &Client{
		token:    token,
		deviceID: randString(16),
		header:   h.Clone(),
		conn:     new(Conn),
	}
}

func (y *Client) getTicket() (*RedirectResponse, error) {
	// потом переделаю
	header := y.header.Clone()
	header.Set("Sec-WebSocket-Protocol", `Bearer, v2, {"Ynison-Device-Id":"`+y.deviceID+`","Ynison-Device-Info":"{\"app_name\":\"Chrome\",\"type\":1}"}`)
	c, resp, err := websocket.DefaultDialer.Dial("wss://ynison.music.yandex.ru/redirector.YnisonRedirectService/GetRedirectToYnison", header)
	if err != nil {
		_ = resp
		return nil, err
	}
	defer c.Close()
	_, message, err := c.ReadMessage()
	if err != nil {
		return nil, err
	}
	r := new(RedirectResponse)
	json.Unmarshal(message, r)
	err = c.WriteMessage(websocket.CloseMessage, websocket.FormatCloseMessage(websocket.CloseNormalClosure, ""))
	if err != nil {
		return nil, err
	}

	return r, nil
}

// Get ticket and connect to websocket
func (y *Client) Connect() error {
	r, err := y.getTicket()
	if err != nil {
		return err
	}
	u := url.URL{Scheme: "wss", Host: r.Host, Path: "/ynison_state.YnisonStateService/PutYnisonState"}
	header := y.header.Clone()
	// некрасиво, но работает
	header.Set("Sec-WebSocket-Protocol", `Bearer, v2, {"Ynison-Device-Id":"`+y.deviceID+`","Ynison-Redirect-Ticket":"`+r.RedirectTicket+`","Ynison-Session-Id":"`+r.SessionID+`","Ynison-Device-Info":"{\"app_name\":\"Chrome\",\"type\":1}"}`)

	y.OnConnect(func() {
		// некрасиво, но работает x2
		y.conn.Listen(`{"update_full_state":{"player_state":{"player_queue":{"current_playable_index":-1,"entity_id":"","entity_type":"VARIOUS","playable_list":[],"options":{"repeat_mode":"NONE"},"entity_context":"BASED_ON_ENTITY_BY_DEFAULT","version":{"device_id":"` + y.deviceID + `","version":9021243204784341000,"timestamp_ms":0},"from_optional":""},"status":{"duration_ms":0,"paused":true,"playback_speed":1,"progress_ms":0,"version":{"device_id":"` + y.deviceID + `","version":8321822175199937000,"timestamp_ms":0}}},"device":{"capabilities":{"can_be_player":true,"can_be_remote_controller":false,"volume_granularity":16},"info":{"device_id":"` + y.deviceID + `","type":"WEB","title":"Chrome Browser","app_name":"Chrome"},"volume_info":{"volume":0},"is_shadow":true},"is_currently_active":false},"rid":"ac281c26-a047-4419-ad00-e4fbfda1cba3","player_action_timestamp_ms":0,"activity_interception_type":"DO_NOT_INTERCEPT_BY_DEFAULT"}`)
	})
	err = y.conn.Connect(u.String(), header)
	return err
}

// Close connection
func (y *Client) Close() {
	y.conn.Close()
}

// OnMessage event
func (y *Client) OnMessage(f func(AudioMessage)) {
	y.conn.OnMessage(f)
}

// OnConnect event
func (y *Client) OnConnect(f func()) {
	y.conn.OnConnect(f)
}

// OnDisconnect event
func (y *Client) OnDisconnect(f func()) {
	y.conn.OnDisconnect(f)
}

// IsConnected returns true if the socket is actively connected
func (y *Client) IsConnected() bool {
	return y.conn.isConnected
}

func randString(n int) string {
	const alphanum = "0123456789abcdefghijklmnopqrstuvwxyz"
	var bytes = make([]byte, n)
	rand.Read(bytes)
	for i, b := range bytes {
		bytes[i] = alphanum[b%byte(len(alphanum))]
	}
	return string(bytes)
}

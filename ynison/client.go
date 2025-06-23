package ynison

import (
	"crypto/rand"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"time"

	"github.com/google/uuid"
	"github.com/gorilla/websocket"
)

type Client struct {
	token         string
	deviceID      string
	header        http.Header
	conn          *Conn
	ConfigMessage string
}

// нельзя использовать в качестве пульта
func NewClient(token string) *Client {
	h := make(http.Header)
	h.Set("Origin", "https://music.yandex.ru")
	h.Set("Authorization", "OAuth "+token)
	deviceID := randString(16)
	return &Client{
		token:         token,
		deviceID:      deviceID,
		header:        h.Clone(),
		conn:          new(Conn),
		ConfigMessage: `{"update_full_state":{"player_state":{"player_queue":{"current_playable_index":-1,"entity_id":"","entity_type":"VARIOUS","playable_list":[],"options":{"repeat_mode":"NONE"},"entity_context":"BASED_ON_ENTITY_BY_DEFAULT","version":{"device_id":"` + deviceID + `","version":9021243204784341000,"timestamp_ms":0},"from_optional":""},"status":{"duration_ms":0,"paused":true,"playback_speed":1,"progress_ms":0,"version":{"device_id":"` + deviceID + `","version":8321822175199937000,"timestamp_ms":0}}},"device":{"capabilities":{"can_be_player":false,"can_be_remote_controller":false,"volume_granularity":0},"info":{"device_id":"` + deviceID + `","type":"WEB","title":"go-YaYnison","app_name":"Chrome"},"volume_info":{"volume":0},"is_shadow":true},"is_currently_active":false},"rid":"ac281c26-a047-4419-ad00-e4fbfda1cba3","player_action_timestamp_ms":0,"activity_interception_type":"DO_NOT_INTERCEPT_BY_DEFAULT"}`,
	}
}

// deviceID задается вручную, поддерживает удаленное управление
func NewClientWithDeviceID(token, deviceID string) *Client {
	h := make(http.Header)
	h.Set("Origin", "https://music.yandex.ru")
	h.Set("Authorization", "OAuth "+token)
	return &Client{
		token:         token,
		deviceID:      deviceID,
		header:        h.Clone(),
		conn:          new(Conn),
		ConfigMessage: `{"update_full_state":{"player_state":{"player_queue":{"current_playable_index":-1,"entity_id":"","entity_type":"VARIOUS","playable_list":[],"options":{"repeat_mode":"NONE"},"entity_context":"BASED_ON_ENTITY_BY_DEFAULT","version":{"device_id":"` + deviceID + `","version":9021243204784341000,"timestamp_ms":0},"from_optional":""},"status":{"duration_ms":0,"paused":true,"playback_speed":1,"progress_ms":0,"version":{"device_id":"` + deviceID + `","version":8321822175199937000,"timestamp_ms":0}}},"device":{"capabilities":{"can_be_player":false,"can_be_remote_controller":true,"volume_granularity":0},"info":{"device_id":"` + deviceID + `","type":"WEB","title":"go-YaYnison","app_name":"Chrome"},"volume_info":{"volume":0},"is_shadow":false},"is_currently_active":false},"rid":"ac281c26-a047-4419-ad00-e4fbfda1cba3","player_action_timestamp_ms":0,"activity_interception_type":"DO_NOT_INTERCEPT_BY_DEFAULT"}`,
	}
}

func (y *Client) getTicket() (*RedirectResponse, error) {
	// потом переделаю
	header := y.header.Clone()
	header.Set("Sec-WebSocket-Protocol", `Bearer, v2, {"Ynison-Device-Id":"`+y.deviceID+`","Ynison-Device-Info":"{\"app_name\":\"Chrome\",\"type\":1}"}`)
	c, resp, err := websocket.DefaultDialer.Dial("wss://ynison.music.yandex.ru/redirector.YnisonRedirectService/GetRedirectToYnison", header)
	if resp != nil {
		resp.Body.Close()
	}
	if err != nil {
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
		y.conn.Send(y.ConfigMessage)
	})
	err = y.conn.Connect(u.String(), header)
	return err
}

// Обновить громкость устройства.
//
// device id устройства, на котором меняется громкость
//
// новое значение состояния громкости.
func (y *Client) UpdateVolumeInfo(deviceID string, volume float64) {
	// Example
	// { "update_volume_info": { "device_id": "target deviceID", "volume_info": { "volume": 0.5, "version": { "device_id": "deviceID", "version": 0, "timestamp_ms": 0 } } }, "rid": ""}
	rid, _ := uuid.NewUUID()
	data := new(PutYnisonStateRequest)
	data.UpdateVolumeInfo = new(UpdateVolumeInfo)
	data.UpdateVolumeInfo.VolumeInfo.Volume = volume
	data.UpdateVolumeInfo.DeviceID = deviceID
	data.Rid = rid.String()
	bytes, _ := json.Marshal(data)
	y.conn.SendBytes(bytes)
}

// Обновить статус воспроизведения.
//
// Отправляется в следующих случаях:
//
// * Старт воспроизведения (после паузы).
//
// * Остановка воспроизведения/пауза (после проигрывания).
//
// * Перемотка.
//
// * Изменение скорости воспроизведения.
func (y *Client) UpdatePlayingStatus(paused bool, ProgressMs, DurationMs string) {
	// Example
	// { "update_playing_status": { "playing_status": { "progress_ms": 0, "duration_ms": 0, "paused": false, "playback_speed": 1, "version": { "device_id": "deviceID here", "version": 0, "timestamp_ms": "" } } }, "rid": ""}
	rid, _ := uuid.NewUUID()
	timestamp := fmt.Sprint(time.Now().UnixMilli())
	data := new(PutYnisonStateRequest)
	data.UpdatePlayingStatus = new(UpdatePlayingStatus)
	data.UpdatePlayingStatus.PlayingStatus = PlayingStatus{
		Paused:        paused,
		ProgressMs:    ProgressMs,
		DurationMs:    DurationMs,
		PlaybackSpeed: 1,
		Version: Version{
			DeviceID:    y.deviceID,
			TimestampMs: timestamp,
		},
	}
	data.Rid = rid.String()
	bytes, _ := json.Marshal(data)
	y.conn.SendBytes(bytes)
}

// Обновить активное устройство.
func (y *Client) UpdateActiveDevice(deviceID string) {
	rid, _ := uuid.NewUUID()
	data := new(PutYnisonStateRequest)
	data.UpdateActiveDevice = new(UpdateActiveDevice)
	data.UpdateActiveDevice.DeviceIDOptional = deviceID
	data.Rid = rid.String()
	bytes, _ := json.Marshal(data)
	y.conn.SendBytes(bytes)
}

// Close connection
func (y *Client) Close() {
	y.conn.Close()
}

// OnMessage event
func (y *Client) OnMessage(f func(PutYnisonStateResponse)) {
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

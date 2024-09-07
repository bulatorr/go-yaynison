package ynisonGrpc

import (
	"context"
	"crypto/rand"
	"crypto/tls"
	"crypto/x509"

	"github.com/bulatorr/go-yaynison/redirector"
	"github.com/bulatorr/go-yaynison/ynisonstate"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/metadata"
	"google.golang.org/protobuf/types/known/wrapperspb"
)

type Client struct {
	token         string
	deviceID      string
	header        map[string]string
	conn          *Conn
	ConfigMessage *ynisonstate.PutYnisonStateRequest
}

// deviceID задается вручную, поддерживает удаленное управление
func NewClient(token, deviceId string) *Client {
	h := make(map[string]string)
	h["Authorization"] = "OAuth " + token
	h["Ynison-Device-Id"] = deviceId
	request := new(ynisonstate.PutYnisonStateRequest)
	request.Parameters = &ynisonstate.PutYnisonStateRequest_UpdateFullState{
		UpdateFullState: &ynisonstate.UpdateFullState{
			PlayerState: &ynisonstate.PlayerState{
				PlayerQueue: &ynisonstate.PlayerQueue{
					CurrentPlayableIndex: -1,
					EntityType:           ynisonstate.PlayerQueue_VARIOUS,
					Options: &ynisonstate.PlayerStateOptions{
						RepeatMode: ynisonstate.PlayerStateOptions_NONE,
					},
					EntityContext: ynisonstate.PlayerQueue_BASED_ON_ENTITY_BY_DEFAULT,
				},
				Status: &ynisonstate.PlayingStatus{
					Paused:        true,
					PlaybackSpeed: 1,
				},
			},
			Device: &ynisonstate.UpdateDevice{
				Info: &ynisonstate.DeviceInfo{
					DeviceId: deviceId,
					Type:     ynisonstate.DeviceType_WEB,
					Title:    "go-YaYnison",
					AppName:  "Chrome",
				},
				Capabilities: &ynisonstate.DeviceCapabilities{
					CanBePlayer:           false,
					CanBeRemoteController: true,
				},
			},
		},
	}
	return &Client{
		token:         token,
		deviceID:      deviceId,
		header:        h,
		conn:          new(Conn),
		ConfigMessage: request,
	}
}

func (y *Client) getTicket() (*redirector.RedirectResponse, error) {
	// потом переделаю x2
	tlsCreds, _ := generateTLSCreds()
	conn, err := grpc.NewClient("ynison.music.yandex.net", grpc.WithTransportCredentials(tlsCreds))
	if err != nil {
		return nil, err
	}
	defer conn.Close()
	r := redirector.NewYnisonRedirectServiceClient(conn)
	md := metadata.New(y.header)
	ctx := metadata.NewOutgoingContext(context.Background(), md)
	resp, err := r.GetRedirectToYnison(ctx, &redirector.RedirectRequest{}, grpc.Header(&md))
	return resp, err
}

// Get ticket and connect to websocket
func (y *Client) Connect() error {
	r, err := y.getTicket()
	if err != nil {
		return err
	}
	header := y.header
	header["ynison-redirect-ticket"] = r.RedirectTicket

	y.OnConnect(func() {
		// некрасиво, но работает x2
		y.conn.Send(y.ConfigMessage)
	})
	err = y.conn.Connect(r.GetHost(), header)
	return err
}

// Обновить громкость устройства.
//
// device id устройства, на котором меняется громкость
//
// новое значение состояния громкости.
func (y *Client) UpdateVolumeInfo(deviceID string, VolumeInfo *ynisonstate.DeviceVolume) {
	r := new(ynisonstate.PutYnisonStateRequest)
	r.Parameters = &ynisonstate.PutYnisonStateRequest_UpdateVolumeInfo{
		UpdateVolumeInfo: &ynisonstate.UpdateVolumeInfo{
			DeviceId:   deviceID,
			VolumeInfo: VolumeInfo,
		},
	}
	y.conn.Send(r)
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
func (y *Client) UpdatePlayingStatus(PlayingStatus *ynisonstate.PlayingStatus) {
	r := new(ynisonstate.PutYnisonStateRequest)
	r.Parameters = &ynisonstate.PutYnisonStateRequest_UpdatePlayingStatus{
		UpdatePlayingStatus: &ynisonstate.UpdatePlayingStatus{
			PlayingStatus: PlayingStatus,
		},
	}
	y.conn.Send(r)
}

// Обновить активное устройство.
func (y *Client) UpdateActiveDevice(deviceID string) {
	r := new(ynisonstate.PutYnisonStateRequest)
	r.Parameters = &ynisonstate.PutYnisonStateRequest_UpdateActiveDevice{
		UpdateActiveDevice: &ynisonstate.UpdateActiveDevice{
			DeviceIdOptional: wrapperspb.String(deviceID),
		},
	}
	y.conn.Send(r)
}

// Close connection
func (y *Client) Close() {
	y.conn.Close()
}

// OnMessage event
func (y *Client) OnMessage(f func(*ynisonstate.PutYnisonStateResponse)) {
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

func generateTLSCreds() (credentials.TransportCredentials, error) {
	systemRoots, err := x509.SystemCertPool()
	if err != nil {
		return nil, err
	}

	return credentials.NewTLS(&tls.Config{
		RootCAs: systemRoots,
	}), nil
}

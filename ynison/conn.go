package ynison

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
	"sync"
	"time"

	"github.com/gorilla/websocket"
)

// based on https://raw.githubusercontent.com/Adeithe/go-twitch/master/pubsub/conn.go

type Conn struct {
	socket *websocket.Conn
	done   chan bool

	isConnected bool
	writer      sync.Mutex

	onMessage    []func(AudioMessage)
	onTicket     []func(string, RedirectResponse)
	onConnect    []func()
	onDisconnect []func()
}

type IConn interface {
	Connect(string, http.Header) error
	Write(int, []byte) error
	Close()

	IsConnected() bool

	Listen(...string) error
	Unlisten(...string) error

	OnMessage(func(AudioMessage))
	OnTicket(func(string, RedirectResponse))
	OnConnect(func())
	OnDisconnect(func())
}

var _ IConn = &Conn{}

func (conn *Conn) Connect(Host string, Header http.Header) error {
	socket, _, err := websocket.DefaultDialer.Dial(Host, Header)
	if err != nil {
		return err
	}
	conn.socket = socket
	conn.done = make(chan bool)
	conn.isConnected = true
	go conn.reader()
	for _, f := range conn.onConnect {
		go f()
	}
	return nil
}

// Write a message and send it to the server
func (conn *Conn) Write(msgType int, data []byte) error {
	conn.writer.Lock()
	defer conn.writer.Unlock()
	return conn.socket.WriteMessage(msgType, data)
}

// Close the connection to server
func (conn *Conn) Close() {
	conn.Write(websocket.CloseMessage, websocket.FormatCloseMessage(websocket.CloseNormalClosure, ""))
	timer := time.NewTimer(time.Second)
	defer timer.Stop()
	select {
	case <-conn.done:
	case <-timer.C:
		conn.socket.Close()
	}
	for _, f := range conn.onDisconnect {
		go f()
	}
}

// IsConnected returns true if the socket is actively connected
func (conn *Conn) IsConnected() bool {
	return conn.isConnected
}

// Listen to a topic
//
// This operation will block, giving the server up to 5 seconds to respond after correcting for latency before failing
func (conn *Conn) Listen(topics ...string) error {
	for _, message := range topics {
		if err := conn.Write(websocket.TextMessage, []byte(message)); err != nil {
			return err
		}
	}
	return nil
}

// Unlisten from the provided topics
//
// Not implemented
func (conn *Conn) Unlisten(topics ...string) error {
	// ???
	return nil
}

// OnMessage event called after a message is receieved
func (conn *Conn) OnMessage(f func(AudioMessage)) {
	conn.onMessage = append(conn.onMessage, f)
}

// OnTicket event called after a ticket is receieved
func (conn *Conn) OnTicket(f func(string, RedirectResponse)) {
	conn.onTicket = append(conn.onTicket, f)
}

// OnConnect event called after the connection is opened
func (conn *Conn) OnConnect(f func()) {
	conn.onConnect = append(conn.onConnect, f)
}

// OnDisconnect event called after the connection is closed
func (conn *Conn) OnDisconnect(f func()) {
	conn.onDisconnect = append(conn.onDisconnect, f)
}

func (conn *Conn) reader() {
	for {
		msgType, bytes, err := conn.socket.ReadMessage()
		if err != nil || msgType == websocket.CloseMessage {
			break
		}
		conn.handleMessage(bytes)
	}
	conn.socket.Close()
	conn.isConnected = false
	close(conn.done)
	for _, f := range conn.onDisconnect {
		go f()
	}
}

func (conn *Conn) handleMessage(bytes []byte) {
	if len(bytes) < 1 {
		return
	}
	var msg1 AudioMessage
	if err := json.Unmarshal(bytes, &msg1); err != nil {
		fmt.Println(strings.TrimSpace(string(bytes)))
	}
	for _, f := range conn.onMessage {
		go f(msg1)
	}
}

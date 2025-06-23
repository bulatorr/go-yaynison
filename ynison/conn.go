package ynison

import (
	"encoding/json"
	"log"
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

	onMessage    []func(PutYnisonStateResponse)
	onTicket     []func(string, RedirectResponse)
	onConnect    []func()
	onDisconnect []func()
}

type IConn interface {
	Connect(string, http.Header) error
	Write(int, []byte) error
	Close()

	IsConnected() bool

	Send(...string) error
	Unlisten(...string) error

	OnMessage(func(PutYnisonStateResponse))
	OnTicket(func(string, RedirectResponse))
	OnConnect(func())
	OnDisconnect(func())
}

var _ IConn = &Conn{}

func (conn *Conn) Connect(Host string, Header http.Header) error {
	socket, resp, err := websocket.DefaultDialer.Dial(Host, Header)
	if resp != nil {
		resp.Body.Close()
	}
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
	if !conn.isConnected {
		return
	}
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

// Send to a topic
//
// This operation will block, giving the server up to 5 seconds to respond after correcting for latency before failing
func (conn *Conn) Send(topics ...string) error {
	for _, message := range topics {
		if err := conn.Write(websocket.TextMessage, []byte(message)); err != nil {
			return err
		}
	}
	return nil
}

// Send to a topic bytes
//
// This operation will block, giving the server up to 5 seconds to respond after correcting for latency before failing
func (conn *Conn) SendBytes(message []byte) error {
	return conn.Write(websocket.TextMessage, message)
}

// Unlisten from the provided topics
//
// Not implemented
func (conn *Conn) Unlisten(topics ...string) error {
	// ???
	return nil
}

// OnMessage event called after a message is receieved
func (conn *Conn) OnMessage(f func(PutYnisonStateResponse)) {
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
	// log.Println(strings.TrimSpace(string(bytes)))
	var msg PutYnisonStateResponse
	if err := json.Unmarshal(bytes, &msg); err != nil {
		log.Println(strings.TrimSpace(string(bytes)))
		log.Println(err.Error())
		return
	}
	for _, f := range conn.onMessage {
		go f(msg)
	}
}

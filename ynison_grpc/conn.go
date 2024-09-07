package ynisonGrpc

import (
	"context"
	"sync"

	"github.com/bulatorr/go-yaynison/ynisonstate"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
)

// based on https://raw.githubusercontent.com/Adeithe/go-twitch/master/pubsub/conn.go

type Conn struct {
	socket *grpc.ClientConn
	stream ynisonstate.YnisonStateService_PutYnisonStateClient

	done chan bool

	isConnected bool
	writer      sync.Mutex

	onMessage    []func(*ynisonstate.PutYnisonStateResponse)
	onConnect    []func()
	onDisconnect []func()
}

type IConn interface {
	Connect(Host string, Header map[string]string) error
	Close()

	IsConnected() bool

	Send(...*ynisonstate.PutYnisonStateRequest) error

	OnMessage(func(*ynisonstate.PutYnisonStateResponse))
	OnConnect(func())
	OnDisconnect(func())
}

var _ IConn = &Conn{}

func (conn *Conn) Connect(Host string, Header map[string]string) error {
	tlsCreds, _ := generateTLSCreds()
	cc, err := grpc.NewClient(Host, grpc.WithTransportCredentials(tlsCreds))
	if err != nil {
		return err
	}
	conn.socket = cc
	conn.done = make(chan bool)
	conn.isConnected = true
	md := metadata.New(Header)
	ctx := metadata.NewOutgoingContext(context.Background(), md)

	Ys := ynisonstate.NewYnisonStateServiceClient(cc)
	stream, err := Ys.PutYnisonState(ctx, grpc.Header(&md))
	if err != nil {
		return err
	}
	conn.stream = stream

	go conn.reader()
	for _, f := range conn.onConnect {
		go f()
	}
	return nil
}

// Close the connection to server
func (conn *Conn) Close() {
	conn.socket.Close()
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
func (conn *Conn) Send(topics ...*ynisonstate.PutYnisonStateRequest) error {
	conn.writer.Lock()
	defer conn.writer.Unlock()
	for _, message := range topics {
		conn.stream.Send(message)
	}
	return nil
}

// OnMessage event called after a message is receieved
func (conn *Conn) OnMessage(f func(*ynisonstate.PutYnisonStateResponse)) {
	conn.onMessage = append(conn.onMessage, f)
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
		resp, err := conn.stream.Recv()
		if err != nil {
			break
		}
		for _, f := range conn.onMessage {
			go f(resp)
		}
	}
	conn.socket.Close()
	conn.isConnected = false
	close(conn.done)
	for _, f := range conn.onDisconnect {
		go f()
	}
}

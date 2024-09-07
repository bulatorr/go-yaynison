package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"os/signal"
	"time"

	ynisonGrpc "github.com/bulatorr/go-yaynison/ynison_grpc"
	"github.com/bulatorr/go-yaynison/ynisonstate"
)

func main() {
	var token string
	flag.StringVar(&token, "t", "empty", "OAuth token")
	flag.Parse()
	if token == "empty" {
		log.Fatal("token required")
	}
	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt)

	y := ynisonGrpc.NewClient(token, "mi007pp2as935j59")
	defer y.Close()

	y.OnMessage(func(am *ynisonstate.PutYnisonStateResponse) {
		fmt.Printf("\n\n[OnMessage]\n")

		for i, device := range am.GetDevices() {
			fmt.Printf("%d %s %s\n", i, device.Info.Title, device.Info.DeviceId)
		}
		current, _ := time.ParseDuration(fmt.Sprint(am.GetPlayerState().GetStatus().GetProgressMs()) + "ms")
		total, _ := time.ParseDuration(fmt.Sprint(am.GetPlayerState().GetStatus().GetDurationMs()) + "ms")
		fmt.Println("Rid:", am.GetRid())
		fmt.Println("Pause:", am.GetPlayerState().GetStatus().GetPaused())
		playableList := am.GetPlayerState().GetPlayerQueue().GetPlayableList()
		if len(playableList) > 0 {
			data := playableList[am.GetPlayerState().GetPlayerQueue().GetCurrentPlayableIndex()]
			fmt.Println("Title:", data.GetTitle())
			fmt.Println("Type:", data.GetPlayableType())
			fmt.Println("Album:", data.GetAlbumIdOptional().GetValue())
			fmt.Println("Track:", data.GetPlayableId())
			fmt.Println("Cover url:", data.GetCoverUrlOptional().GetValue())
			fmt.Println("From:", data.GetFrom())
		}
		fmt.Printf("Played: %s of %s\n", timePrettify(current), timePrettify(total))
	})
	y.OnConnect(func() {
		log.Println("[OnConnect] connected to ynison")
	})
	y.OnDisconnect(func() {
		log.Println("[OnDisconnect] disconnected from ynison")
	})

	err := y.Connect()
	if err != nil {
		log.Fatal(err)
	}
	<-interrupt
}

func timePrettify(t time.Duration) string {
	d := t.Round(time.Millisecond)
	minute := int(d.Minutes()) % 60
	second := int(d.Seconds()) % 60
	milli := int(d.Milliseconds()) % 1000
	return fmt.Sprintf("%02d:%02d:%02d", minute, second, milli) // 09:37:228
}

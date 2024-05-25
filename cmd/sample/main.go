package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"os/signal"
	"time"

	"github.com/bulatorr/go-yaynison/ynison"
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

	y := ynison.NewClient(token)
	defer y.Close()

	y.OnMessage(func(am ynison.AudioMessage) {
		fmt.Printf("\n\n[OnMessage]\n")

		current, _ := time.ParseDuration(am.PlayerState.Status.ProgressMs + "ms")
		total, _ := time.ParseDuration(am.PlayerState.Status.DurationMs + "ms")
		fmt.Println("Rid:", am.Rid)
		fmt.Println("Pause:", am.PlayerState.Status.Paused)
		fmt.Println("Title:", am.PlayerState.PlayerQueue.PlayableList[am.PlayerState.PlayerQueue.CurrentPlayableIndex].Title)
		fmt.Println("TrackID:", am.PlayerState.PlayerQueue.PlayableList[am.PlayerState.PlayerQueue.CurrentPlayableIndex].PlayableID)
		fmt.Println("From:", am.PlayerState.PlayerQueue.PlayableList[am.PlayerState.PlayerQueue.CurrentPlayableIndex].From)
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

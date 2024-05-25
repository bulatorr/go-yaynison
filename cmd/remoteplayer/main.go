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

var ProgressMs string
var DurationMs string

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
		ProgressMs = am.PlayerState.Status.ProgressMs
		DurationMs = am.PlayerState.Status.DurationMs

		fmt.Printf("\n\n[OnMessage]\n")
		for i, device := range am.Devices {
			fmt.Printf("%d %s %s\n", i, device.Info.Title, device.Info.DeviceID)
		}
		current, _ := time.ParseDuration(am.PlayerState.Status.ProgressMs + "ms")
		total, _ := time.ParseDuration(am.PlayerState.Status.DurationMs + "ms")
		fmt.Println("Rid:", am.Rid)
		fmt.Println("Pause:", am.PlayerState.Status.Paused)
		if len(am.PlayerState.PlayerQueue.PlayableList) > 0 {
			fmt.Println("Title:", am.PlayerState.PlayerQueue.PlayableList[am.PlayerState.PlayerQueue.CurrentPlayableIndex].Title)
			fmt.Println("TrackID:", am.PlayerState.PlayerQueue.PlayableList[am.PlayerState.PlayerQueue.CurrentPlayableIndex].PlayableID)
			fmt.Println("From:", am.PlayerState.PlayerQueue.PlayableList[am.PlayerState.PlayerQueue.CurrentPlayableIndex].From)
		}
		fmt.Printf("Played: %s of %s\n", timePrettify(current), timePrettify(total))
	})
	y.OnConnect(func() {
		log.Println("[OnConnect] connected to ynison")
		time.Sleep(3 * time.Second)
		// изменение громкости. принимает значение от 0 до 1
		y.UpdateVolumeInfo("device_id here", 0.2)
		// изменение состояние плеера. пауза, перемотка.
		y.UpdatePlayingStatus(true, ProgressMs, DurationMs)
		time.Sleep(3 * time.Second)
		// cнимаем с паузы
		y.UpdatePlayingStatus(false, ProgressMs, DurationMs)
		log.Println("[onConnect] sended")
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

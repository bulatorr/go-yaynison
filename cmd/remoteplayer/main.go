package main

import (
	"flag"
	"fmt"
	"log"
	"time"

	"github.com/bulatorr/go-yaynison/ynison"
)

var ProgressMs string
var DurationMs string
var pauseState bool

func main() {
	var token string
	var deviceid string
	var paused bool
	var fetch bool
	var volume float64

	flag.StringVar(&token, "t", "", "OAuth token")
	flag.StringVar(&deviceid, "d", "", "deviceID")
	flag.BoolVar(&paused, "p", false, "Change state of Player (if not paused, it will stop the player, and vice versa)")
	flag.BoolVar(&fetch, "f", false, "Show data from ynison server")
	flag.Float64Var(&volume, "v", 228, "Volume level")
	flag.Parse()

	if token == "" {
		log.Fatal("token required")
	}

	interrupt := make(chan bool, 1)
	programexit := make(chan bool, 1)
	defer close(interrupt)
	defer close(programexit)

	y := ynison.NewClientWithDeviceID(token, "mi007pp2as935j60")
	defer y.Close()

	y.OnMessage(func(am ynison.PutYnisonStateResponse) {
		ProgressMs = am.PlayerState.Status.ProgressMs
		DurationMs = am.PlayerState.Status.DurationMs
		pauseState = am.PlayerState.Status.Paused
		deviceid = am.ActiveDeviceIDOptional
		if fetch {
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
				fmt.Println("From:", am.PlayerState.PlayerQueue.PlayableList[am.PlayerState.PlayerQueue.CurrentPlayableIndex])
			}
			fmt.Printf("Played: %s of %s\n", timePrettify(current), timePrettify(total))
		}
		interrupt <- true
	})
	y.OnConnect(func() {
		log.Println("[OnConnect] connected to ynison")
		// задержка 1-3 секунды, иначе работает через раз
		time.Sleep(2 * time.Second)
		<-interrupt
		if paused {
			// изменение состояние плеера. пауза, перемотка.
			log.Println("New state paused:", !pauseState)
			y.UpdatePlayingStatus(!pauseState, ProgressMs, DurationMs)
		}
		if volume != 228 {
			if volume >= 0 && volume <= 1 {
				// изменение громкости. принимает значение от 0 до 1
				log.Println("New state volume:", volume)
				y.UpdateVolumeInfo(deviceid, volume)
			} else {
				log.Println("Volume level should be [0, 1] Example: 0.5")
			}
		}
		log.Println("[onConnect] done")
		time.Sleep(2 * time.Second)
		programexit <- true
	})
	y.OnDisconnect(func() {
		log.Println("[OnDisconnect] disconnected from ynison")
	})
	err := y.Connect()
	if err != nil {
		log.Fatal(err)
	}
	<-programexit
}

func timePrettify(t time.Duration) string {
	d := t.Round(time.Millisecond)
	minute := int(d.Minutes()) % 60
	second := int(d.Seconds()) % 60
	milli := int(d.Milliseconds()) % 1000
	return fmt.Sprintf("%02d:%02d:%02d", minute, second, milli) // 09:37:228
}

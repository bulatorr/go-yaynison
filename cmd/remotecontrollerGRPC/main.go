package main

import (
	"flag"
	"fmt"
	"log"
	"time"

	ynisonGrpc "github.com/bulatorr/go-yaynison/ynison_grpc"
	"github.com/bulatorr/go-yaynison/ynisonstate"
)

var ProgressMs int64
var DurationMs int64
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

	y := ynisonGrpc.NewClient(token, "mi007pp2as935j60")
	defer y.Close()

	y.OnMessage(func(am *ynisonstate.PutYnisonStateResponse) {
		ProgressMs = am.PlayerState.Status.ProgressMs
		DurationMs = am.PlayerState.Status.DurationMs
		pauseState = am.PlayerState.Status.Paused
		if deviceid == "" {
			deviceid = am.ActiveDeviceIdOptional.Value
		}
		if fetch {
			fmt.Printf("\n\n[OnMessage]\n")
			for i, device := range am.Devices {
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
		}
		interrupt <- true
	})
	y.OnConnect(func() {
		log.Println("[OnConnect] connected to ynison")
		<-interrupt
		if paused {
			// изменение состояние плеера. пауза, перемотка.
			log.Println("New state paused:", !pauseState)
			y.UpdatePlayingStatus(&ynisonstate.PlayingStatus{
				ProgressMs:    ProgressMs,
				DurationMs:    DurationMs,
				Paused:        !pauseState,
				PlaybackSpeed: 1,
			})
		}
		if volume != 228 {
			if volume >= 0 && volume <= 1 {
				// изменение громкости. принимает значение от 0 до 1
				log.Println("New state volume:", volume)
				y.UpdateVolumeInfo(deviceid, &ynisonstate.DeviceVolume{
					Volume: volume,
				})
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

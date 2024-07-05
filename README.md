# go-YaYnison 🎵

A Go library for interacting with Yandex Ynison.

## Build

1. Clone the repository:
    >git clone https://github.com/bulatorr/go-yaynison
   

2. Navigate to the sample directory:
    >cd cmd/sample
   

3. Run the build script:
    >./build.sh
   

## Usage

To use the sample, run the following command in your terminal:

>sample -t \<token\>


## Output

The program will output messages like below:

2024/05/25 11:09:51 [OnConnect] connected to ynison

[OnMessage]
Rid: 493763ef-2d95-4cac-8db6-940c6469ab6d
Pause: true
Title: Cornuto
TrackID: 82330963
From: web-album_track-track-track-saved
Played: 00:22:613 of 02:04:760

[OnMessage]
Rid: 694adf04-44de-4682-bb9b-2a395f79d2d2
Pause: false
Title: Cornuto
TrackID: 82330963
From: web-album_track-track-track-saved
Played: 00:22:619 of 02:04:760

...

2024/05/25 11:10:27 [OnDisconnect] disconnected from ynison

## RemoteController

Использование RemoteController:

- -d string
   deviceID
- -f
   Показать данные с сервера Ynison
- -p
   Изменить состояние плеера (если не приостановлен, он остановит воспроизведение плеера, и наоборот)
- -t string
   Токен OAuth
- -v float
   Уровень громкости [0,1]

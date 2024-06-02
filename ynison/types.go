package ynison

type RedirectResponse struct {
	Host            string `json:"host"`
	RedirectTicket  string `json:"redirect_ticket"`
	SessionID       string `json:"session_id"`
	KeepAliveParams struct {
		KeepAliveTimeSeconds    int `json:"keep_alive_time_seconds"`
		KeepAliveTimeoutSeconds int `json:"keep_alive_timeout_seconds"`
	} `json:"keep_alive_params"`
}

// Сообщение с обновленным состоянием.
//
// Сервер отправляет на клиент в следующих случаях:
// * В ответ на `PutYnisonStateRequest`, который меняет состояние на сервере.
// * Если произошло обновление на другом клиенте и его нужно транслировать всем подключенным устройствам
// (например, сменилась очередь воспроизведения).
// * Если изменился список устройств.
type PutYnisonStateResponse struct {
	PlayerState struct {
		Status struct {
			ProgressMs    string  `json:"progress_ms,omitempty"`
			DurationMs    string  `json:"duration_ms,omitempty"`
			Paused        bool    `json:"paused,omitempty"`
			PlaybackSpeed float64 `json:"playback_speed,omitempty"`
			Version       struct {
				DeviceID    string `json:"device_id,omitempty"`
				Version     string `json:"version,omitempty"`
				TimestampMs string `json:"timestamp_ms,omitempty"`
			} `json:"version,omitempty"`
		} `json:"status,omitempty"`
		PlayerQueue struct {
			EntityID   string `json:"entity_id,omitempty"`
			EntityType string `json:"entity_type,omitempty"`
			Queue      struct {
				WaveQueue struct {
					RecommendedPlayableList []struct {
						PlayableID   string `json:"playable_id,omitempty"`
						PlayableType string `json:"playable_type,omitempty"`
						Title        string `json:"title,omitempty"`
						CoverURL     struct {
							Value string `json:"value,omitempty"`
						} `json:"cover_url,omitempty"`
					} `json:"recommended_playable_list,omitempty"`
					LivePlayableIndex int `json:"live_playable_index,omitempty"`
					EntityOptions     struct {
						WaveEntityOptional struct {
							SessionID string `json:"session_id,omitempty"`
						} `json:"wave_entity_optional,omitempty"`
						TrackSources []struct {
							Key        int `json:"key,omitempty"`
							WaveSource struct {
							} `json:"wave_source,omitempty"`
							PhonotekaSource struct {
								ArtistID struct {
									ID string `json:"id,omitempty"`
								} `json:"artist_id,omitempty"`
								PlaylistID struct {
									ID string `json:"id,omitempty"`
								} `json:"playlist_id,omitempty"`
								AlbumID struct {
									ID string `json:"id,omitempty"`
								} `json:"album_id,omitempty"`
								EntityContext string `json:"entity_context,omitempty"`
							} `json:"phonoteka_source,omitempty"`
						} `json:"track_sources,omitempty"`
					} `json:"entity_options,omitempty"`
				} `json:"wave_queue,omitempty"`
				GenerativeQueue struct {
					ID string `json:"id,omitempty"`
				} `json:"generative_queue,omitempty"`
				FmRadioQueue struct {
					ID string `json:"id,omitempty"`
				} `json:"fm_radio_queue,omitempty"`
				VideoWaveQueue struct {
					ID string `json:"id,omitempty"`
				} `json:"video_wave_queue,omitempty"`
				LocalTracksQueue struct {
				} `json:"local_tracks_queue,omitempty"`
			} `json:"queue,omitempty"`
			CurrentPlayableIndex int `json:"current_playable_index,omitempty"`
			PlayableList         []struct {
				PlayableID   string `json:"playable_id,omitempty"`
				PlayableType string `json:"playable_type,omitempty"`
				From         string `json:"from"`
				Title        string `json:"title,omitempty"`
				CoverURL     string `json:"cover_url,omitempty"`
			} `json:"playable_list,omitempty"`
			Options struct {
				RepeatMode string `json:"repeat_mode,omitempty"`
			} `json:"options,omitempty"`
			Version struct {
				DeviceID    string `json:"device_id,omitempty"`
				Version     string `json:"version,omitempty"`
				TimestampMs string `json:"timestamp_ms,omitempty"`
			} `json:"version,omitempty"`
			ShuffleOptional struct {
				PlayableIndices []int `json:"playable_indices,omitempty"`
			} `json:"shuffle_optional,omitempty"`
			EntityContext         string `json:"entity_context,omitempty"`
			FromOptional          string `json:"from_optional,omitempty"`
			InitialEntityOptional struct {
				EntityID   string `json:"entity_id,omitempty"`
				EntityType string `json:"entity_type,omitempty"`
			} `json:"initial_entity_optional,omitempty"`
			AddingOptionsOptional struct {
				RadioOptions struct {
					SessionID string `json:"session_id,omitempty"`
				} `json:"radio_options,omitempty"`
			} `json:"adding_options_optional,omitempty"`
		} `json:"player_queue,omitempty"`
		PlayerQueueInjectOptional struct {
			PlayingStatus struct {
				ProgressMs    string  `json:"progress_ms,omitempty"`
				DurationMs    string  `json:"duration_ms,omitempty"`
				Paused        bool    `json:"paused,omitempty"`
				PlaybackSpeed float64 `json:"playback_speed,omitempty"`
				Version       struct {
					DeviceID    string `json:"device_id,omitempty"`
					Version     string `json:"version,omitempty"`
					TimestampMs string `json:"timestamp_ms,omitempty"`
				} `json:"version,omitempty"`
			} `json:"playing_status,omitempty"`
			Playable struct {
				PlayableID   string `json:"playable_id,omitempty"`
				PlayableType string `json:"playable_type,omitempty"`
				Title        string `json:"title,omitempty"`
				CoverURL     struct {
					Value string `json:"value,omitempty"`
				} `json:"cover_url,omitempty"`
			} `json:"playable,omitempty"`
			Version struct {
				DeviceID    string `json:"device_id,omitempty"`
				Version     string `json:"version,omitempty"`
				TimestampMs string `json:"timestamp_ms,omitempty"`
			} `json:"version,omitempty"`
		} `json:"player_queue_inject_optional,omitempty"`
	} `json:"player_state,omitempty"`
	Devices []struct {
		Info struct {
			DeviceID   string `json:"device_id,omitempty"`
			Title      string `json:"title,omitempty"`
			Type       string `json:"type,omitempty"`
			AppName    string `json:"app_name,omitempty"`
			AppVersion string `json:"app_version,omitempty"`
		} `json:"info,omitempty"`
		Volume       float64 `json:"volume,omitempty"`
		Capabilities struct {
			CanBePlayer           bool `json:"can_be_player,omitempty"`
			CanBeRemoteController bool `json:"can_be_remote_controller,omitempty"`
			VolumeGranularity     int  `json:"volume_granularity,omitempty"`
		} `json:"capabilities,omitempty"`
		Session struct {
			ID string `json:"id,omitempty"`
		} `json:"session,omitempty"`
		IsOffline  bool `json:"is_offline,omitempty"`
		VolumeInfo struct {
			Volume  float64 `json:"volume,omitempty"`
			Version struct {
				DeviceID    string `json:"device_id,omitempty"`
				Version     string `json:"version,omitempty"`
				TimestampMs string `json:"timestamp_ms,omitempty"`
			} `json:"version,omitempty"`
		} `json:"volume_info,omitempty"`
	} `json:"devices,omitempty"`
	ActiveDeviceIDOptional string `json:"active_device_id_optional,omitempty"`
	TimestampMs            string `json:"timestamp_ms,omitempty"`
	Rid                    string `json:"rid,omitempty"`
}

// Сообщение с обновленным состоянием.
// Клиент отправляет на сервер один из параметров в зависимости
// от произошедшего события.
type PutYnisonStateRequest struct {
	UpdateFullState          *UpdateFullState         `json:"update_full_state,omitempty"`
	UpdateActiveDevice       *UpdateActiveDevice      `json:"update_active_device,omitempty"`
	UpdatePlayingStatus      *UpdatePlayingStatus     `json:"update_playing_status,omitempty"`
	UpdatePlayerState        *UpdatePlayerState       `json:"update_player_state,omitempty"`
	UpdateVolume             *UpdateVolume            `json:"update_volume,omitempty"`
	UpdatePlayerQueueInject  *UpdatePlayerQueueInject `json:"update_player_queue_inject,omitempty"`
	UpdateSessionParams      *UpdateSessionParams     `json:"update_session_params,omitempty"`
	UpdateVolumeInfo         *UpdateVolumeInfo        `json:"update_volume_info,omitempty"`
	SyncStateFromEov         *SyncStateFromEov        `json:"sync_state_from_eov,omitempty"`
	PlayerActionTimestampMs  string                   `json:"player_action_timestamp_ms,omitempty"`
	Rid                      string                   `json:"rid,omitempty"`
	ActivityInterceptionType string                   `json:"activity_interception_type,omitempty"`
}

// Версия изменений.
type Version struct {
	// Идентификатор устройства, которое инициировало изменение.
	DeviceID string `json:"device_id,omitempty"`
	// Версия последнего изменения.
	// random(int64)
	Version string `json:"version,omitempty"`
	// Время последнего изменения.
	// Диагностическое значение, не используется в бизнес-логике на клиентах.
	TimestampMs string `json:"timestamp_ms,omitempty"`
}
type Status struct {
	ProgressMs    string  `json:"progress_ms,omitempty"`
	DurationMs    string  `json:"duration_ms,omitempty"`
	Paused        bool    `json:"paused,omitempty"`
	PlaybackSpeed int     `json:"playback_speed,omitempty"`
	Version       Version `json:"version,omitempty"`
}
type CoverURL struct {
	Value string `json:"value,omitempty"`
}
type RecommendedPlayableList struct {
	PlayableID   string   `json:"playable_id,omitempty"`
	PlayableType string   `json:"playable_type,omitempty"`
	Title        string   `json:"title,omitempty"`
	CoverURL     CoverURL `json:"cover_url,omitempty"`
}
type WaveEntityOptional struct {
	SessionID string `json:"session_id,omitempty"`
}
type WaveSource struct {
}
type ArtistID struct {
	ID string `json:"id,omitempty"`
}
type PlaylistID struct {
	ID string `json:"id,omitempty"`
}
type AlbumID struct {
	ID string `json:"id,omitempty"`
}
type PhonotekaSource struct {
	ArtistID      ArtistID   `json:"artist_id,omitempty"`
	PlaylistID    PlaylistID `json:"playlist_id,omitempty"`
	AlbumID       AlbumID    `json:"album_id,omitempty"`
	EntityContext string     `json:"entity_context,omitempty"`
}
type TrackSources struct {
	Key             int             `json:"key,omitempty"`
	WaveSource      WaveSource      `json:"wave_source,omitempty"`
	PhonotekaSource PhonotekaSource `json:"phonoteka_source,omitempty"`
}
type EntityOptions struct {
	WaveEntityOptional WaveEntityOptional `json:"wave_entity_optional,omitempty"`
	TrackSources       []TrackSources     `json:"track_sources,omitempty"`
}
type WaveQueue struct {
	RecommendedPlayableList []RecommendedPlayableList `json:"recommended_playable_list,omitempty"`
	LivePlayableIndex       int                       `json:"live_playable_index,omitempty"`
	EntityOptions           EntityOptions             `json:"entity_options,omitempty"`
}
type GenerativeQueue struct {
	ID string `json:"id,omitempty"`
}
type FmRadioQueue struct {
	ID string `json:"id,omitempty"`
}
type VideoWaveQueue struct {
	ID string `json:"id,omitempty"`
}
type LocalTracksQueue struct {
}
type Queue struct {
	WaveQueue        WaveQueue        `json:"wave_queue,omitempty"`
	GenerativeQueue  GenerativeQueue  `json:"generative_queue,omitempty"`
	FmRadioQueue     FmRadioQueue     `json:"fm_radio_queue,omitempty"`
	VideoWaveQueue   VideoWaveQueue   `json:"video_wave_queue,omitempty"`
	LocalTracksQueue LocalTracksQueue `json:"local_tracks_queue,omitempty"`
}
type PlayableList struct {
	PlayableID   string   `json:"playable_id,omitempty"`
	PlayableType string   `json:"playable_type,omitempty"`
	Title        string   `json:"title,omitempty"`
	CoverURL     CoverURL `json:"cover_url,omitempty"`
}
type Options struct {
	RepeatMode string `json:"repeat_mode,omitempty"`
}
type ShuffleOptional struct {
	PlayableIndices []int `json:"playable_indices,omitempty"`
}
type FromOptional struct {
	Value string `json:"value,omitempty"`
}
type InitialEntityOptional struct {
	EntityID   string `json:"entity_id,omitempty"`
	EntityType string `json:"entity_type,omitempty"`
}
type RadioOptions struct {
	SessionID string `json:"session_id,omitempty"`
}
type AddingOptionsOptional struct {
	RadioOptions RadioOptions `json:"radio_options,omitempty"`
}
type PlayerQueue struct {
	EntityID              string                `json:"entity_id,omitempty"`
	EntityType            string                `json:"entity_type,omitempty"`
	Queue                 Queue                 `json:"queue,omitempty"`
	CurrentPlayableIndex  int                   `json:"current_playable_index,omitempty"`
	PlayableList          []PlayableList        `json:"playable_list,omitempty"`
	Options               Options               `json:"options,omitempty"`
	Version               Version               `json:"version,omitempty"`
	ShuffleOptional       ShuffleOptional       `json:"shuffle_optional,omitempty"`
	EntityContext         string                `json:"entity_context,omitempty"`
	FromOptional          FromOptional          `json:"from_optional,omitempty"`
	InitialEntityOptional InitialEntityOptional `json:"initial_entity_optional,omitempty"`
	AddingOptionsOptional AddingOptionsOptional `json:"adding_options_optional,omitempty"`
}
type PlayingStatus struct {
	ProgressMs    string  `json:"progress_ms,omitempty"`
	DurationMs    string  `json:"duration_ms,omitempty"`
	Paused        bool    `json:"paused,omitempty"`
	PlaybackSpeed int     `json:"playback_speed,omitempty"`
	Version       Version `json:"version,omitempty"`
}
type Playable struct {
	PlayableID   string   `json:"playable_id,omitempty"`
	PlayableType string   `json:"playable_type,omitempty"`
	Title        string   `json:"title,omitempty"`
	CoverURL     CoverURL `json:"cover_url,omitempty"`
}
type PlayerQueueInjectOptional struct {
	PlayingStatus PlayingStatus `json:"playing_status,omitempty"`
	Playable      Playable      `json:"playable,omitempty"`
	Version       Version       `json:"version,omitempty"`
}

// Состояние плеера.
type PlayerState struct {
	Status                    Status                    `json:"status,omitempty"`
	PlayerQueue               PlayerQueue               `json:"player_queue,omitempty"`
	PlayerQueueInjectOptional PlayerQueueInjectOptional `json:"player_queue_inject_optional,omitempty"`
}
type Info struct {
	DeviceID   string `json:"device_id,omitempty"`
	Title      string `json:"title,omitempty"`
	Type       string `json:"type,omitempty"`
	AppName    string `json:"app_name,omitempty"`
	AppVersion string `json:"app_version,omitempty"`
}
type Capabilities struct {
	CanBePlayer           bool `json:"can_be_player,omitempty"`
	CanBeRemoteController bool `json:"can_be_remote_controller,omitempty"`
	VolumeGranularity     int  `json:"volume_granularity,omitempty"`
}
type VolumeInfo struct {
	Volume  float64 `json:"volume,omitempty"`
	Version Version `json:"version,omitempty"`
}
type Device struct {
	Info         Info         `json:"info,omitempty"`
	Volume       int          `json:"volume,omitempty"`
	Capabilities Capabilities `json:"capabilities,omitempty"`
	VolumeInfo   VolumeInfo   `json:"volume_info,omitempty"`
}
type SyncStateFromEovOptional struct {
	ActualQueueID string `json:"actual_queue_id,omitempty"`
}
type PlayerQueueInject struct {
	PlayingStatus PlayingStatus `json:"playing_status,omitempty"`
	Playable      Playable      `json:"playable,omitempty"`
	Version       Version       `json:"version,omitempty"`
}

// Обновить общее состояние.
type UpdateFullState struct {
	PlayerState              PlayerState              `json:"player_state,omitempty"`
	IsCurrentlyActive        bool                     `json:"is_currently_active,omitempty"`
	Device                   Device                   `json:"device,omitempty"`
	SyncStateFromEovOptional SyncStateFromEovOptional `json:"sync_state_from_eov_optional,omitempty"`
}

// Обновить активное устройство.
type UpdateActiveDevice struct {
	DeviceIDOptional string `json:"device_id_optional,omitempty"`
}

// Обновить статус воспроизведения.
type UpdatePlayingStatus struct {
	PlayingStatus PlayingStatus `json:"playing_status,omitempty"`
}

// Обновить состояние плеера.
type UpdatePlayerState struct {
	PlayerState PlayerState `json:"player_state,omitempty"`
}

// Обновить громкость. Устаревшее значение, смотри [update_volume_info].
type UpdateVolume struct {
	Volume   float64 `json:"volume,omitempty"`
	DeviceID string  `json:"device_id,omitempty"`
}

// Обновить состояние проигрывания ижектируемой в очередь сущности.
type UpdatePlayerQueueInject struct {
	PlayerQueueInject PlayerQueueInject `json:"player_queue_inject,omitempty"`
}

// Обновить информацию об устройстве отправившем информацию.
type UpdateSessionParams struct {
	MuteEventsIfPassive bool `json:"mute_events_if_passive,omitempty"`
}

// Обновить громкость.
type UpdateVolumeInfo struct {
	DeviceID   string     `json:"device_id,omitempty"`
	VolumeInfo VolumeInfo `json:"volume_info,omitempty"`
}

// Запросить синхронизацию с сервисом ЕОВ.
type SyncStateFromEov struct {
	ActualQueueID string `json:"actual_queue_id,omitempty"`
}

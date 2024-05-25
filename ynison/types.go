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

type AudioMessage struct {
	PlayerState struct {
		Status struct {
			ProgressMs    string  `json:"progress_ms"`
			DurationMs    string  `json:"duration_ms"`
			Paused        bool    `json:"paused"`
			PlaybackSpeed float64 `json:"playback_speed"`
			Version       struct {
				DeviceID    string `json:"device_id"`
				Version     string `json:"version"`
				TimestampMs string `json:"timestamp_ms"`
			} `json:"version"`
		} `json:"status"`
		PlayerQueue struct {
			EntityID             string `json:"entity_id"`
			EntityType           string `json:"entity_type"`
			CurrentPlayableIndex int    `json:"current_playable_index"`
			PlayableList         []struct {
				PlayableID       string `json:"playable_id"`
				AlbumIDOptional  string `json:"album_id_optional,omitempty"`
				PlayableType     string `json:"playable_type"`
				From             string `json:"from"`
				Title            string `json:"title"`
				CoverURLOptional string `json:"cover_url_optional,omitempty"`
				TrackInfo        struct {
					TrackSourceKey int `json:"track_source_key"`
				} `json:"track_info"`
			} `json:"playable_list"`
			Options struct {
				RepeatMode string `json:"repeat_mode"`
			} `json:"options"`
			Version struct {
				DeviceID    string `json:"device_id"`
				Version     string `json:"version"`
				TimestampMs string `json:"timestamp_ms"`
			} `json:"version"`
			ShuffleOptional struct {
				PlayableIndices []int `json:"playable_indices"`
			} `json:"shuffle_optional"`
			EntityContext string `json:"entity_context"`
			Queue         struct {
				WaveQueue struct {
					RecommendedPlayableList []struct {
						PlayableID       string `json:"playable_id"`
						AlbumIDOptional  string `json:"album_id_optional"`
						PlayableType     string `json:"playable_type"`
						From             string `json:"from"`
						Title            string `json:"title"`
						CoverURLOptional string `json:"cover_url_optional"`
						TrackInfo        struct {
							TrackSourceKey int `json:"track_source_key"`
						} `json:"track_info"`
					} `json:"recommended_playable_list"`
					LivePlayableIndex int `json:"live_playable_index"`
					EntityOptions     struct {
						TrackSources []struct {
							Key             int `json:"key"`
							PhonotekaSource struct {
								EntityContext string `json:"entity_context"`
								PlaylistID    struct {
									ID string `json:"id"`
								} `json:"playlist_id"`
							} `json:"phonoteka_source"`
						} `json:"track_sources"`
					} `json:"entity_options"`
				} `json:"wave_queue"`
			} `json:"queue"`
		} `json:"player_queue"`
	} `json:"player_state"`
	Devices []struct {
		Info struct {
			DeviceID   string `json:"device_id"`
			Title      string `json:"title"`
			Type       string `json:"type"`
			AppName    string `json:"app_name"`
			AppVersion string `json:"app_version"`
		} `json:"info"`
		Volume       float64 `json:"volume"`
		Capabilities struct {
			CanBePlayer           bool `json:"can_be_player"`
			CanBeRemoteController bool `json:"can_be_remote_controller"`
			VolumeGranularity     int  `json:"volume_granularity"`
		} `json:"capabilities"`
		Session struct {
			ID string `json:"id"`
		} `json:"session"`
		IsOffline  bool `json:"is_offline"`
		VolumeInfo struct {
			Volume  float64 `json:"volume"`
			Version struct {
				DeviceID    string `json:"device_id"`
				Version     string `json:"version"`
				TimestampMs string `json:"timestamp_ms"`
			} `json:"version"`
		} `json:"volume_info"`
	} `json:"devices"`
	ActiveDeviceIDOptional string `json:"active_device_id_optional"`
	TimestampMs            string `json:"timestamp_ms"`
	Rid                    string `json:"rid"`
}

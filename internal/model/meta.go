package model

type AudioMetaData struct {
	Codec    string `json:"codec"`
	Bitrate  string `json:"bitrate"`
	Channels int    `json:"channels,omitempty"`
	Language string `json:"language"`
}

type VideoMetaData struct {
	Codec      string `json:"codec"`
	Bitrate    string `json:"bitrate"`
	Resolution string `json:"resolution"`
}

type MediaMetaData struct {
	Audios []AudioMetaData `json:"audios,omitempty"`
	Videos []VideoMetaData `json:"videos,omitempty"`
}

type Result struct {
	Periods []MediaMetaData `json:"periods,omitempty"`
}

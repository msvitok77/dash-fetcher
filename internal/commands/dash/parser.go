package dash

import (
	"encoding/json"
	"fmt"
	"strconv"
	"strings"

	"github.com/msvitok77/dash-fetcher/internal/model"
	"github.com/msvitok77/dash-fetcher/internal/urls"
	"github.com/zencoder/go-dash/v3/mpd"
)

const (
	audioRepresentationType = "audio"
	videoRepresentationType = "video"
)

// Parse parses the given url pointing to a .mpd source
func Parse(url string) error {
	reader, err := urls.URLResourceReader(url)
	if err != nil {
		return fmt.Errorf("parsing .mpd: %w", err)
	}
	defer reader.Close()

	// reading mpd source
	mpd, err := mpd.Read(reader)
	if err != nil {
		return fmt.Errorf("reading .mpd resource: %w", err)
	}

	// result which is then serialized to JSON
	result := model.Result{
		Periods: make([]model.MediaMetaData, 0, len(mpd.Periods)),
	}

	var mPeriod model.MediaMetaData
	for _, period := range mpd.Periods {
		mPeriod, err = parsePeriod(period)
		if err != nil {
			return err
		}
		result.Periods = append(result.Periods, mPeriod)
	}

	prettyJSON, err := json.MarshalIndent(result, "", "  ")
	if err != nil {
		return fmt.Errorf("marshaling JSON output: %w", err)
	}

	fmt.Printf("💯 JSON:\n%s", prettyJSON)
	return nil
}

// parsePeriod parses the mpd.Period to a model.MediaMetaData
func parsePeriod(period *mpd.Period) (model.MediaMetaData, error) {
	if period == nil {
		return model.MediaMetaData{}, nil
	}

	var (
		audios []model.AudioMetaData
		videos []model.VideoMetaData
	)

	for _, adaptationSet := range period.AdaptationSets {
		pAudios, pVideos, err := parseAdaptationSet(adaptationSet)
		if err != nil {
			return model.MediaMetaData{}, fmt.Errorf("parsing period: %w", err)
		}
		audios = append(audios, pAudios...)
		videos = append(videos, pVideos...)
	}

	return model.MediaMetaData{
		Audios: audios,
		Videos: videos,
	}, nil
}

// parseAdaptationSet parses the mpd.AdaptationSet to bothe model.AudioMetaData and model.VideoMetaData
func parseAdaptationSet(adaptationSet *mpd.AdaptationSet) ([]model.AudioMetaData, []model.VideoMetaData, error) {
	var (
		audios []model.AudioMetaData
		videos []model.VideoMetaData
	)

	if adaptationSet == nil {
		return audios, videos, nil
	}

	audioLang := value(adaptationSet.Lang, "")

	for _, representation := range adaptationSet.Representations {
		if representation.MimeType == nil || *representation.MimeType == "" {
			fmt.Println("unknown mimetype, skipping")
			continue
		}

		mimeType := strings.SplitN(*representation.MimeType, "/", 2)

		switch mimeType[0] {
		case audioRepresentationType:
			audio, err := parseAudio(representation, audioLang)
			if err != nil {
				return nil, nil, fmt.Errorf("parsing audio: %w", err)
			}
			audios = append(audios, audio)
		case videoRepresentationType:
			video, err := parseVideo(representation)
			if err != nil {
				return nil, nil, fmt.Errorf("parsing video: %w", err)
			}
			videos = append(videos, video)
		default:
			fmt.Println("unknown mimetype, skipping")
		}
	}

	return audios, videos, nil
}

// parseAudio parses the audio metadata to model.AudioMetaData
func parseAudio(representation *mpd.Representation, lang string) (model.AudioMetaData, error) {
	var (
		channels int
		err      error
	)

	if representation == nil {
		return model.AudioMetaData{}, nil
	}

	if representation.AudioChannelConfiguration != nil && representation.AudioChannelConfiguration.Value != nil {
		channels, err = strconv.Atoi(*representation.AudioChannelConfiguration.Value)
		if err != nil {
			return model.AudioMetaData{}, fmt.Errorf("parsing channel: %w", err)
		}
	}

	return model.AudioMetaData{
		Codec:    value(representation.Codecs, ""),
		Bitrate:  strconv.FormatInt(value(representation.Bandwidth, 0), 10),
		Channels: channels,
		Language: lang,
	}, nil
}

// parseAudio parses the video metadata to model.VideoMetaData
func parseVideo(representation *mpd.Representation) (model.VideoMetaData, error) {
	var (
		width  = value(representation.Width, 0)
		height = value(representation.Height, 0)

		resolution string
	)

	if width != 0 && height != 0 {
		resolution = fmt.Sprintf("%dx%d", width, height)
	}

	return model.VideoMetaData{
		Codec:      value(representation.Codecs, ""),
		Bitrate:    strconv.FormatInt(value(representation.Bandwidth, 0), 10),
		Resolution: resolution,
	}, nil
}

func value[T any](value *T, defaultValue T) T {
	if value == nil {
		return defaultValue
	}
	return *value
}

package vlc

// #cgo LDFLAGS: -lvlc
// #include <vlc/vlc.h>
import "C"
import (
	"unsafe"
)

// MediaTrackType represents the type of a media track.
type MediaTrackType int

// Media track types.
const (
	MediaTrackUnknown MediaTrackType = iota - 1
	MediaTrackAudio
	MediaTrackVideo
	MediaTrackText
)

// MediaAudioTrack contains information specific to audio media tracks.
type MediaAudioTrack struct {
	Channels uint // number of audio channels.
	Rate     uint // audio sample rate.
}

// MediaVideoTrack contains information specific to video media tracks.
type MediaVideoTrack struct {
	Width  uint // video width.
	Height uint // video height.

	// Aspect ratio information.
	AspectRatioNum uint // aspect ratio numerator.
	AspectRatioDen uint // aspect ratio denominator.

	// Frame rate information.
	FrameRateNum uint // frame rate numerator.
	FrameRateDen uint // frame rate denominator.
}

// MediaSubtitleTrack contains information specific to subtitle media tracks.
type MediaSubtitleTrack struct {
	Encoding string // character encoding of the subtitle.
}

// MediaTrack contains information regarding a media track.
type MediaTrack struct {
	ID      int            // Media track identifier.
	Type    MediaTrackType // Media track type.
	BitRate uint           // Media track bit rate.

	// libVLC representation of the four-character code of the codec used by
	// the media track.
	Codec uint

	// The original four-character code of the codec used by the media track,
	// extracted from the container.
	OriginalCodec uint

	// Codec profile (real audio flavor, MPEG audio layer, H264 profile, etc.).
	// NOTE: Codec specific.
	Profile int

	// Stream restriction level (resolution, bitrate, codec features, etc.).
	// NOTE: Codec specific.
	Level int

	Language    string // Media track language name.
	Description string // Description of the media track.

	// Type specific information.
	Audio    *MediaAudioTrack
	Video    *MediaVideoTrack
	Subtitle *MediaSubtitleTrack
}

func (mt *MediaTrack) assertInit() error {
	if mt == nil {
		return ErrMediaTrackNotInitialized
	}

	return nil
}

func parseMediaTrack(cTrack *C.libvlc_media_track_t) (*MediaTrack, error) {
	if cTrack == nil {
		return nil, ErrMediaTrackNotInitialized
	}

	mt := &MediaTrack{
		ID:            int(cTrack.i_id),
		Type:          MediaTrackType(cTrack.i_type),
		BitRate:       uint(cTrack.i_bitrate),
		Codec:         uint(cTrack.i_codec),
		OriginalCodec: uint(cTrack.i_original_fourcc),
		Profile:       int(cTrack.i_profile),
		Level:         int(cTrack.i_level),
		Language:      C.GoString(cTrack.psz_language),
		Description:   C.GoString(cTrack.psz_description),
	}

	switch mt.Type {
	case MediaTrackAudio:
		audio := *(**C.libvlc_audio_track_t)(unsafe.Pointer(&cTrack.anon0[0]))
		if audio == nil {
			break
		}

		mt.Audio = &MediaAudioTrack{
			Channels: uint(audio.i_channels),
			Rate:     uint(audio.i_rate),
		}
	case MediaTrackVideo:
		video := *(**C.libvlc_video_track_t)(unsafe.Pointer(&cTrack.anon0[0]))
		if video == nil {
			break
		}

		mt.Video = &MediaVideoTrack{
			Width:          uint(video.i_width),
			Height:         uint(video.i_height),
			AspectRatioNum: uint(video.i_sar_num),
			AspectRatioDen: uint(video.i_sar_den),
			FrameRateNum:   uint(video.i_frame_rate_num),
			FrameRateDen:   uint(video.i_frame_rate_den),
		}
	case MediaTrackText:
		subtitle := *(**C.libvlc_subtitle_track_t)(unsafe.Pointer(&cTrack.anon0[0]))
		if subtitle == nil {
			break
		}

		mt.Subtitle = &MediaSubtitleTrack{
			Encoding: C.GoString(subtitle.psz_encoding),
		}
	}

	return mt, nil
}

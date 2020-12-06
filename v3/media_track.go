package vlc

// #cgo LDFLAGS: -lvlc
// #include <vlc/vlc.h>
import "C"
import (
	"unsafe"
)

// VideoOrientation represents the orientation of a video media track.
type VideoOrientation int

// Video orientations.
const (
	// Normal.
	OrientationTopLeft VideoOrientation = iota

	// Flipped horizontally.
	OrientationTopRight

	// Flipped vertically.
	OrientationBottomLeft

	// Rotated 180 degrees.
	OrientationBottomRight

	// Transposed.
	OrientationLeftTop

	// Rotated 90 degrees anti-clockwise.
	OrientationLeftBottom

	// Rotated 90 degrees clockwise.
	OrientationRightTop

	// Anti-transposed.
	OrientationRightBottom
)

// VideoProjection represents the projection mode of a video media track.
type VideoProjection int

// Video projections.
const (
	ProjectionRectangular           VideoProjection = 0
	ProjectionEquirectangular                       = 1
	ProjectionCubemapLayoutStandard                 = 0x100
)

// VideoViewpoint contains viewpoint information for a video media track.
type VideoViewpoint struct {
	Yaw   float64 // Viewpoint yaw in degrees [-180-180].
	Pitch float64 // Viewpoint pitch in degrees [-90-90].
	Roll  float64 // Viewpoint roll in degrees [-180-180].
	FOV   float64 // Viewpoint field of view in degrees [0-180]. Default: 80.
}

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
	Width       uint             // video width.
	Height      uint             // video height.
	Orientation VideoOrientation // video orientation.
	Projection  VideoProjection  // video projection mode.
	Pose        VideoViewpoint   // video initial viewpoint.

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

// CodecDescription returns the description of the codec used by the media track.
func (mt *MediaTrack) CodecDescription() (string, error) {
	if err := mt.assertInit(); err != nil {
		return "", err
	}

	codec := mt.Codec
	if codec == 0 {
		codec = mt.OriginalCodec
	}

	// Get codec description.
	return C.GoString(C.libvlc_media_get_codec_description(
		C.libvlc_track_type_t(mt.Type),
		C.uint(codec),
	)), nil
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
			Orientation:    VideoOrientation(video.i_orientation),
			Projection:     VideoProjection(video.i_projection),
			Pose: VideoViewpoint{
				Yaw:   float64(video.pose.f_yaw),
				Pitch: float64(video.pose.f_pitch),
				Roll:  float64(video.pose.f_roll),
				FOV:   float64(video.pose.f_field_of_view),
			},
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

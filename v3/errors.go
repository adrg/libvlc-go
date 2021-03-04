package vlc

import "errors"

// Module errors.
var (
	ErrModuleInitialize     = errors.New("could not initialize module")
	ErrModuleNotInitialized = errors.New("module not initialized")
	ErrUserInterfaceStart   = errors.New("could not start user interface")
)

// Player errors.
var (
	ErrPlayerCreate             = errors.New("could not create player")
	ErrPlayerNotInitialized     = errors.New("player not initialized")
	ErrPlayerSetRenderer        = errors.New("could not set player renderer")
	ErrPlayerSetEqualizer       = errors.New("could not set player equalizer")
	ErrListPlayerCreate         = errors.New("could not create list player")
	ErrListPlayerNotInitialized = errors.New("list player not initialized")
)

// Media errors.
var (
	ErrMediaCreate             = errors.New("could not create media")
	ErrMediaNotFound           = errors.New("could not find media")
	ErrMediaNotInitialized     = errors.New("media not initialized")
	ErrMediaListCreate         = errors.New("could not create media list")
	ErrMediaListNotFound       = errors.New("could not find media list")
	ErrMediaListNotInitialized = errors.New("media list not initialized")
	ErrMissingMediaStats       = errors.New("could not get media statistics")
	ErrInvalidMediaStats       = errors.New("invalid media statistics")
	ErrMissingMediaLocation    = errors.New("could not get media location")
	ErrMediaMetaSave           = errors.New("could not save media metadata")
	ErrMediaParse              = errors.New("could not parse media")
)

// Media track errors.
var (
	ErrMediaTrackNotInitialized = errors.New("media track not initialized")
	ErrMediaTrackNotFound       = errors.New("could not find media track")
	ErrInvalidMediaTrack        = errors.New("invalid media track")
)

// Event manager errors.
var (
	ErrMissingEventManager  = errors.New("could not get event manager instance")
	ErrInvalidEventCallback = errors.New("invalid event callback")
)

// Audio/Video errors.
var (
	ErrAudioOutputListMissing = errors.New("could not get audio output list")
	ErrAudioOutputSet         = errors.New("could not set audio output")
	ErrFilterListMissing      = errors.New("could not get filter list")
	ErrStereoModeSet          = errors.New("could not set stereo mode")
)

// Renderer discoverer errors.
var (
	ErrRendererDiscovererParse          = errors.New("could not parse renderer discoverer")
	ErrRendererDiscovererCreate         = errors.New("could not create renderer discoverer")
	ErrRendererDiscovererNotInitialized = errors.New("renderer discoverer not initialized")
	ErrRendererDiscovererStart          = errors.New("could not start renderer discoverer")
	ErrRendererNotInitialized           = errors.New("renderer not initialized")
)

// Equalizer errors.
var (
	ErrEqualizerCreate         = errors.New("could not create equalizer")
	ErrEqualizerNotInitialized = errors.New("equalizer not initialized")
	ErrEqualizerAmpValueSet    = errors.New("could not set equalizer amplification value")
)

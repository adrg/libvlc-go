package vlc

import "errors"

// Module errors.
var (
	ErrModuleInitialize     = errors.New("could not initialize module")
	ErrModuleNotInitialized = errors.New("module not initialized")
)

// Player errors.
var (
	ErrPlayerCreate             = errors.New("could not create player")
	ErrPlayerNotInitialized     = errors.New("player not initialized")
	ErrListPlayerCreate         = errors.New("could not create list player")
	ErrListPlayerNotInitialized = errors.New("list player not initialized")
)

// Media errors.
var (
	ErrMediaCreate             = errors.New("could not create media")
	ErrMediaNotInitialized     = errors.New("media not initialized")
	ErrMediaListCreate         = errors.New("could not create media list")
	ErrMediaListNotInitialized = errors.New("media list not initialized")
	ErrMissingMediaStats       = errors.New("could not get media statistics")
	ErrInvalidMediaStats       = errors.New("invalid media statistics")
)

// Event manager errors.
var (
	ErrMissingEventManager  = errors.New("could not get event manager instance")
	ErrInvalidEventCallback = errors.New("invalid event callback")
)

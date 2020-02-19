package vlc

import "errors"

// Module errors.
var (
	ErrModuleNotInitialized = errors.New("module not initialized")
)

// Player errors.
var (
	ErrPlayerNotInitialized     = errors.New("player not initialized")
	ErrListPlayerNotInitialized = errors.New("list player not initialized")
)

// Media errors.
var (
	ErrMediaNotInitialized     = errors.New("media not initialized")
	ErrMediaListNotInitialized = errors.New("media list not initialized")
	ErrMissingMediaStats       = errors.New("could not get media statistics")
	ErrInvalidMediaStats       = errors.New("invalid media statistics")
)

// Event manager errors.
var (
	ErrMissingEventManager  = errors.New("could not get event manager instance")
	ErrInvalidEventCallback = errors.New("invalid event callback")
)

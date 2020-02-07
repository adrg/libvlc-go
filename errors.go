package vlc

import "errors"

// Common errors.
var (
	ErrModuleNotInitialized     = errors.New("module not initialized")
	ErrPlayerNotInitialized     = errors.New("player not initialized")
	ErrListPlayerNotInitialized = errors.New("list player not initialized")
	ErrMediaNotInitialized      = errors.New("media not initialized")
	ErrMediaListNotInitialized  = errors.New("media list not initialized")
	ErrMissingEventManager      = errors.New("could not get event manager instance")
	ErrInvalidEventCallback     = errors.New("invalid event callback")
)

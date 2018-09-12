package vlc

type Event uint

// Media events
const (
	// Metadata of a media item changed.
	MediaMetaChanged Event = iota

	// Subitem was added to a media item.
	MediaSubItemAdded

	// Duration of a media item changed.
	MediaDurationChanged

	// Parsing state of a media item changed.
	MediaParsedChanged

	// A media item was freed.
	MediaFreed

	// State of the media item changed.
	MediaStateChanged

	// Subitem tree was added to a media item.
	MediaSubItemTreeAdded
)

// Player events
const (
	MediaPlayerMediaChanged Event = 0x100
	MediaPlayerIdle
	MediaPlayerOpening
	MediaPlayerBuffering
	MediaPlayerPlaying
	MediaPlayerPaused
	MediaPlayerStopped
	MediaPlayerForward
	MediaPlayerBackward
	MediaPlayerEndReached
	MediaPlayerEncounteredError
	MediaPlayerTimeChanged
	MediaPlayerPositionChanged
	MediaPlayerSeekableChanged
	MediaPlayerPausableChanged
	MediaPlayerTitleChanged
	MediaPlayerSnapshotTaken
	MediaPlayerLengthChanged
	MediaPlayerVout
	MediaPlayerScrambledChanged
	MediaPlayerESAdded
	MediaPlayerESDeleted
	MediaPlayerESSelected
	MediaPlayerCorked
	MediaPlayerUncorked
	MediaPlayerMuted
	MediaPlayerUnmuted
	MediaPlayerAudioVolume
	MediaPlayerAudioDevice
	MediaPlayerChapterChanged
)

// Media list events
const (
	// A media item was added to a media list.
	MediaListItemAdded Event = 0x200

	// A media item is about to get added to a media list.
	MediaListWillAddItem

	// A media item was deleted from a media list.
	MediaListItemDeleted

	// A media item is about to get deleted from a media list.
	MediaListWillDeleteItem

	// A media list has reached the end.
	MediaListEndReached

	// Playback of a media list player has started.
	MediaListPlayerPlayed = 0x400

	// The current item of a media list player has changed to a different item.
	MediaListPlayerNextItemSet

	// Playback of a media list player has stopped.
	MediaListPlayerStopped
)

// Renderer events
const (
	// A new renderer item was found by a renderer discoverer.
	// The renderer item is valid until deleted.
	RendererDiscovererItemAdded Event = 0x502

	// A previously discovered renderer item was deleted by a renderer
	// discoverer. The renderer item is no longer valid.
	RendererDiscovererItemDeleted
)

// VideoLAN Manager events
const (
	VlmMediaAdded Event = 0x600
	VlmMediaRemoved
	VlmMediaChanged
	VlmMediaInstanceStarted
	VlmMediaInstanceStopped
	VlmMediaInstanceStatusInit
	VlmMediaInstanceStatusOpening
	VlmMediaInstanceStatusPlaying
	VlmMediaInstanceStatusPause
	VlmMediaInstanceStatusEnd
	VlmMediaInstanceStatusError
)

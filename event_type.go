package vlc

type EventType uint16

// Event types
const (
	MediaMetaChanged EventType = iota
	MediaSubItemAdded
	MediaDurationChanged
	MediaParsedChanged
	MediaFreed
	MediaStateChanged
	MediaSubItemTreeAdded
)

const (
	MediaPlayerMediaChanged = iota + 0x100
	MediaPlayerNothingSpecial
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

const (
	MediaListItemAdded = iota + 0x200
	MediaListWillAddItem
	MediaListItemDeleted
	MediaListWillDeleteItem
	MediaListEndReached
)

const (
	MediaListViewItemAdded = iota + 0x300
	MediaListViewWillAddItem
	MediaListViewItemDeleted
	MediaListViewWillDeleteItem
)

const (
	MediaListPlayerPlayed = iota + 0x400
	MediaListPlayerNextItemSet
	MediaListPlayerStopped
)

const (

	// Deprected
	MediaDiscovererStarted = iota + ox500
	MediaDiscovererEnded

	RendererDiscovererItemAdded
	RendererDiscovererItemDeleted
)

const (
	VlmMediaAdded = iota + ox600
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

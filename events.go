package vlc

// Event represents an event that can occur inside libvlc.
type Event int

// Media events.
const (
	// MediaMetaChanged is triggered when the metadata of a media item changes.
	MediaMetaChanged Event = iota

	// MediaSubItemAdded is triggered when a Subitem is added to a media item.
	MediaSubItemAdded

	// MediaDurationChanged is triggered when the duration
	// of a media item changes.
	MediaDurationChanged

	// MediaParsedChanged is triggered when the parsing state
	// of a media item changes.
	MediaParsedChanged

	// MediaFreed is triggered when a media item is freed.
	MediaFreed

	// MediaStateChanged is triggered when the state of the media item changes.
	MediaStateChanged

	// MediaSubItemTreeAdded is triggered when a Subitem tree is
	// added to a media item.
	MediaSubItemTreeAdded

	// MediaThumbnailGenerated is triggered when a thumbnail
	// generation is completed.
	MediaThumbnailGenerated
)

// Player events.
const (
	MediaPlayerMediaChanged Event = 0x100 + iota
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

// Media list events.
const (
	// MediaListItemAdded is triggered when a media item is added to a media list.
	MediaListItemAdded Event = 0x200 + iota

	// MediaListWillAddItem is triggered when a media item is about to get
	// added to a media list.
	MediaListWillAddItem

	// MediaListItemDeleted is triggered when a media item is deleted
	// from a media list.
	MediaListItemDeleted

	// MediaListWillDeleteItem is triggered when a media item is about to get
	// deleted from a media list.
	MediaListWillDeleteItem

	// MediaListEndReached is triggered when a media list has reached the end.
	MediaListEndReached
)

// Deprecated events.
const (
	MediaListViewItemAdded = 0x300 + iota
	MediaListViewWillAddItem
	MediaListViewItemDeleted
	MediaListViewWillDeleteItem
)

const (
	// MediaListPlayerPlayed is triggered when playback of the media list
	// of the list player has ended.
	MediaListPlayerPlayed = 0x400 + iota

	// MediaListPlayerNextItemSet is triggered when the current item
	// of a media list player has changed to a different item.
	MediaListPlayerNextItemSet

	// MediaListPlayerStopped is triggered when playback
	// of a media list player is stopped programmatically.
	MediaListPlayerStopped
)

// Deprecated events.
const (
	MediaDiscovererStarted Event = 0x500 + iota
	MediaDiscovererEnded
)

// Renderer events.
const (
	// RendererDiscovererItemAdded is triggered when a new renderer item is
	// found by a renderer discoverer. The renderer item is valid until deleted.
	RendererDiscovererItemAdded Event = 0x502 + iota

	// RendererDiscovererItemDeleted is triggered when a previously discovered
	// renderer item was deleted by a renderer discoverer. The renderer item
	// is no longer valid.
	RendererDiscovererItemDeleted
)

// VideoLAN Manager events.
const (
	VlmMediaAdded Event = 0x600 + iota
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

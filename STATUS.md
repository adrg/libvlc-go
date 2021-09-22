## Core

| ☐ | Binding                                | Implementation         | Versions   |
|---|:---------------------------------------|:-----------------------|:-----------|
| ☒ | libvlc_new                             | vlc.Init               | `v2`, `v3` |
| ☒ | libvlc_release                         | vlc.Release            | `v2`, `v3` |
| ☒ | libvlc_add_intf                        | vlc.StartUserInterface | `v2`, `v3` |
| ☐ | libvlc_set_exit_handler                |                        | `v2`, `v3` |
| ☒ | libvlc_set_user_agent                  | vlc.SetAppName         | `v2`, `v3` |
| ☒ | libvlc_set_app_id                      | vlc.SetAppID           | `v2`, `v3` |
| ☒ | libvlc_get_version                     | vlc.Version.Runtime    | `v2`, `v3` |
| ☒ | libvlc_get_compiler                    | vlc.Version.Compiler   | `v2`, `v3` |
| ☒ | libvlc_get_changeset                   | vlc.Version.Changeset  | `v2`, `v3` |

Reference: [libVLC core](https://www.videolan.org/developers/vlc/doc/doxygen/html/group__libvlc__core.html).

## Player

| ☐ | Binding                                           | Implementation             | Versions   |
|---|:--------------------------------------------------|:---------------------------|:-----------|
| ☒ | libvlc_media_player_new                           | vlc.NewPlayer              | `v2`, `v3` |
| ☐ | libvlc_media_player_new_from_media                |                            | `v2`, `v3` |
| ☒ | libvlc_media_player_release                       | Player.Release             | `v2`, `v3` |
| ☒ | libvlc_media_player_set_media                     | Player.SetMedia            | `v2`, `v3` |
| ☒ | libvlc_media_player_get_media                     | Player.Media               | `v2`, `v3` |
| ☒ | libvlc_media_player_event_manager                 | Player.EventManager        | `v2`, `v3` |
| ☒ | libvlc_media_player_is_playing                    | Player.IsPlaying           | `v2`, `v3` |
| ☒ | libvlc_media_player_play                          | Player.Play                | `v2`, `v3` |
| ☒ | libvlc_media_player_set_pause                     | Player.SetPause            | `v2`, `v3` |
| ☒ | libvlc_media_player_pause                         | Player.TogglePause         | `v2`, `v3` |
| ☒ | libvlc_media_player_stop                          | Player.Stop                | `v2`, `v3` |
| ☒ | libvlc_media_player_set_renderer                  | Player.SetRenderer         | `v3`       |
| ☒ | libvlc_media_player_set_nsobject                  | Player.SetNSObject         | `v2`, `v3` |
| ☒ | libvlc_media_player_get_nsobject                  | Player.NSObject            | `v2`, `v3` |
| ☒ | libvlc_media_player_set_xwindow                   | Player.SetXWindow          | `v2`, `v3` |
| ☒ | libvlc_media_player_get_xwindow                   | Player.XWindow             | `v2`, `v3` |
| ☒ | libvlc_media_player_set_hwnd                      | Player.SetHWND             | `v2`, `v3` |
| ☒ | libvlc_media_player_get_hwnd                      | Player.HWND                | `v2`, `v3` |
| ☐ | libvlc_media_player_set_android_context           |                            | `v2`, `v3` |
| ☒ | libvlc_media_player_get_length                    | Player.MediaLength         | `v2`, `v3` |
| ☒ | libvlc_media_player_get_time                      | Player.MediaTime           | `v2`, `v3` |
| ☒ | libvlc_media_player_set_time                      | Player.SetMediaTime        | `v2`, `v3` |
| ☒ | libvlc_media_player_get_position                  | Player.MediaPosition       | `v2`, `v3` |
| ☒ | libvlc_media_player_set_position                  | Player.SetMediaPosition    | `v2`, `v3` |
| ☒ | libvlc_media_player_set_chapter                   | Player.SetChapter          | `v2`, `v3` |
| ☒ | libvlc_media_player_get_chapter                   | Player.ChapterIndex        | `v2`, `v3` |
| ☒ | libvlc_media_player_get_chapter_count             | Player.ChapterCount        | `v2`, `v3` |
| ☒ | libvlc_media_player_will_play                     | Player.WillPlay            | `v2`, `v3` |
| ☒ | libvlc_media_player_get_chapter_count_for_title   | Player.TitleChapterCount   | `v2`, `v3` |
| ☒ | libvlc_media_player_set_title                     | Player.SetTitle            | `v2`, `v3` |
| ☒ | libvlc_media_player_get_title                     | Player.TitleIndex          | `v2`, `v3` |
| ☒ | libvlc_media_player_get_title_count               | Player.TitleCount          | `v2`, `v3` |
| ☒ | libvlc_media_player_previous_chapter              | Player.PreviousChapter     | `v2`, `v3` |
| ☒ | libvlc_media_player_next_chapter                  | Player.NextChapter         | `v2`, `v3` |
| ☒ | libvlc_media_player_get_full_title_descriptions   | Player.Titles              | `v3`       |
| ☒ | libvlc_media_player_get_full_chapter_descriptions | Player.TitleChapters       | `v3`       |
| ☒ | libvlc_media_player_get_rate                      | Player.SetPlaybackRate     | `v2`, `v3` |
| ☒ | libvlc_media_player_set_rate                      | Player.PlaybackRate        | `v2`, `v3` |
| ☒ | libvlc_media_player_get_state                     | Player.MediaState          | `v2`, `v3` |
| ☒ | libvlc_media_player_has_vout                      | Player.VideoOutputCount    | `v2`, `v3` |
| ☒ | libvlc_media_player_is_seekable                   | Player.IsSeekable          | `v2`, `v3` |
| ☒ | libvlc_media_player_can_pause                     | Player.CanPause            | `v2`, `v3` |
| ☒ | libvlc_media_player_program_scrambled             | Player.IsScrambled         | `v2`, `v3` |
| ☒ | libvlc_media_player_next_frame                    | Player.NextFrame           | `v2`, `v3` |
| ☒ | libvlc_media_player_set_equalizer                 | Player.SetEqualizer        | `v2`, `v3` |
| ☒ | libvlc_media_player_get_role                      | Player.Role                | `v3`       |
| ☒ | libvlc_media_player_set_role                      | Player.SetRole             | `v3`       |
| ☐ | libvlc_media_player_navigate                      |                            | `v2`, `v3` |
| ☐ | libvlc_media_player_set_video_title_display       |                            | `v2`, `v3` |
| ☐ | libvlc_media_player_add_slave                     |                            | `v2`, `v3` |

Reference: [libVLC media player](https://www.videolan.org/developers/vlc/doc/doxygen/html/group__libvlc__media__player.html).

## Audio controls

| ☐ | Binding                                         | Implementation               | Versions   |
|---|:------------------------------------------------|:-----------------------------|:-----------|
| ☒ | libvlc_audio_output_list_get                    | vlc.AudioOutputList          | `v2`, `v3` |
| ☒ | libvlc_audio_filter_list_get                    | vlc.ListAudioFilters         | `v2`, `v3` |
| ☒ | libvlc_audio_output_set                         | Player.SetAudioOutput        | `v2`, `v3` |
| ☒ | libvlc_audio_output_device_enum                 | Player.AudioOutputDevices    | `v2`, `v3` |
| ☒ | libvlc_audio_output_device_list_get             | vlc.ListAudioOutputDevices   | `v2`, `v3` |
| ☒ | libvlc_audio_output_device_set                  | vlc.SetAudioOutputDevice     | `v2`, `v3` |
| ☒ | libvlc_audio_output_device_get                  | vlc.AudioOutputDevice        | `v3`       |
| ☒ | libvlc_audio_toggle_mute                        | Player.ToggleMute            | `v2`, `v3` |
| ☒ | libvlc_audio_get_mute                           | Player.IsMuted               | `v2`, `v3` |
| ☒ | libvlc_audio_set_mute                           | Player.SetMute               | `v2`, `v3` |
| ☒ | libvlc_audio_get_volume                         | Player.Volume                | `v2`, `v3` |
| ☒ | libvlc_audio_set_volume                         | Player.SetVolume             | `v2`, `v3` |
| ☒ | libvlc_audio_get_track_count                    | Player.AudioTrackCount       | `v2`, `v3` |
| ☒ | libvlc_audio_get_track_description              | Player.AudioTrackDescriptors | `v2`, `v3` |
| ☒ | libvlc_audio_get_track                          | Player.AudioTrackID          | `v2`, `v3` |
| ☒ | libvlc_audio_set_track                          | Player.SetAudioTrack         | `v2`, `v3` |
| ☒ | libvlc_audio_get_channel                        | Player.StereoMode            | `v2`, `v3` |
| ☒ | libvlc_audio_set_channel                        | Player.SetStereoMode         | `v2`, `v3` |
| ☒ | libvlc_audio_get_delay                          | Player.AudioDelay            | `v2`, `v3` |
| ☒ | libvlc_audio_set_delay                          | Player.SetAudioDelay         | `v2`, `v3` |
| ☒ | libvlc_audio_equalizer_get_preset_count         | vlc.EqualizerPresetCount     | `v2`, `v3` |
| ☒ | libvlc_audio_equalizer_get_preset_name          | vlc.EqualizerPresetName      | `v2`, `v3` |
| ☒ | libvlc_audio_equalizer_get_band_count           | vlc.EqualizerBandCount       | `v2`, `v3` |
| ☒ | libvlc_audio_equalizer_get_band_frequency       | vlc.EqualizerBandFrequency   | `v2`, `v3` |
| ☒ | libvlc_audio_equalizer_new                      | vlc.NewEqualizer             | `v2`, `v3` |
| ☒ | libvlc_audio_equalizer_new_from_preset          | vlc.NewEqualizerFromPreset   | `v2`, `v3` |
| ☒ | libvlc_audio_equalizer_release                  | Equalizer.Release            | `v2`, `v3` |
| ☒ | libvlc_audio_equalizer_set_preamp               | Equalizer.SetPreampValue     | `v2`, `v3` |
| ☒ | libvlc_audio_equalizer_get_preamp               | Equalizer.PreampValue        | `v2`, `v3` |
| ☒ | libvlc_audio_equalizer_set_amp_at_index         | Equalizer.SetAmpValueAtIndex | `v2`, `v3` |
| ☒ | libvlc_audio_equalizer_get_amp_at_index         | Equalizer.AmpValueAtIndex    | `v2`, `v3` |
| ☐ | libvlc_audio_set_callbacks                      |                              | `v2`, `v3` |
| ☐ | libvlc_audio_set_volume_callback                |                              | `v2`, `v3` |
| ☐ | libvlc_audio_set_format_callbacks               |                              | `v2`, `v3` |
| ☐ | libvlc_audio_set_format                         |                              | `v2`, `v3` |

Reference: [libVLC audio controls](https://www.videolan.org/developers/vlc/doc/doxygen/html/group__libvlc__audio.html).

## Video controls

| ☐ | Binding                                           | Implementation                  | Versions   |
|---|:--------------------------------------------------|:--------------------------------|:-----------|
| ☒ | libvlc_video_filter_list_get                      | vlc.ListVideoFilters            | `v2`, `v3` |
| ☒ | libvlc_toggle_fullscreen                          | Player.ToggleFullScreen         | `v2`, `v3` |
| ☒ | libvlc_set_fullscreen                             | Player.SetFullScreen            | `v2`, `v3` |
| ☒ | libvlc_get_fullscreen                             | Player.IsFullScreen             | `v2`, `v3` |
| ☒ | libvlc_video_set_key_input                        | Player.SetKeyInput              | `v2`, `v3` |
| ☒ | libvlc_video_set_mouse_input                      | Player.SetMouseInput            | `v2`, `v3` |
| ☒ | libvlc_video_get_size                             | Player.VideoDimensions          | `v2`, `v3` |
| ☒ | libvlc_video_get_cursor                           | Player.CursorPosition           | `v2`, `v3` |
| ☒ | libvlc_video_get_scale                            | Player.Scale                    | `v2`, `v3` |
| ☒ | libvlc_video_set_scale                            | Player.SetScale                 | `v2`, `v3` |
| ☒ | libvlc_video_get_aspect_ratio                     | Player.AspectRatio              | `v2`, `v3` |
| ☒ | libvlc_video_set_aspect_ratio                     | Player.SetAspectRatio           | `v2`, `v3` |
| ☒ | libvlc_video_update_viewpoint                     | Player.UpdateVideoViewpoint     | `v3`       |
| ☒ | libvlc_video_get_spu                              | Player.SubtitleTrackID          | `v2`, `v3` |
| ☒ | libvlc_video_get_spu_count                        | Player.SubtitleTrackCount       | `v2`, `v3` |
| ☒ | libvlc_video_get_spu_description                  | Player.SubtitleTrackDescriptors | `v2`, `v3` |
| ☒ | libvlc_video_set_spu                              | Player.SetSubtitleTrack         | `v2`, `v3` |
| ☒ | libvlc_video_get_spu_delay                        | Player.SubtitleDelay            | `v2`, `v3` |
| ☒ | libvlc_video_set_spu_delay                        | Player.SetSubtitleDelay         | `v2`, `v3` |
| ☐ | libvlc_video_get_teletext                         |                                 | `v2`, `v3` |
| ☐ | libvlc_video_set_teletext                         |                                 | `v2`, `v3` |
| ☒ | libvlc_video_get_track_count                      | Player.VideoTrackCount          | `v2`, `v3` |
| ☒ | libvlc_video_get_track_description                | Player.VideoTrackDescriptors    | `v2`, `v3` |
| ☒ | libvlc_video_get_track                            | Player.VideoTrackID             | `v2`, `v3` |
| ☒ | libvlc_video_set_track                            | Player.SetVideoTrack            | `v2`, `v3` |
| ☐ | libvlc_video_take_snapshot                        |                                 | `v2`, `v3` |
| ☐ | libvlc_video_get_marquee_int                      |                                 | `v2`, `v3` |
| ☐ | libvlc_video_set_marquee_int                      |                                 | `v2`, `v3` |
| ☐ | libvlc_video_set_marquee_string                   |                                 | `v2`, `v3` |
| ☐ | libvlc_video_get_logo_int                         |                                 | `v2`, `v3` |
| ☐ | libvlc_video_set_logo_int                         |                                 | `v2`, `v3` |
| ☐ | libvlc_video_set_logo_string                      |                                 | `v2`, `v3` |
| ☐ | libvlc_video_get_adjust_int                       |                                 | `v2`, `v3` |
| ☐ | libvlc_video_set_adjust_int                       |                                 | `v2`, `v3` |
| ☐ | libvlc_video_get_adjust_float                     |                                 | `v2`, `v3` |
| ☐ | libvlc_video_set_adjust_float                     |                                 | `v2`, `v3` |
| ☐ | libvlc_video_set_callbacks                        |                                 | `v2`, `v3` |
| ☐ | libvlc_video_set_format                           |                                 | `v2`, `v3` |
| ☐ | libvlc_video_set_format_callbacks                 |                                 | `v2`, `v3` |

Reference: [libVLC video controls](https://www.videolan.org/developers/vlc/doc/doxygen/html/group__libvlc__video.html).

## List player

| ☐ | Binding                                      | Implementation             | Versions   |
|---|:---------------------------------------------|:---------------------------|:-----------|
| ☒ | libvlc_media_list_player_new                 | vlc.NewListPlayer          | `v2`, `v3` |
| ☒ | libvlc_media_list_player_release             | ListPlayer.Release         | `v2`, `v3` |
| ☒ | libvlc_media_list_player_event_manager       | ListPlayer.EventManager    | `v2`, `v3` |
| ☒ | libvlc_media_list_player_set_media_player    | ListPlayer.SetPlayer       | `v2`, `v3` |
| ☒ | libvlc_media_list_player_get_media_player    | ListPlayer.Player          | `v2`, `v3` |
| ☒ | libvlc_media_list_player_set_media_list      | ListPlayer.SetMediaList    | `v2`, `v3` |
| ☒ | libvlc_media_list_player_play                | ListPlayer.Play            | `v2`, `v3` |
| ☒ | libvlc_media_list_player_pause               | ListPlayer.TogglePause     | `v2`, `v3` |
| ☒ | libvlc_media_list_player_set_pause           | ListPlayer.SetPause        | `v3`       |
| ☒ | libvlc_media_list_player_is_playing          | ListPlayer.IsPlaying       | `v2`, `v3` |
| ☒ | libvlc_media_list_player_get_state           | ListPlayer.MediaState      | `v2`, `v3` |
| ☒ | libvlc_media_list_player_play_item_at_index  | ListPlayer.PlayAtIndex     | `v2`, `v3` |
| ☒ | libvlc_media_list_player_play_item           | ListPlayer.PlayItem        | `v2`, `v3` |
| ☒ | libvlc_media_list_player_stop                | ListPlayer.Stop            | `v2`, `v3` |
| ☒ | libvlc_media_list_player_next                | ListPlayer.PlayNext        | `v2`, `v3` |
| ☒ | libvlc_media_list_player_previous            | ListPlayer.PlayPrevious    | `v2`, `v3` |
| ☒ | libvlc_media_list_player_set_playback_mode   | ListPlayer.SetPlaybackMode | `v2`, `v3` |

Reference: [libVLC media list player](https://www.videolan.org/developers/vlc/doc/doxygen/html/group__libvlc__media__list__player.html).

## Media

| ☐ | Binding                                | Implementation              | Versions   |
|---|:---------------------------------------|:----------------------------|:-----------|
| ☒ | libvlc_media_new_location              | vlc.NewMediaFromPath        | `v2`, `v3` |
| ☒ | libvlc_media_new_path                  | vlc.NewMediaFromURL         | `v2`, `v3` |
| ☐ | libvlc_media_new_fd                    |                             | `v2`, `v3` |
| ☒ | libvlc_media_new_callbacks             | vlc.NewMediaFromReadSeeker  | `v3`       |
| ☐ | libvlc_media_new_as_node               |                             | `v2`, `v3` |
| ☒ | libvlc_media_add_option                | Media.AddOptions            | `v2`, `v3` |
| ☐ | libvlc_media_add_option_flag           |                             | `v2`, `v3` |
| ☒ | libvlc_media_release                   | Media.Release               | `v2`, `v3` |
| ☒ | libvlc_media_get_mrl                   | Media.Location              | `v2`, `v3` |
| ☒ | libvlc_media_duplicate                 | Media.Duplicate             | `v2`, `v3` |
| ☒ | libvlc_media_get_meta                  | Media.Meta                  | `v2`, `v3` |
| ☒ | libvlc_media_set_meta                  | Media.SetMeta               | `v2`, `v3` |
| ☒ | libvlc_media_save_meta                 | Media.SaveMeta              | `v2`, `v3` |
| ☒ | libvlc_media_get_state                 | Media.State                 | `v2`, `v3` |
| ☒ | libvlc_media_get_stats                 | Media.Stats                 | `v2`, `v3` |
| ☒ | libvlc_media_subitems                  | Media.SubItems              | `v2`, `v3` |
| ☒ | libvlc_media_event_manager             | Media.EventManager          | `v2`, `v3` |
| ☒ | libvlc_media_get_duration              | Media.Duration              | `v2`, `v3` |
| ☒ | libvlc_media_parse_with_options        | Media.ParseWithOptions      | `v3`       |
| ☒ | libvlc_media_parse_stop                | Media.StopParse             | `v3`       |
| ☒ | libvlc_media_get_parsed_status         | Media.ParseStatus           | `v3`       |
| ☒ | libvlc_media_set_user_data             | Media.UserData              | `v2`, `v3` |
| ☒ | libvlc_media_get_user_data             | Media.SetUserData           | `v2`, `v3` |
| ☒ | libvlc_media_tracks_get                | Media.Tracks                | `v2`, `v3` |
| ☒ | libvlc_media_get_codec_description     | MediaTrack.CodecDescription | `v3`       |
| ☒ | libvlc_media_get_type                  | Media.Type                  | `v3`       |
| ☐ | libvlc_media_slaves_add                |                             | `v3`       |
| ☐ | libvlc_media_slaves_clear              |                             | `v3`       |
| ☐ | libvlc_media_slaves_get                |                             | `v3`       |
| ☐ | libvlc_media_slaves_release            |                             | `v3`       |
| ☒ | libvlc_media_parse                     | Media.Parse                 | `v3`, `v3` |
| ☒ | libvlc_media_parse_async               | Media.ParseAsync            | `v2`, `v3` |
| ☒ | libvlc_media_is_parsed                 | Media.IsParsed              | `v2`, `v3` |

Reference: [libVLC media](https://www.videolan.org/developers/vlc/doc/doxygen/html/group__libvlc__media.html).

## Media list

| ☐ | Binding                         | Implementation               | Versions   |
|---|:--------------------------------|:---------------------------  |:-----------|
| ☒ | libvlc_media_list_new           | vlc.NewMediaList             | `v2`, `v3` |
| ☒ | libvlc_media_list_release       | MediaList.Release            | `v2`, `v3` |
| ☒ | libvlc_media_list_set_media     | MediaList.AssociateMedia     | `v2`, `v3` |
| ☒ | libvlc_media_list_media         | MediaList.AssociatedMedia    | `v2`, `v3` |
| ☒ | libvlc_media_list_add_media     | MediaList.AddMedia           | `v2`, `v3` |
| ☒ | libvlc_media_list_insert_media  | MediaList.InsertMedia        | `v2`, `v3` |
| ☒ | libvlc_media_list_remove_index  | MediaList.RemoveMediaAtIndex | `v2`, `v3` |
| ☒ | libvlc_media_list_count         | MediaList.Count              | `v2`, `v3` |
| ☒ | libvlc_media_list_item_at_index | MediaList.MediaAtIndex       | `v2`, `v3` |
| ☒ | libvlc_media_list_index_of_item | MediaList.IndexOfMedia       | `v2`, `v3` |
| ☒ | libvlc_media_list_is_readonly   | MediaList.IsReadOnly         | `v2`, `v3` |
| ☒ | libvlc_media_list_lock          | MediaList.Lock               | `v2`, `v3` |
| ☒ | libvlc_media_list_unlock        | MediaList.Unlock             | `v2`, `v3` |
| ☒ | libvlc_media_list_event_manager | MediaList.EventManager       | `v2`, `v3` |

Reference: [libVLC media list](https://www.videolan.org/developers/vlc/doc/doxygen/html/group__libvlc__media__list.html).

## Event manager

| ☐ | Binding             | Implementation      | Versions   |
|---|:--------------------|:--------------------|:-----------|
| ☒ | libvlc_event_attach | EventManager.Attach | `v2`, `v3` |
| ☒ | libvlc_event_detach | EventManager.Detach | `v2`, `v3` |

Reference: [libVLC events](https://www.videolan.org/developers/vlc/doc/doxygen/html/group__libvlc__event.html).

## Renderer discoverer

| ☐ | Binding                                  | Implementation             | Versions   |
|---|:-----------------------------------------|:---------------------------|:-----------|
| ☒ | libvlc_renderer_item_name                | Renderer.Name              | `v3`       |
| ☒ | libvlc_renderer_item_type                | Renderer.Type              | `v3`       |
| ☒ | libvlc_renderer_item_icon_uri            | Renderer.IconURI           | `v3`       |
| ☒ | libvlc_renderer_item_flags               | Renderer.Flags             | `v3`       |
| ☒ | libvlc_renderer_discoverer_new           | vlc.NewRendererDiscoverer  | `v3`       |
| ☒ | libvlc_renderer_discoverer_release       | RendererDiscoverer.Release | `v3`       |
| ☒ | libvlc_renderer_discoverer_start         | RendererDiscoverer.Start   | `v3`       |
| ☒ | libvlc_renderer_discoverer_stop          | RendererDiscoverer.Stop    | `v3`       |
| ☒ | libvlc_renderer_discoverer_list_get      | ListRendererDiscoverers    | `v3`       |

Reference: [libVLC renderer discoverer](https://www.videolan.org/developers/vlc/doc/doxygen/html/group__libvlc__renderer__discoverer.html).

| ☐ | Binding                                  | Implementation             | Versions   |
|---|:-----------------------------------------|:---------------------------|:-----------|
| ☒ | libvlc_media_discoverer_new              | vlc.NewMediaDiscoverer     | `v3`       |
| ☒ | libvlc_media_discoverer_start            | MediaDiscoverer.Start      | `v3`       |
| ☒ | libvlc_media_discoverer_stop             | MediaDiscoverer.Stop       | `v3`       |
| ☒ | libvlc_media_discoverer_release          | MediaDiscoverer.Release    | `v3`       |
| ☒ | libvlc_media_discoverer_media_list       | MediaDiscoverer.MediaList  | `v3`       |
| ☒ | libvlc_media_discoverer_is_running       | MediaDiscoverer.IsRunning  | `v3`       |
| ☒ | libvlc_media_discoverer_list_get         | vlc.ListMediaDiscoverers   | `v3`       |

Reference: [libVLC media discovery](https://www.videolan.org/developers/vlc/doc/doxygen/html/group__libvlc__media__discoverer.html).

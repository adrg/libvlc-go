<h1 align="center">
  <div>
    <img src="https://raw.githubusercontent.com/adrg/adrg.github.io/master/assets/projects/libvlc-go/logo.svg" alt="libvlc-go logo"/>
  </div>
</h1>

<h3 align="center">Go bindings for libVLC and high-level media player interface.</h3>

<p align="center">
    <a href="https://github.com/adrg/libvlc-go/actions/workflows/ci.yml">
        <img alt="Build status" src="https://github.com/adrg/libvlc-go/actions/workflows/ci.yml/badge.svg">
    </a>
    <a href="https://pkg.go.dev/github.com/adrg/libvlc-go/v3">
        <img alt="pkg.go.dev documentation" src="https://img.shields.io/badge/go.dev-reference-007d9c?logo=go&logoColor=white">
    </a>
    <a href="https://opensource.org/licenses/MIT" rel="nofollow">
        <img alt="MIT license" src="https://img.shields.io/github/license/adrg/libvlc-go"/>
    </a>
    <a href="https://github.com/avelino/awesome-go#video">
        <img alt="Awesome Go" src="https://awesome.re/mentioned-badge.svg"/>
    </a>
    <a href="https://ko-fi.com/T6T72WATK">
        <img alt="Buy me a coffee" src="https://img.shields.io/static/v1.svg?label=%20&message=Buy%20me%20a%20coffee&color=579fbf&logo=buy%20me%20a%20coffee&logoColor=white"/>
    </a>
    <br />
    <a href="https://goreportcard.com/report/github.com/adrg/libvlc-go">
        <img alt="Go report card" src="https://goreportcard.com/badge/github.com/adrg/libvlc-go" />
    </a>
    <a href="https://github.com/adrg/libvlc-go/graphs/contributors">
        <img alt="GitHub contributors" src="https://img.shields.io/github/contributors/adrg/libvlc-go" />
    </a>
    <a href="https://discord.gg/3h3K3JF">
        <img alt="Discord channel" src="https://img.shields.io/discord/716939396464508958?label=discord" />
    </a>
    <a href="https://github.com/adrg/libvlc-go/issues?q=is%3Aopen+is%3Aissue">
        <img alt="GitHub open issues" src="https://img.shields.io/github/issues-raw/adrg/libvlc-go">
    </a>
    <a href="https://github.com/adrg/libvlc-go/issues?q=is%3Aissue+is%3Aclosed">
        <img alt="GitHub closed issues" src="https://img.shields.io/github/issues-closed-raw/adrg/libvlc-go" />
    </a>
</p>

The package can be useful for adding multimedia capabilities to applications
through the provided player interfaces. It relies on Go modules in order to
mirror each supported major version of [libVLC](https://www.videolan.org/vlc/libvlc.html).

Documentation for v3, which implements bindings for libVLC 3.X, can be found on [pkg.go.dev](https://pkg.go.dev/github.com/adrg/libvlc-go/v3).  
Documentation for v2, which implements bindings for libVLC 2.X, can be found on [pkg.go.dev](https://pkg.go.dev/github.com/adrg/libvlc-go/v2).

<p align="center">
    <a href="https://github.com/adrg/libvlc-go-examples/tree/master/v3/gtk3_player">
      <img align="center" width="49%" alt="libvlc-go media player" src="https://raw.githubusercontent.com/adrg/adrg.github.io/master/assets/projects/libvlc-go/gtk3-media-player-example/libvlc-gtk3-media-player.jpg">
    </a>
    <a href="https://github.com/adrg/libvlc-go-examples/tree/master/v3/gtk3_screen_recorder">
      <img align="center" width="49%" alt="libvlc-go screen recorder" src="https://raw.githubusercontent.com/adrg/adrg.github.io/master/assets/projects/libvlc-go/gtk3-screen-recorder-example/libvlc-gtk3-screen-recorder.jpg">
    </a>
    <a href="https://github.com/adrg/libvlc-go-examples/tree/master/v3/gtk3_equalizer">
      <img align="center" width="49%" alt="libvlc-go equalizer" src="https://raw.githubusercontent.com/adrg/adrg.github.io/master/assets/projects/libvlc-go/gtk3-equalizer-example/libvlc-gtk3-equalizer.jpg">
    </a>
    <a href="https://github.com/adrg/libvlc-go-examples/tree/master/v3/gtk3_media_discovery">
      <img align="center" width="49%" alt="libvlc-go media discovery" src="https://raw.githubusercontent.com/adrg/adrg.github.io/master/assets/projects/libvlc-go/gtk3-media-discovery-example/libvlc-gtk3-media-discovery.jpg">
    </a>
</p>

Example applications:

* [GUI media player](https://github.com/adrg/libvlc-go-examples/tree/master/v3/gtk3_player)
* [GUI screen recorder](https://github.com/adrg/libvlc-go-examples/tree/master/v3/gtk3_screen_recorder)
* [GUI equalizer](https://github.com/adrg/libvlc-go-examples/tree/master/v3/gtk3_equalizer)
* [GUI media discovery](https://github.com/adrg/libvlc-go-examples/tree/master/v3/gtk3_media_discovery)

## Prerequisites

The libVLC development files are required. Instructions for installing the
VLC SDK on multiple operating systems can be found on the wiki pages of this project.

- [Install on Linux](https://github.com/adrg/libvlc-go/wiki/Install-on-Linux)
- [Install on Windows](https://github.com/adrg/libvlc-go/wiki/Install-on-Windows)
- [Install on macOS](https://github.com/adrg/libvlc-go/wiki/Install-on-macOS)

## Installation

In order to support multiple versions of libVLC, the package contains a Go
module for each major version of the API. Choose an installation option
depending on the version of libVLC you want to use.

**libVLC v3.X**

```bash
go get github.com/adrg/libvlc-go/v3
```

**libVLC v2.X**

```bash
go get github.com/adrg/libvlc-go/v2

# Build for libVLC < v2.2.0
go build -tags legacy
```

All versions above also work for projects which are not using Go modules.
However, consider switching to modules.

## Examples

* [GTK 3 media player](https://github.com/adrg/libvlc-go-examples/tree/master/v3/gtk3_player) (using [gotk3](https://github.com/gotk3/gotk3))
* [GTK 3 screen recorder](https://github.com/adrg/libvlc-go-examples/tree/master/v3/gtk3_screen_recorder) (using [gotk3](https://github.com/gotk3/gotk3))
* [GTK 3 media discovery](https://github.com/adrg/libvlc-go-examples/tree/master/v3/gtk3_media_discovery) (using [gotk3](https://github.com/gotk3/gotk3))
* [GTK 3 equalizer](https://github.com/adrg/libvlc-go-examples/tree/master/v3/gtk3_equalizer) (using [gotk3](https://github.com/gotk3/gotk3))
* [GTK 2 media player](https://github.com/adrg/libvlc-go-examples/tree/master/v3/gtk2_player) (using [go-gtk](https://github.com/mattn/go-gtk))
* [GTK 2 screen recorder](https://github.com/adrg/libvlc-go-examples/tree/master/v3/gtk2_screen_recorder) (using [go-gtk](https://github.com/mattn/go-gtk))
* [Basic player usage](https://github.com/adrg/libvlc-go-examples/blob/master/v3/player/player.go)
* [Basic list player usage](https://github.com/adrg/libvlc-go-examples/tree/master/v3/list_player/list_player.go)
* [Handling events](https://github.com/adrg/libvlc-go-examples/tree/master/v3/event_handling/event_handling.go)
* [Retrieve media tracks](https://github.com/adrg/libvlc-go-examples/blob/master/v3/media_tracks/media_tracks.go)
* [Retrieve media information](https://github.com/adrg/libvlc-go-examples/blob/master/v3/media_information/media_information.go)
* [Display screen as player media](https://github.com/adrg/libvlc-go-examples/blob/master/v3/display_screen_media/display_screen_media.go)
* [Stream media to Chromecast](https://github.com/adrg/libvlc-go-examples/blob/master/v3/chromecast_streaming/chromecast_streaming.go)
* [Player equalizer usage](https://github.com/adrg/libvlc-go-examples/blob/master/v3/equalizer/equalizer.go)

Examples for all supported API versions can be found at https://github.com/adrg/libvlc-go-examples.

## Usage

```go
package main

import (
    "log"

    vlc "github.com/adrg/libvlc-go/v3"
)

func main() {
    // Initialize libVLC. Additional command line arguments can be passed in
    // to libVLC by specifying them in the Init function.
    if err := vlc.Init("--no-video", "--quiet"); err != nil {
        log.Fatal(err)
    }
    defer vlc.Release()

    // Create a new player.
    player, err := vlc.NewPlayer()
    if err != nil {
        log.Fatal(err)
    }
    defer func() {
        player.Stop()
        player.Release()
    }()

    // Add a media file from path or from URL.
    // Set player media from path:
    // media, err := player.LoadMediaFromPath("localpath/test.mp4")
    // Set player media from URL:
    media, err := player.LoadMediaFromURL("http://stream-uk1.radioparadise.com/mp3-32")
    if err != nil {
        log.Fatal(err)
    }
    defer media.Release()

    // Retrieve player event manager.
    manager, err := player.EventManager()
    if err != nil {
        log.Fatal(err)
    }

    // Register the media end reached event with the event manager.
    quit := make(chan struct{})
    eventCallback := func(event vlc.Event, userData interface{}) {
        close(quit)
    }

    eventID, err := manager.Attach(vlc.MediaPlayerEndReached, eventCallback, nil)
    if err != nil {
        log.Fatal(err)
    }
    defer manager.Detach(eventID)

    // Start playing the media.
    err = player.Play()
    if err != nil {
        log.Fatal(err)
    }

    <-quit
}
```

## In action

A list of projects using libvlc-go, in alphabetical order. If you want to
showcase your project in this section, please create a pull request with it.

- [Alio](https://github.com/fenimore/alio) - Command-line music player with Emacs style key bindings.
- [Tripbot](https://github.com/adanalife/tripbot) - An ongoing 24/7 slow-TV art project.

## Stargazers over time

[![Stargazers over time](https://starchart.cc/adrg/libvlc-go.svg)](https://starchart.cc/adrg/libvlc-go)

## Contributing

Contributions in the form of pull requests, issues or just general feedback,
are always welcome.  
See [CONTRIBUTING.MD](CONTRIBUTING.md).

**Contributors**:
[adrg](https://github.com/adrg),
[fenimore](https://github.com/fenimore),
[tarrsalah](https://github.com/tarrsalah),
[danielpellon](https://github.com/danielpellon),
[patknight](https://github.com/patknight),
[sndnvaps](https://github.com/sndnvaps).

## Discord server

libvlc-go is part of the libVLC Discord Community [server](https://discord.gg/3h3K3JF). Feel free to come say hello!

## References

For more information see the
[libVLC](https://www.videolan.org/developers/vlc/doc/doxygen/html/group__libvlc.html) documentation.

## License

Copyright (c) 2018 Adrian-George Bostan.

This project is licensed under the [MIT license](https://opensource.org/licenses/MIT).
See [LICENSE](LICENSE) for more details.

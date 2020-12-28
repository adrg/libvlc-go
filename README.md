<h1 align="center">
  <div>
    <img src="https://raw.githubusercontent.com/adrg/adrg.github.io/master/assets/projects/libvlc-go/logo.png" height="45px" alt="libvlc-go logo"/>
  </div>
</h1>

<h3 align="center">Go bindings for libVLC 2.X/3.X/4.X and high-level media player interface.</h3>

<p align="center">
    <a href="https://pkg.go.dev/github.com/adrg/libvlc-go/v3">
        <img alt="pkg.go.dev documentation" src="https://img.shields.io/badge/go.dev-reference-007d9c?logo=go&logoColor=white">
    </a>
    <a href="https://opensource.org/licenses/MIT" rel="nofollow">
        <img alt="MIT license" src="https://img.shields.io/github/license/adrg/libvlc-go"/>
    </a>
    <a href="https://github.com/avelino/awesome-go#video">
        <img alt="Awesome Go" src="https://cdn.rawgit.com/sindresorhus/awesome/d7305f38d29fed78fa85652e3a63e154dd8e8829/media/badge.svg"/>
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

Documentation for v3, which implements bindings for libVLC 3.X, can be found on [pkg.go.dev](https://pkg.go.dev/github.com/adrg/libvlc-go/v3) and on [GoDoc](https://godoc.org/github.com/adrg/libvlc-go/v3).  
Documentation for v2, which implements bindings for libVLC 2.X, can be found on [pkg.go.dev](https://pkg.go.dev/github.com/adrg/libvlc-go/v2) and on [GoDoc](https://godoc.org/github.com/adrg/libvlc-go/v2).

![libvlc-go examples](https://raw.githubusercontent.com/adrg/adrg.github.io/master/assets/projects/libvlc-go/libvlc-go-examples.jpg)

Example applications:

* [GUI media player](https://github.com/adrg/libvlc-go-examples/tree/master/v3/gtk3_player)
* [GUI screen recorder](https://github.com/adrg/libvlc-go-examples/tree/master/v3/gtk3_screen_recorder)

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

**libVLC v3.X or later**

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
However, please consider switching to modules.

## Examples

* [GTK 3 media player](https://github.com/adrg/libvlc-go-examples/tree/master/v3/gtk3_player) (using [gotk3](https://github.com/gotk3/gotk3))
* [GTK 3 screen recorder](https://github.com/adrg/libvlc-go-examples/tree/master/v3/gtk3_screen_recorder) (using [gotk3](https://github.com/gotk3/gotk3))
* [GTK 2 media player](https://github.com/adrg/libvlc-go-examples/tree/master/v3/gtk2_player) (using [go-gtk](https://github.com/mattn/go-gtk))
* [GTK 2 screen recorder](https://github.com/adrg/libvlc-go-examples/tree/master/v3/gtk2_screen_recorder) (using [go-gtk](https://github.com/mattn/go-gtk))
* [Basic player usage](https://github.com/adrg/libvlc-go-examples/blob/master/v3/player/player.go)
* [Basic list player usage](https://github.com/adrg/libvlc-go-examples/tree/master/v3/list_player/list_player.go)
* [Handling events](https://github.com/adrg/libvlc-go-examples/tree/master/v3/event_handling/event_handling.go)
* [Retrieve media tracks](https://github.com/adrg/libvlc-go-examples/blob/master/v3/media_tracks/media_tracks.go)
* [Retrieve media information](https://github.com/adrg/libvlc-go-examples/blob/master/v3/media_information/media_information.go)
* [Display screen as player media](https://github.com/adrg/libvlc-go-examples/blob/master/v3/display_screen_media/display_screen_media.go)

Examples for the older version of the API can be found in [here](https://github.com/adrg/libvlc-go-examples/tree/master/v2).

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

## Stargazers over time

[![Stargazers over time](https://starchart.cc/adrg/libvlc-go.svg)](https://starchart.cc/adrg/libvlc-go)

## Contributing

Contributions in the form of pull requests, issues or just general feedback,
are always welcome. See [CONTRIBUTING.MD](CONTRIBUTING.md).

**Contributors**:
[adrg](https://github.com/adrg),
[fenimore](https://github.com/fenimore),
[tarrsalah](https://github.com/tarrsalah),
[danielpellon](https://github.com/danielpellon),
[patknight](https://github.com/patknight),
[sndnvaps](https://github.com/sndnvaps).

## References

For more information see the
[libVLC](https://www.videolan.org/developers/vlc/doc/doxygen/html/group__libvlc.html) documentation.

## License

Copyright (c) 2018 Adrian-George Bostan.

This project is licensed under the [MIT license](https://opensource.org/licenses/MIT).
See [LICENSE](LICENSE) for more details.

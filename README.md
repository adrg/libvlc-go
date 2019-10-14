libvlc-go
=========
[![GoDoc](https://img.shields.io/badge/godoc-reference-blue.svg?style=flat-square)](https://godoc.org/github.com/adrg/libvlc-go)
[![License: MIT](https://img.shields.io/badge/license-MIT-red.svg?style=flat-square)](https://opensource.org/licenses/MIT)
[![Go Report Card](https://goreportcard.com/badge/github.com/adrg/libvlc-go)](https://goreportcard.com/report/github.com/adrg/libvlc-go)

Provides golang bindings for libvlc version 2.X/3.X/4.X. This is a work in
progress and it is not safe for use in a production environment. The current
implementation contains only a small portion of libvlc's functionality.

Full documentation can be found at: https://godoc.org/github.com/adrg/libvlc-go

## Prerequisites
In order to use this project you need to have libvlc-dev installed. On Debian
based distributions it can be installed using apt.
```sh
sudo apt-get install libvlc-dev
```

## Installation
```
go get github.com/adrg/libvlc-go
```

## Build for libvlc < v3.0.0

```
go build -tags legacy
```

## Usage

### Player usage
```go
package main

import (
    "log"

    vlc "github.com/adrg/libvlc-go"
)

func main() {
    // Initialize libvlc. Additional command line arguments can be passed in
    // to libvlc by specifying them in the Init function.
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
    media, err := player.LoadMediaFromURL("https://stream-uk1.radioparadise.com/mp3-32")
    if err != nil {
        log.Fatal(err)
    }
    defer media.Release()

    // Start playing the media.
    err = player.Play()
    if err != nil {
        log.Fatal(err)
    }

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

    <-quit
}
```

### List player usage
```go
package main

import (
    "log"

    vlc "github.com/adrg/libvlc-go"
)

func main() {
    // Initialize libvlc. Additional command line arguments can be passed in
    // to libvlc by specifying them in the Init function.
    if err := vlc.Init("--no-video", "--quiet"); err != nil {
        log.Fatal(err)
    }
    defer vlc.Release()

    // Create a new list player.
    player, err := vlc.NewListPlayer()
    if err != nil {
        log.Fatal(err)
    }
    defer func() {
        player.Stop()
        player.Release()
    }()

    // Create a new media list.
    list, err := vlc.NewMediaList()
    if err != nil {
        log.Fatal(err)
    }
    defer list.Release()

    err = list.AddMediaFromPath("localpath/example1.mp3")
    if err != nil {
        log.Fatal(err)
    }

    err = list.AddMediaFromURL("https://example.com")
    if err != nil {
        log.Fatal(err)
    }

    // Set player media list.
    err = player.SetMediaList(list)
    if err != nil {
        log.Fatal(err)
    }

    // Media files can be added to the list after the list has been added
    // to the player. The player will play these files as well.
    err = list.AddMediaFromPath("localpath/example2.mp3")
    if err != nil {
        log.Fatal(err)
    }

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

    <-quit
}
```

## Contributing

Contributions in the form of pull requests, issues or just general feedback,
are always welcome.

**Contributors**:
[adrg](https://github.com/adrg),
[fenimore](https://github.com/fenimore),
[tarrsalah](https://github.com/tarrsalah),
[danielpellon](https://github.com/danielpellon),
[patknight](https://github.com/patknight)

## References
For more information see [libvlc](https://videolan.org).

## License
Copyright (c) 2018 Adrian-George Bostan.

This project is licensed under the [MIT license](https://opensource.org/licenses/MIT). See LICENSE for more details.

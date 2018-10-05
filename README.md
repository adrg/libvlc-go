libvlc-go
=========
[![GoDoc](http://img.shields.io/badge/godoc-reference-blue.svg?style=flat-square)](https://godoc.org/github.com/adrg/libvlc-go)
[![License: MIT](http://img.shields.io/badge/license-MIT-red.svg?style=flat-square)](http://opensource.org/licenses/MIT)

Provides golang bindings for libvlc version 2.X/3.X/4.X. This is a work in
progress and it is not safe for use in a production environment. The current
implementation contains only a small portion of libvlc's functionality.

Full documentation can be found at: http://godoc.org/github.com/adrg/libvlc-go

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

## Usage

### Player usage
```go
package main

import (
    "log"
    "time"

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
    media, err := player.LoadMediaFromURL("http://stream-uk1.radioparadise.com/mp3-32")
    if err != nil {
        log.Fatal(err)
    }
    defer media.Release()

    // Start playing the media.
    err = player.Play()
    if err != nil {
        log.Fatal(err)
    }

    // Wait some amount of time for the media to start playing.
    // Depends on the version of libvlc. From my tests, libvlc 3.X does not
    // need this delay.
    // TODO: Implement proper callbacks for getting the state of the media.
    time.Sleep(1 * time.Second)

    // If the media played is a live stream the length will be 0.
    length, err := player.MediaLength()
    if err != nil || length == 0 {
        length = 1000 * 60
    }

    time.Sleep(time.Duration(length) * time.Millisecond)
}
```

### List player usage
```go
package main

import (
    "log"
    "time"

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

    err = list.AddMediaFromURL("http://example.com")
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

    // Wait some amount of time for the media to start playing.
    // Depends on the version of libvlc. From my tests, libvlc 3.X does not
    // need this delay.
    // TODO: Implement proper callbacks for getting the state of the media.
    time.Sleep(1 * time.Second)

    // Start playing the media list.
    err = player.Play()
    if err != nil {
        log.Fatal(err)
    }

    time.Sleep(60 * 1000 * time.Millisecond)
}
```

## Contributing

Contributions in the form of pull requests, issues or just general feedback,
are always welcome.

**Contributors**:
[adrg](https://github.com/adrg), [fenimore](https://github.com/fenimore),
[tarrsalah](https://github.com/tarrsalah)

## References
For more information see [libvlc](http://videolan.org).

## License
Copyright (c) 2018 Adrian-George Bostan.

This project is licensed under the [MIT license](http://opensource.org/licenses/MIT). See LICENSE for more details.

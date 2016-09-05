libvlc-go
=========
[![GoDoc](http://img.shields.io/badge/godoc-reference-blue.svg?style=flat-square)](https://godoc.org/github.com/adrg/libvlc-go)
[![License: MIT](http://img.shields.io/badge/license-MIT-red.svg?style=flat-square)](http://opensource.org/licenses/MIT)

Provides golang bindings for libvlc version 2.X. This is a work in progress
and it is not safe for use in a production environment. The current
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

### Usage
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

	// Create a new player
	player, err := vlc.NewPlayer()
	if err != nil {
		log.Fatal(err)
	}
	defer func() {
		player.Stop()
		player.Release()
	}()

	// Set player media. The second parameter of the method specifies if
	// the media resource is local or remote.
	// err = player.SetMedia("localPath/test.mp4", true)
	err = player.SetMedia("http://stream-uk1.radioparadise.com/mp3-32", false)
	if err != nil {
		log.Fatal(err)
	}

	// Play
	err = player.Play()
	if err != nil {
		log.Fatal(err)
	}

	// Wait some amount of time for the media to start playing
	// TODO: Implement proper callbacks for getting the state of the media
	time.Sleep(1 * time.Second)

	// If the media played is a live stream the length will be 0
	length, err := player.MediaLength()
	if err != nil || length == 0 {
		length = 1000 * 60
	}

	time.Sleep(time.Duration(length) * time.Millisecond)
}
```

## References
For more information see [libvlc](http://videolan.org).

## License
Copyright (c) 2015 Adrian-George Bostan.

This project is licensed under the [MIT license](http://opensource.org/licenses/MIT). See LICENSE for more details.

wallbase
========

The wallbase package facilitates requests to [wallbase.cc][]. Useful until the
public API is released.

[wallbase.cc]: http://wallbase.cc/

Documentation
-------------

Documentation provided by GoDoc.

- [wallbase][]: implements search and download functions for [wallbase.cc][].

[wallbase]: http://godoc.org/github.com/mewmew/playground/archive/wallbase

walls
=====

walls updates the desktop wallpaper at specified time intervals.

It performs a search for wallpapers on [wallbase.cc] based on the provided
search query. The wallpaper result order is random.

For persistent storage a custom download directory can be specified.

Installation
------------

	go get github.com/mewmew/playground/archive/wallbase/cmd/walls

Dependencies
------------

The `hsetroot` command is used to set the wallpaper.

Command documentation
---------------------

Command documentation provided by GoDoc.

- [walls][]

[walls]: http://godoc.org/github.com/mewmew/playground/archive/wallbase/cmd/walls

Usage
-----

	walls [OPTION]... QUERY

Flags:

	-o (default="/tmp/wallbase")
		Output directory.
	-res (default="")
		Screen resolution (ex. "1920x1080").
	-t (default="30m")
		Timeout interval between updates.
	-v (default=false)
		Verbose.

Examples
--------

1. Search for "nature waterfall" and update active wallpaper each 10s.

		wallbase -t 10s nature waterfall

2. Search for "nature" and store each wallpaper in "download/".

		wallbase -t 0s -o download/ nature

3. Search for "nature" wallpapers with the screen resolution 1920x1080.

		wallbase -res 1920x1080 nature

public domain
-------------

This code is hereby released into the *[public domain][]*.

[public domain]: https://creativecommons.org/publicdomain/zero/1.0/

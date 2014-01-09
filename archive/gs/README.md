gs
==

The gs package implements an unofficial API for [grooveshark.com][].

[grooveshark.com]: http://grooveshark.com/

Documentation
-------------

Documentation provided by GoDoc.

- [gs][]

[gs]: http://godoc.org/github.com/mewmew/playground/archive/gs

aquarium
========

aquarium is a backup utility for grooveshark users.

It creates a list of all songs in a user's collection, favorites and playlists.

Note that only the artist name and song title is stored, not the actual audio
file.

Installation
------------

	go get github.com/mewmew/playground/archive/gs/cmd/aquarium

Usage
-----

	aquarium USERNAME

Examples
--------

1. Create a backup of testuser's song collection, favorites and playlists.

		aquarium testuser

public domain
-------------

This code is hereby released into the *[public domain][]*.

[public domain]: https://creativecommons.org/publicdomain/zero/1.0/

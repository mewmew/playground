gitsync
=======

This tool keeps forked git repositories in sync with their parents. It does so
by locating the repositories of provided usernames and organizations. Then it
creates a shell script which will clone all repository forks, pull changes from
their parens and push those changes to the forked repository.

Installation
------------

	go get github.com/mewmew/playground/cmd/gitsync

Usage
-----

	gitsync USER...

Examples
--------

	gitsync mewbak > sync.sh
	chmod +x sync.sh
	./sync.sh

public domain
-------------

This code is hereby released into the *[public domain][]*.

[public domain]: https://creativecommons.org/publicdomain/zero/1.0/

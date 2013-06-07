WIP
---

This project is a *work in progress*. The implementation is *incomplete* and
subject to change. The documentation can be inaccurate.

breyting
========

[breyting][] is Icelandic and translates to change or alteration. This
application will monitor a set of pages and check for changes at a specified
time interval.

It will also monitor the config file of said pages, so the application doesn't
have to be restarted after changing the set of pages to be monitored.

[breyting]: https://en.wiktionary.org/wiki/breyting

Installation
------------

	$ go get github.com/mewmew/playground/breyting/cmd/breyting
	$ mkdir -p ~/.config/breyting/
	$ cp $GOPATH/src/github.com/mewmew/playground/breyting/breyting.ini ~/.config/breyting/

Examples
--------

	# Specify which pages (URL and CSS selector) to monitor in the config file.

	$ breyting -http=:4000
	$ firefox http://localhost:4000

public domain
-------------

This code is hereby released into the *[public domain][]*.

[public domain]: https://creativecommons.org/publicdomain/zero/1.0/

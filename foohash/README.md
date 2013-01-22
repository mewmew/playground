WIP
---

This project is a *work in progress*. The implementation is *incomplete* and
subject to change. The documentation can be inaccurate.

foohash
=======

foohash recovers passwords from hashes.

Installation
------------

	$ go get github.com/mewmew/foohash/cmd/foohash

Usage
-----

	foohash [OPTION]... [HASH]...

Flags:

	-w (default="")
		Wordlist path.
	-wr (default=true)
		Wordlist attack, using regular (unmodified) words.
	-wt (default=true)
		Wordlist attack, using titled words.
	-wu (default=true
		Wordlist attack, using upper case words.
	-wl (default=true)
		Wordlist attack, using leet speak words.
	-wn (default=true)
		Wordlist attack, using words with number suffixes.
	-sp (default="")
		Salt prefix.
	-ss (default="")
		Salt suffix.
	-g (default=false)
		Google.

public domain
-------------

This code is hereby released into the *[public domain][]*.

[public domain]: https://creativecommons.org/publicdomain/zero/1.0/

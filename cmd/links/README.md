links
=====

This tool examines HTML files and reports invalid links.

Installation
------------

	go get github.com/mewmew/playground/cmd/links

Usage
-----

	links PATH...

Examples
--------

	cd testdata/
	links *.html
	// Output:
	// invalid fragment id to "b.html#a" in file "a.html".
	// invalid link target "c.html" in file "b.html".

[a.html]: https://raw.github.com/mewmew/playground/master/cmd/links/testdata/a.html
[b.html]: https://raw.github.com/mewmew/playground/master/cmd/links/testdata/b.html

public domain
-------------

This code is hereby released into the *[public domain][]*.

[public domain]: https://creativecommons.org/publicdomain/zero/1.0/

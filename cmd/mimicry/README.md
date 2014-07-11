mimicry
=======

This tool creates a git repository which mimics an image using a contribution
history of carefully crafted commit dates. It expects a 51x7 image with a
transparent background.

Installation
------------

	go get github.com/mewmew/playground/cmd/mimicry

Usage
-----

	mimicry IMAGE

Examples
--------

The mimicry command was used in a second attempt to create the [hello][]
repository. Below is a screenshot of its contribution history as of 2014-07-11.

	mimicry hello.png

![Screenshot - first try](https://raw.github.com/mewmew/playground/master/cmd/mimicry/hello world.png)

[hello]: https://github.com/yumpie2/hello

public domain
-------------

This code is hereby released into the *[public domain][]*.

[public domain]: https://creativecommons.org/publicdomain/zero/1.0/

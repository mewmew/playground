Partial implementation
----------------------

This implementation is only *partially complete*. While useable it is still
missing a lot of features.

pcap
====

This package provides support for reading pcap files.

Documentation
-------------

Documentation provided by GoDoc.

	- [pcap][]

[pcap]: http://godoc.org/github.com/mewmew/playground/archive/pcap

Installation
------------

	go get github.com/mewmew/playground/archive/pcap

pcapsulate
==========

pcapsulate encapsulates the provided files as packets in a pcap file.

Installation
------------

	go get github.com/mewmew/playground/archive/pcap/cmd/pcapsulate

Usage
-----

	pcapsulate [FILE]...

Flags:

	-o (default="pcapsulate.pcap")
		Output path.

public domain
-------------

This code is hereby released into the *[public domain][]*.

[public domain]: https://creativecommons.org/publicdomain/zero/1.0/

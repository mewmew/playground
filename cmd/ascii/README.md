# ascii

This tool examines files and reports non-ascii characters.

## Installation

	go get github.com/mewmew/playground/cmd/ascii

## Usage

	ascii [OPTION]... PATH...

Flags:

	-v (default=false)
		Verbose.

## Examples

	echo "Hello, 世界" > foo.txt
	ascii foo.txt
	// Output:
	// foo.txt (1:7) - non-ascii character '世'

## Public domain

The source code and any original content of this repository is hereby released into the [public domain].

[public domain]: https://creativecommons.org/publicdomain/zero/1.0/

# usagen

`usagen` generates usage documentation for a given command. It does so by executing the command with the "--help" flag and parsing the output.

## Installation

    go get github.com/mewmew/playground/cmd/usagen

## Usage

    usagen CMD

## Examples

    $ usagen ll2dot
    // Usage:
    //
    //     ll2dot [OPTION]... FILE...
    //
    // Flags:
    //
    //    -f=false:   Force overwrite existing graph directories.
    //    -funcs="":  Comma separated list of functions to parse (e.g. "foo,bar").
    //    -img=false: Generate an image representation of the CFG.
    //    -q=false:   Suppress non-error messages.
    package main

## Public domain

The source code and any original content of this repository is hereby released into the [public domain].

[public domain]: https://creativecommons.org/publicdomain/zero/1.0/

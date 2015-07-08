# thesound

This tool extracts the original sound from videos. The raw audio stream is
copied (not converted).

## Dependencies

`thesound` uses [ffprobe][] to determine the audio codec, and [ffmpeg][] to
extract the sound.

[ffprobe]: http://www.ffmpeg.org/ffprobe.html
[ffmpeg]: http://www.ffmpeg.org/ffmpeg.html

## Installation

	go get github.com/mewmew/playground/cmd/thesound

## Usage

	thesound PATH...

## Examples

	thesound "Birdy - Not About Angels.mp4"
	// Output:
	// Created "Birdy - Not About Angels.aac".

## Public domain

The source code and any original content of this repository is hereby released into the [public domain].

[public domain]: https://creativecommons.org/publicdomain/zero/1.0/

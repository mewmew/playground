# wallhaven

The wallhaven package implements an unofficial API for [wallhaven.cc](http://wallhaven.cc/).

## Documentation

Documentation provided by GoDoc.

- [wallhaven][]: implements search and download functionality for wallhaven.cc.

[wallhaven]: http://godoc.org/github.com/mewmew/playground/archive/wallhaven

# walls

walls is a tool which downloads wallpapers from wallhaven.cc.

It performs a search for wallpapers on wallbase.cc based on the given search query and search options.

For persistent storage, a custom download directory may be specified.

## Installation

	go get github.com/mewmew/playground/archive/wallhaven/cmd/walls

## Usage

	walls [OPTION]... QUERY

Flags:

	-cats value
		Wallpaper categories (general,anime,people)
	-n int
		Download at most n wallpapers (0 = infinite)
	-o string
		Output directory (default "/tmp/walls")
	-order value
		Sorting order (asc,desc)
	-page value
		Result page number
	-purity value
		Purity modes (SFW,sketchy,NSFW)
	-ratios value
		Aspect ratios (4x3,5x4,16x9,16x10,21x9,32x9,48x9)
	-res value
		Screen resolutions (1024x768,1280x800,1366x768,1280x960,1440x900,1600x900,1280x1024,1600x1200,1680x1050,1920x1080,1920x1200,2560x1440,2560x1600,3840x1080,5760x1080,3840x2160)
	-sorting value
		Sorting method (relevance,random,date_added,views,favorites)

## Examples

1. Download a random wallpaper and set it as desktop background.

		walls -n 1 -sorting random | xargs -I {} feh --bg-fill "{}"

2. Download wallpapers matching the search query *nature* to the `download/` directory.

		walls -o download/ nature

3. Download at most 100 SFW wallpapers with a screen resolution 1920x1080.

		wallbase -n 100 -purity SFW -res 1920x1080

## Public domain

The source code and any original content of this repository is hereby released into the [public domain].

[public domain]: https://creativecommons.org/publicdomain/zero/1.0/

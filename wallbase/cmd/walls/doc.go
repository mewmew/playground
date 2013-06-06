/*

walls updates the desktop wallpaper at specified time intervals.

It performs a search for wallpapers on http://wallbase.cc/ based on the provided
search query. The wallpaper result order is random.

For persistent storage a custom download directory can be specified.

Installation:

	go get github.com/mewmew/wallbase/cmd/walls

Dependencies:

The `hsetroot` command is used to set the wallpaper.

Usage:

	walls [OPTION]... QUERY

Flags:

	-o (default="/tmp/wallbase")
		Output directory.
	-t (default="30m")
		Timeout interval between updates.

Examples:

1. Search for "nature waterfall" and update active wallpaper each 10s.

	wallbase -t 10s nature waterfall

2. Search for "nature" and store each wallpaper in "download/".

	wallbase -t 0s -o download/ nature

*/
package main

package main

import "bytes"
import "image"
import "image/png"

import "github.com/mewmew/playground/eg"

// getImage downloads and returns an image for the challenge.
func getImage() (img image.Image, err error) {
	buf, err := eg.Get("http://www.enigmagroup.org/missions/programming/8/image.php")
	if err != nil {
		return nil, err
	}
	r := bytes.NewBuffer(buf)
	img, err = png.Decode(r)
	if err != nil {
		return nil, err
	}
	return img, nil
}

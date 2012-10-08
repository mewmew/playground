package main

import "os"

import "github.com/mewkiz/pkg/imgutil"
import "github.com/mewmew/playground/pic"

// dump converts each sub image into a PNG file, whose name is based on the hash
// of it's content.
func dump(subs []pic.SubImager) (err error) {
	for _, sub := range subs {
		hashSum, err := getHashSum(sub)
		if err != nil {
			return err
		}
		imgName := hashSum + ".png"
		fr, err := os.Open(imgName)
		if err == nil {
			fr.Close()
			continue
		}
		err = imgutil.WriteFile(imgName, sub)
		if err != nil {
			return err
		}
	}
	return nil
}

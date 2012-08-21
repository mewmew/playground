package main

import "fmt"
import "os"

import "github.com/mewkiz/pkg/pngutil"
import "github.com/mewmew/playground/pic"

// charCount keeps track of the number of times an image content hash has been
// encountered.
var charCount = map[string]int{}

// dump converts each sub image into a PNG file, whose name is based on the hash
// of it's content.
func dump(subs []pic.SubImager) (err error) {
	for _, sub := range subs {
		hashSum, err := getHashSum(sub)
		if err != nil {
			return err
		}
		charCount[hashSum]++
		imgName := hashSum + ".png"
		fr, err := os.Open(imgName)
		if err == nil {
			fr.Close()
			continue
		}
		err = pngutil.WriteFile(imgName, sub)
		if err != nil {
			return err
		}
	}
	return nil
}

// printCount prints the total number of times an image content hash has been
// encountered.
func printCount() {
	for hashSum, count := range charCount {
		fmt.Printf("%4d - %s\n", count, hashSum)
	}
}

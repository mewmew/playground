package main

import "errors"
import "flag"
import "fmt"
import "log"
import "os"

import "github.com/mewkiz/pkg/imgutil"
import "github.com/mewmew/playground/eg"
import "github.com/mewmew/playground/pic"

func init() {
	flag.StringVar(&eg.FieldV4, "f", "", "Set enigmafiedV4 cookie value.")
	flag.StringVar(&eg.PhpSessid, "p", "", "Set PHPSESSID cookie value.")
	flag.Parse()
}

func main() {
	if !eg.HasSession() {
		flag.Usage()
		os.Exit(1)
	}
	err := prog8()
	if err != nil {
		log.Fatalln(err)
	}
}

// prog8 translates the captcha into text and submits the answer.
//
// 1) Download:
//    - Download the image.
// 2) Sort:
//    - Sort it's lines based on the background gradient.
// 3) Mono:
//    - Convert the image to monochrome, using a transparent background and
//      black foreground.
// 4) Crop:
//    - Crop the image.
// 5) Split:
//    - Split the image into several horizontal sub images, which also have been
//       cropped.
// 6) Translate:
//    - Translate each sub image back into a character by compare them against a
//      stored version.
// 7) Submit answer.
//    - Submit the answer.
func prog8() (err error) {
	img, err := getImage()
	if err != nil {
		return err
	}
	orig, ok := img.(pic.SubImager)
	if !ok {
		return errors.New("image contains no SubImage method.")
	}
	err = imgutil.WriteFile("01_orig.png", orig)
	if err != nil {
		return err
	}
	sorted := getSorted(orig)
	err = imgutil.WriteFile("02_sort.png", sorted)
	if err != nil {
		return err
	}
	mono := getMono(sorted)
	err = imgutil.WriteFile("03_mono.png", mono)
	if err != nil {
		return err
	}
	crop := pic.Crop(mono)
	err = imgutil.WriteFile("04_crop.png", crop)
	if err != nil {
		return err
	}
	subs := pic.HSubs(crop)
	for subNum, sub := range subs {
		err = imgutil.WriteFile(fmt.Sprintf("05_sub_%d.png", subNum), sub)
		if err != nil {
			return err
		}
	}
	text, err := translate(subs)
	if err != nil {
		return err
	}
	err = submitAnswer(text)
	if err != nil {
		return err
	}
	return nil
}

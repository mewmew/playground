package main

import dbg "fmt"
import "errors"
import "flag"
import "fmt"
///import "log"
import "os"
import "time"

import "github.com/mewkiz/pkg/pngutil"
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
   for i := 0; i < 1000; i++ {
      err := captcha2()
      if err != nil {
         ///log.Fatalln(err)
         dbg.Println(err)
      }
      time.Sleep(1 * time.Second)
   }
}

// captcha2 translates the captcha into text and submits the answer.
//
// 1) Download:
//    - Download the image.
// 2) Mono:
//    - Convert the image to monochrome, using a transparent background and
//      black foreground.
// 3) Crop:
//    - Crop the image.
// 4) Split:
//    - Split the image into several horizontal sub images, which also have been
//       cropped.
// 5) Translate:
//    - Translate each sub image back into a character by compare them against a
//      stored version.
// 6) Submit answer.
//    - Submit the answer.
func captcha2() (err error) {
   img, err := getImage()
   if err != nil {
      return err
   }
   orig, ok := img.(pic.SubImager)
   if !ok {
      return errors.New("image contains no SubImage method.")
   }
   err = pngutil.WriteFile("01_orig.png", orig)
   if err != nil {
      return err
   }
   mono := getMono(orig)
   err = pngutil.WriteFile("02_mono.png", mono)
   if err != nil {
      return err
   }
   crop := pic.Crop(mono)
   err = pngutil.WriteFile("03_crop.png", crop)
   if err != nil {
      return err
   }
   subs := pic.HSubs(crop)
   for subNum, sub := range subs {
      err = pngutil.WriteFile(fmt.Sprintf("04_sub_%d.png", subNum), sub)
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

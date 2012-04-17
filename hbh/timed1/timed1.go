package main

import "github.com/mewmew/playground/hbh"
import "github.com/mewmew/playground/hbh/parse"
import "github.com/mewmew/playground/str"

import "encoding/base64"
import "errors"
import "flag"
import "fmt"
import "log"
import "net/url"
import "strings"

func init() {
   flag.StringVar(&hbh.PhpSessid, "p", "", "Set PHPSESSID cookie value.")
   flag.StringVar(&hbh.FusionUser, "f", "", "Set fusion_user cookie value.")
   flag.Parse()
}

func main() {
   if !hbh.HasSession() {
      flag.Usage()
      return
   }
   err := timed1()
   if err != nil {
      log.Fatalln(err)
   }
}

func timed1() (err error) {
   enc, err := getEncStr()
   if err != nil {
      return err
   }
   buf, err := base64.StdEncoding.DecodeString(enc)
   if err != nil {
      return err
   }
   dec := string(buf)
   fmt.Println("decoded string:", dec)
   err = submitSolution(dec)
   if err != nil {
      return err
   }
   return nil
}

func getEncStr() (enc string, err error){
   rawUrl := "http://www.hellboundhackers.org/challenges/timed/timed1/index.php"
   text, err := hbh.Get(rawUrl)
   if err != nil {
      return "", err
   }
   pos := str.IndexAfter(text, "Decrypt the following random string: ")
   if pos == -1 {
      return "", errors.New("base64 encoded string start not found.")
   }
   encLen := strings.IndexRune(text[pos:], ' ')
   if encLen == -1 {
      return "", errors.New("base64 encoded string end not found.")
   }
   return text[pos:pos + encLen], nil
}

func submitSolution(dec string) (err error) {
   rawUrl := "http://www.hellboundhackers.org/challenges/timed/timed1/index.php?b64="
   text, err := hbh.Get(rawUrl + url.QueryEscape(dec))
   if err != nil {
      return err
   }
   err = parse.Html(text, &parse.Search{Tag: "div", Class: "open_table"})
   if err != nil {
      return err
   }
   return nil
}

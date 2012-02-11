package main

import "hbh"
import "parse"
import "str"

import "errors"
import "flag"
import "fmt"
import "log"
import "strconv"
import "strings"
import "unicode"

func init() {
   flag.StringVar(&hbh.PhpSessid, "p", "", "Set PHPSESSID cookie value.")
   flag.StringVar(&hbh.FusionUser, "f", "", "Set fusion_user cookie value.")
}

func main() {
   flag.Parse()
   err := timed1()
   if err != nil {
      log.Fatalln(err)
   }
}

func timed1() (err error) {
   md5, err := getMd5Hash()
   var sum int
   for _, r := range md5 {
      if unicode.IsNumber(r) {
         n, err := strconv.Atoi(string(r))
         if err != nil {
            return err
         }
         sum += n
      }
   }
   fmt.Println("sum:", sum)
   err = submitSolution(sum)
   if err != nil {
      return err
   }
   return nil
}

func getMd5Hash() (md5 string, err error){
   rawUrl := "http://www.hellboundhackers.org/challenges/timed/timed2/index.php"
   text, err := hbh.Get(rawUrl)
   if err != nil {
      return "", err
   }
   pos := str.IndexAfter(text, "You have <strong>two</strong> seconds to answer this challenge and your string is: ")
   if pos == -1 {
      return "", errors.New("md5 hash start not found.")
   }
   md5Len := strings.Index(text[pos:], "<br />")
   if md5Len == -1 {
      return "", errors.New("md5 hash end not found.")
   }
   return text[pos:pos + md5Len], nil
}

func submitSolution(sum int) (err error) {
   rawUrl := "http://www.hellboundhackers.org/challenges/timed/timed2/index.php?check"
   text, err := hbh.Post(rawUrl, fmt.Sprintf("ans=%d&submit=Check", sum))
   if err != nil {
      return err
   }
   err = parse.Html(text, &parse.Search{Tag: "div", Class: "open_table"})
   if err != nil {
      return err
   }
   return nil
}

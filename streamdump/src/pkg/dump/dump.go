package dump

import "errors"
import "fmt"
import "io/ioutil"
import "net/http"
import "os"
import "os/exec"
import "path"
import "strings"

func Url(rawUrl string) (err error) {
   r, err := http.Get(rawUrl)
   if err != nil {
      return err
   }
   defer r.Body.Close()
   buf, err := ioutil.ReadAll(r.Body)
   if err != nil {
      return err
   }
   text := string(buf)
   pos := strings.Index(text, "rtmpe://")
   if pos == -1 {
      pos = strings.Index(text, "rtmp://")
      if pos == -1 {
         return errors.New("stream URL start not found.")
      }
   }
   streamUrlLen := strings.Index(text[pos:], ",")
   if streamUrlLen == -1 {
      return errors.New("stream URL end not found.")
   }
   streamUrl := text[pos:pos + streamUrlLen]
   err = Stream(streamUrl)
   if err != nil {
      return err
   }
   return nil
}

func Stream(streamUrl string) (err error) {
   fmt.Println("dumping:", streamUrl)
   cmd := exec.Command("rtmpdump", "-r", streamUrl, "-o", path.Base(streamUrl))
   cmd.Stderr = os.Stderr
   err = cmd.Run()
   if err != nil {
      return err
   }
   return nil
}

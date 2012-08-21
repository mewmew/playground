package main

import dbg "fmt"
import "fmt"

import "github.com/mewmew/playground/eg"

// submitAnswer submits the answer to enigmagroup's server.
func submitAnswer(text string) (err error) {
	dbg.Println("answer:", text)
	data := fmt.Sprintf("answer=%s&submit=true", text)
	buf, err := eg.Post("http://www.enigmagroup.org/missions/captcha/2/image.php", data)
	dbg.Println("buf:", string(buf))
	if err != nil {
		return err
	}
	return nil
}

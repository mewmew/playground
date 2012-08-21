package main

import dbg "fmt"
import "errors"
import "fmt"

import "github.com/mewmew/playground/eg"

// submitAnswer submits the answer to enigmagroup's server.
func submitAnswer() (err error) {
	if circleValue == 0 {
		return errors.New("unable to calculate value of circle.")
	}
	if heartValue == 0 {
		return errors.New("unable to calculate value of heart.")
	}
	if squareValue == 0 {
		return errors.New("unable to calculate value of square.")
	}
	if starValue == 0 {
		return errors.New("unable to calculate value of star.")
	}
	if triangleValue == 0 {
		return errors.New("unable to calculate value of triangle.")
	}
	data := fmt.Sprintf("circle=%d&heart=%d&square=%d&star=%d&triangle=%d&submit=Submit", circleValue, heartValue, squareValue, starValue, triangleValue)
	buf, err := eg.Post("http://www.enigmagroup.org/missions/programming/9/", data)
	dbg.Println("buf:", string(buf))
	if err != nil {
		return err
	}
	return nil
}

package main

import (
	"log"
)

func main() {
	err := play()
	if err != nil {
		log.Fatalln(err)
	}
}

func play() (err error) {
	return nil
}

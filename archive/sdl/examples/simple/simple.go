package main

import (
	"log"

	"github.com/mewmew/playground/archive/sdl"
)

func main() {
	err := simple()
	if err != nil {
		log.Fatalln(err)
	}
}

func simple() (err error) {
	err = sdl.Init(sdl.InitVideo)
	if err != nil {
		return err
	}
	defer sdl.Quit()
	win, err := sdl.CreateWindow("foo", 300, 300, 640, 480, 0)
	if err != nil {
		return err
	}
	win.Update()
	select {}
}

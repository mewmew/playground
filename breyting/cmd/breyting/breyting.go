package main

import (
	"flag"
	dbg "fmt"
	"log"
	"os"
	"time"

	"github.com/mewmew/playground/breyting/conf"
)

// confPath is the path to the config file. The default is:
//    $HOME/.config/breyting/breyting.ini
var confPath string

func init() {
	defaultConfPath := os.Getenv("HOME") + "/.config/breyting/breyting.ini"
	flag.StringVar(&confPath, "conf", defaultConfPath, "Path to config file.")
}

func main() {
	flag.Parse()
	err := breyting()
	if err != nil {
		log.Fatalln(err)
	}
}

// breyting loads the config file and adds a watcher to it. Each time the config
// file is reloaded, it will spawn one watcher for each page.
func breyting() (err error) {
	dbg.Println("using:", confPath)
	err = conf.Reload(confPath)
	if err != nil {
		return err
	}
	go conf.Watch(confPath)
	/// ### [ tmp ] ###
	///   - server.Listen() should be added here.
	/// ### [/ tmp ] ###
	time.Sleep(1 * time.Hour)
	return nil
}

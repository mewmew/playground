// inject is a tool which satisfies a TED talk addiction.
package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"log"

	"github.com/mewmew/playground/archive/ted"
)

func main() {
	flag.Parse()
	for _, path := range flag.Args() {
		talks, err := parse(path)
		if err != nil {
			log.Fatalln(err)
		}
		inject(talks)
	}
}

// inject outputs a wget script for downloading TED talks.
func inject(talks []ted.Talk) {
	for _, talk := range talks {
		date := talk.Date.Format("2006-01")
		fmt.Printf("wget -O \"%s - %s (%s) [%v].mp4\" %s\n", date, talk.Title, talk.Event, talk.Duration, talk.Download)
	}
}

// parse parses the provided JSON file of TED talks.
func parse(path string) (talks []ted.Talk, err error) {
	buf, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(buf, &talks)
	if err != nil {
		return nil, err
	}
	return talks, nil
}

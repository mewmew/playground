// order is a tool which satisfies a TED talk addiction in an ordered fashion.
package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"

	"github.com/mewmew/playground/archive/ted"
)

func main() {
	flag.Parse()
	for _, path := range flag.Args() {
		talks, err := parse(path)
		if err != nil {
			log.Fatalln(err)
		}
		order(talks)
	}
}

// order outputs one wget script for downloading TED talks per release month.
func order(talks []ted.Talk) {
	var prev string
	var w io.Writer
	for _, talk := range talks {
		date := talk.Date.Format("2006-01")
		if prev == "" || prev != date {
			f, err := os.Create(fmt.Sprintf("%s.sh", date))
			if err != nil {
				log.Fatalln(err)
			}
			err = f.Chmod(0755)
			if err != nil {
				log.Fatalln(err)
			}
			defer f.Close()
			w = f
			fmt.Fprintf(w, "mkdir -p %q\n", date)
		}
		fmt.Fprintf(w, "wget -O \"%s/%s (%s) [%v].mp4\" %s\n", date, talk.Title, talk.Event, talk.Duration, talk.Download)
		prev = date
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

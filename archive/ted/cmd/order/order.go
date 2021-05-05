// order is a tool which satisfies a TED talk addiction in an ordered fashion.
package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"

	"github.com/mewmew/playground/archive/ted"
	"github.com/pkg/errors"
)

func main() {
	flag.Parse()
	for _, path := range flag.Args() {
		talks, err := parse(path)
		if err != nil {
			log.Fatalf("%+v", err)
		}
		if err := order(talks); err != nil {
			log.Fatalf("%+v", err)
		}
	}
}

// files tracks open files.
var files = make(map[string]*os.File)

// order outputs one wget script for downloading TED talks per release month.
func order(talks []ted.Talk) error {
	for _, talk := range talks {
		date := talk.Date.Format("2006-01")
		shPath := fmt.Sprintf("%s.sh", date)
		file, ok := files[shPath]
		if !ok {
			fmt.Printf("Creating %q\n", shPath)
			f, err := os.Create(shPath)
			if err != nil {
				return errors.WithStack(err)
			}
			if err := f.Chmod(0755); err != nil {
				return errors.WithStack(err)
			}
			file = f
			files[shPath] = file
			fmt.Fprint(file, "#!/bin/bash\n\n")
			fmt.Fprintf(file, "mkdir -p %q\n\n", date)
		}
		if len(talk.Download) > 0 {
			fmt.Fprintf(file, "wget -O \"%s/%s (%s) [%v].mp4\" %q\n", date, escape(talk.Title), escape(talk.Event), talk.Duration, talk.Download)
		} else {
			fmt.Fprintf(file, "# firefox %q\n", talk.Link)
		}
	}
	for _, file := range files {
		if err := file.Close(); err != nil {
			return errors.WithStack(err)
		}
	}
	return nil
}

// parse parses the provided JSON file of TED talks.
func parse(path string) (talks []ted.Talk, err error) {
	buf, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, errors.WithStack(err)
	}
	err = json.Unmarshal(buf, &talks)
	if err != nil {
		return nil, errors.WithStack(err)
	}
	return talks, nil
}

// escape escapes double quotes.
func escape(s string) string {
	return strings.ReplaceAll(s, `"`, `\"`)
}

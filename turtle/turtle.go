// Package turtle is meant to facilitate the learning of Japanese.
package turtle

import (
	"log"

	"github.com/mewkiz/pkg/goutil"
	"github.com/mewkiz/pkg/pathutil"
)

func init() {
	// Locate base data directory.
	dataDir, err := goutil.SrcDir("github.com/mewmew/playground/turtle/data")
	if err != nil {
		log.Fatalln(err)
	}
	data = pathutil.Base(dataDir)
}

// data is the base data directory.
var data pathutil.Base

// When Verbose is set to true, verbose output is enabled.
var Verbose bool

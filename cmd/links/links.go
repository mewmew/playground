// links examines HTML files and reports invalid links.
package main

import (
	"flag"
	"fmt"
	"log"
	"net/url"
	"os"
	"path"
	"path/filepath"

	"golang.org/x/net/html"
)

var flagVerbose bool

func init() {
	flag.BoolVar(&flagVerbose, "v", false, "Verbose.")
	flag.Usage = usage
}

func usage() {
	fmt.Fprintln(os.Stderr, "Usage: links PATH...")
	fmt.Fprintln(os.Stderr, "Reports invalid links in HTML files.")
	fmt.Fprintln(os.Stderr)
	fmt.Fprintln(os.Stderr, "Invoke links with one or more filenames or directories.")
}

func main() {
	flag.Parse()
	for _, filePath := range flag.Args() {
		if isDir(filePath) {
			err := parseDir(filePath)
			if err != nil {
				log.Fatalln(err)
			}
		} else {
			err := parseFile(filePath)
			if err != nil {
				log.Fatalln(err)
			}
		}
	}
	checkLinks()
}

var documents = make(map[string]*Doc)

type Doc struct {
	Ids  map[string]bool
	URLs []*url.URL
}

func parseFile(filePath string) (err error) {
	ext := path.Ext(filePath)
	_, ok := whitelist[ext]
	if !ok {
		if flagVerbose {
			log.Printf("ignoring file %q with extension %q.\n", filePath, ext)
		}
		return nil
	}
	fr, err := os.Open(filePath)
	if err != nil {
		return err
	}
	defer fr.Close()

	doc, err := html.Parse(fr)
	if err != nil {
		return err
	}

	documents[filePath] = &Doc{
		Ids:  make(map[string]bool),
		URLs: make([]*url.URL, 0),
	}

	document := documents[filePath]

	var f func(*html.Node) error
	f = func(n *html.Node) (err error) {
		if n.Type == html.ElementNode {
			for _, attr := range n.Attr {
				if attr.Key == "id" {
					document.Ids[attr.Val] = true
				}
				if n.Data == "a" && attr.Key == "href" {
					link, err := url.Parse(attr.Val)
					if err != nil {
						return err
					}
					document.URLs = append(document.URLs, link)
				}
			}
		}
		for c := n.FirstChild; c != nil; c = c.NextSibling {
			err = f(c)
			if err != nil {
				return err
			}
		}
		return nil
	}
	err = f(doc)
	if err != nil {
		return err
	}

	return nil
}

func checkLinks() {
	for documentPath, document := range documents {
		for _, u := range document.URLs {
			if u.IsAbs() {
				if flagVerbose {
					log.Println("ignoring absolute URL:", u)
					continue
				}
			}
			target, ok := documents[u.Path]
			if !ok {
				fmt.Fprintf(os.Stderr, "invalid link target %q in file %q.\n", u.Path, documentPath)
				continue
			}
			if len(u.Fragment) == 0 {
				continue
			}
			_, ok = target.Ids[u.Fragment]
			if !ok {
				fmt.Fprintf(os.Stderr, "invalid fragment id to %q in file %q.\n", u, documentPath)
				continue
			}
		}
	}
}

func isDir(path string) bool {
	fi, err := os.Stat(path)
	if err != nil {
		return false
	}
	return fi.IsDir()
}

func parseDir(dir string) (err error) {
	err = filepath.Walk(dir, walk)
	if err != nil {
		return err
	}
	return nil
}

// whitelist contains a list of all extensions believed to be HTML files.
var whitelist = map[string]bool{
	".html": true,
}

func walk(path string, info os.FileInfo, err error) error {
	if err != nil {
		return err
	}
	if info.Mode().IsRegular() {
		// only parse whitelisted extensions.
		ext := filepath.Ext(path)
		_, ok := whitelist[ext]
		if !ok {
			return nil
		}

		err = parseFile(path)
		if err != nil {
			return err
		}
	}
	return nil
}

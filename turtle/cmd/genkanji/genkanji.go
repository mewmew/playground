// As you will notice this is kind of a hack, but it works :)
//
// Parse table of Jōyō kanji [1]. Use the following steps to prepare the indata
// for this application:
//
//    1. view-source:https://en.wikipedia.org/wiki/List_of_j%C5%8Dy%C5%8D_kanji
//    2. Copy from "<table>" to "</table>".
//    3. Remove the first "<tr>...<tr>"
//    4. Add "<tr></tr>" after the last kanji.

package main

import (
	"bytes"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"sort"
	"strconv"

	"code.google.com/p/go.net/html"
)

func main() {
	flag.Parse()
	for _, filePath := range flag.Args() {
		err := genkanji(filePath)
		if err != nil {
			log.Fatalln(err)
		}
	}
	sort.Sort(GradeOrder(kanjis))
	for _, kanji := range kanjis {
		fmt.Printf("%#v\n", kanji)
	}
}

func genkanji(filePath string) (err error) {
	buf, err := ioutil.ReadFile(filePath)
	if err != nil {
		return err
	}

	doc, err := html.Parse(bytes.NewReader(buf))
	if err != nil {
		return err
	}
	err = f(doc)
	if err != nil {
		return err
	}

	return nil
}

var tdCount int

type Kanji struct {
	Rune           string
	OldRune        string
	Radical        string
	Grade          int
	Meaning        string
	Pronounciation string
}

var kanjis = make([]*Kanji, 0, 2136)

var kanji *Kanji

func f(n *html.Node) (err error) {
	if n.Type == html.ElementNode {
		switch n.Data {
		case "td":
			switch tdCount {
			case 0:
				// Kanji num in Jōyō kanji
				// ignore.
			case 1:
				kanji.Rune, err = parseAnchorText(n)
				if err != nil {
					return err
				}
			case 2:
				// old rune is optional, ignore error.
				kanji.OldRune, _ = parseAnchorText(n)
			case 3:
				kanji.Radical, err = parseAnchorText(n)
				if err != nil {
					return err
				}
			case 4:
				// Unknown.
				// TODO(u): ignore?
			case 5:
				kanji.Grade, err = parseNum(n)
				if err != nil {
					return err
				}
			case 6:
				// Year added?
				// TODO(u): ignore?
			case 7:
				kanji.Meaning, err = parseText(n)
				if err != nil {
					return err
				}
			case 8:
				kanji.Pronounciation, err = parseFullText(n)
				if err != nil {
					return err
				}
			}
			tdCount++
		case "tr":
			if kanji != nil {
				kanjis = append(kanjis, kanji)
			}
			kanji = new(Kanji)
			tdCount = 0
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

func parseAnchorText(n *html.Node) (s string, err error) {
	if n.Type == html.ElementNode && n.Data == "a" {
		return parseText(n.FirstChild)
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		s, err = parseAnchorText(c)
		if err == nil {
			return s, nil
		}
	}
	return "", fmt.Errorf("unable to parse anchor text (td=%d) for kanji %d.", tdCount, len(kanjis)+1)
}

func parseText(n *html.Node) (s string, err error) {
	if n.Type == html.TextNode {
		return n.Data, nil
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		s, err = parseText(c)
		if err == nil {
			return s, nil
		}
	}
	return "", fmt.Errorf("unable to parse nested text (td=%d) for kanji %d.", tdCount, len(kanjis)+1)
}

func parseFullText(n *html.Node) (s string, err error) {
	f := func(n *html.Node) {
		switch n.Type {
		case html.TextNode:
			s += n.Data
		case html.ElementNode:
			if n.Data == "br" {
				s += "\n"
			}
		}
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		f(c)
	}
	if len(s) == 0 {
		return "", fmt.Errorf("unable to parse nested text (td=%d) for kanji %d.", tdCount, len(kanjis)+1)
	}
	return s, nil
}

func parseNum(n *html.Node) (num int, err error) {
	if n.Type == html.TextNode {
		num, err = strconv.Atoi(n.Data)
		if err != nil {
			return 0, err
		}
		return num, nil
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		num, err = parseNum(c)
		if err == nil {
			return num, nil
		}
	}
	return 0, fmt.Errorf("unable to parse nested number (td=%d) for kanji %d.", tdCount, len(kanjis)+1)
}

type GradeOrder []*Kanji

func (l GradeOrder) Len() int {
	return len(l)
}

func (l GradeOrder) Swap(i, j int) {
	l[i], l[j] = l[j], l[i]
}

func (l GradeOrder) Less(i, j int) bool {
	a := l[i]
	b := l[j]
	if a.Grade == b.Grade {
		if a.Radical == b.Radical {
			return a.Rune < b.Rune
		}
		return a.Radical < b.Radical
	}
	return a.Grade < b.Grade
}

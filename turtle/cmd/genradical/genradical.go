// As you will notice this is kind of a hack, but it works :)
//
// Parse table of radicals [1]. Use the following steps to prepare the indata
// for this application:
//
//    1. view-source:https://en.wikipedia.org/wiki/Kangxi_radical#Table_of_radicals
//    2. Copy from "<table>" to "</table>" with start below [1].
//    3. Remove the first "<tr>...<tr>"
//    4. Add "<tr></tr>" after the last radical.
//
// [1]: https://en.wikipedia.org/wiki/Kangxi_radical#Table_of_radicals

package main

import (
	"bytes"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"strconv"
	"strings"

	"code.google.com/p/go.net/html"
)

func main() {
	flag.Parse()
	for _, filePath := range flag.Args() {
		err := genradical(filePath)
		if err != nil {
			log.Fatalln(err)
		}
	}
	for _, radical := range radicals {
		fmt.Printf("%#v\n", radical)
	}
}

func genradical(filePath string) (err error) {
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

type Radical struct {
	Num        int
	Rune       string
	AltRune    string
	Strokes    int
	Pinyin     string
	Hiragana   string
	Romaji     string
	Meaning    string
	Frequency  int
	Simplified string
	Examples   string
}

var radicals = make([]*Radical, 0, 214)

var radical *Radical

func f(n *html.Node) (err error) {
	if n.Type == html.ElementNode {
		switch n.Data {
		case "td":
			switch tdCount {
			case 0:
				radical.Num, err = parseNum(n)
				if err != nil {
					return err
				}
			case 1:
				radical.Rune, radical.AltRune, err = parseRune(n)
				if err != nil {
					return err
				}
			case 2:
				radical.Strokes, err = parseStrokes(n)
				if err != nil {
					return err
				}
			case 3:
				radical.Pinyin, err = parsePinyin(n)
				if err != nil {
					return err
				}
			case 4:
				radical.Hiragana, radical.Romaji, err = parseHiragana(n)
				if err != nil {
					return err
				}
			case 5:
				radical.Meaning, err = parseMeaning(n)
				if err != nil {
					return err
				}
			case 6:
				radical.Frequency, err = parseFrequency(n)
				if err != nil {
					return err
				}
			case 7:
				radical.Simplified = parseSimplified(n)
			case 8:
				radical.Examples, err = parseExamples(n)
				if err != nil {
					return err
				}
			}
			tdCount++
		case "tr":
			if radical != nil {
				radicals = append(radicals, radical)
			}
			radical = new(Radical)
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

func parseNum(n *html.Node) (num int, err error) {
	if n.Type == html.TextNode {
		num, err = strconv.Atoi(n.Data)
		if err != nil {
			return 0, err
		}
		log.Println("num:", num)
		return num, nil
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		return parseNum(c)
	}
	return 0, fmt.Errorf("unable to locate num for radical %d.", len(radicals)+1)
}

func parseRune(n *html.Node) (r, altR string, err error) {
	if n.Type == html.TextNode {
		pos := strings.Index(n.Data, " ")
		if pos == -1 {
			r = n.Data
		} else {
			r = n.Data[:pos]
			altR = n.Data[pos+1:]
		}
		log.Println("rune:", r)
		log.Println("alt rune:", altR)
		return r, altR, nil
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		return parseRune(c)
	}
	return "", "", fmt.Errorf("unable to locate rune for radical %d.", len(radicals)+1)
}

func parseStrokes(n *html.Node) (strokes int, err error) {
	if n.Type == html.TextNode {
		strokes, err = strconv.Atoi(n.Data)
		if err != nil {
			return 0, err
		}
		log.Println("strokes:", strokes)
		return strokes, nil
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		return parseStrokes(c)
	}
	return 0, fmt.Errorf("unable to locate strokes for radical %d.", len(radicals)+1)
}

func parsePinyin(n *html.Node) (pinyin string, err error) {
	if n.Type == html.TextNode {
		pinyin = n.Data
		log.Println("pinyin:", pinyin)
		return pinyin, nil
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		return parsePinyin(c)
	}
	return "", fmt.Errorf("unable to locate pinyin for radical %d.", len(radicals)+1)
}

func parseHiragana(n *html.Node) (hiragana, romaji string, err error) {
	if n.Type == html.TextNode {
		a := strings.Split(n.Data, "-")
		if len(a) != 2 {
			return "", "", fmt.Errorf("unable to locate hiragana for radical %d.", len(radicals)+1)
		}
		log.Println("hiragana:", hiragana)
		log.Println("romaji:", romaji)
		return a[0], a[1], nil
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		return parseHiragana(c)
	}
	return "", "", fmt.Errorf("unable to locate hiragana for radical %d.", len(radicals)+1)
}

func parseMeaning(n *html.Node) (meaning string, err error) {
	if n.Type == html.TextNode {
		meaning = n.Data
		log.Println("meaning:", meaning)
		return meaning, nil
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		return parseMeaning(c)
	}
	return "", fmt.Errorf("unable to locate meaning for radical %d.", len(radicals)+1)
}

func parseFrequency(n *html.Node) (frequency int, err error) {
	if n.Type == html.TextNode {
		data := strings.Replace(n.Data, ",", "", -1)
		frequency, err = strconv.Atoi(data)
		if err != nil {
			return 0, err
		}
		log.Println("frequency:", frequency)
		return frequency, nil
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		return parseFrequency(c)
	}
	return 0, fmt.Errorf("unable to locate frequency for radical %d.", len(radicals)+1)
}

func parseSimplified(n *html.Node) (simplified string) {
	if n.Type == html.TextNode {
		simplified = n.Data
		log.Println("simplified:", simplified)
		return simplified
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		return parseSimplified(c)
	}
	return ""
}

func parseExamples(n *html.Node) (examples string, err error) {
	if n.Type == html.TextNode {
		examples = n.Data
		log.Println("examples:", examples)
		return examples, nil
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		return parseExamples(c)
	}
	return "", fmt.Errorf("unable to locate examples for radical %d.", len(radicals)+1)
}

// stash is a tool which extends the stash of TED talks.
package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"os"
	"strings"
	"time"

	"github.com/PuerkitoBio/goquery"
	"github.com/mewmew/playground/archive/ted"
)

// flagAll specifies if all TED talk pages should be crawled (default: only
// crawl first page).
var flagAll bool

func main() {
	flag.BoolVar(&flagAll, "all", false, "List all TED talks (default: only show first page).")
	flag.Parse()
	err := stash("http://www.ted.com/talks/quick-list")
	if err != nil {
		log.Fatalln(err)
	}
	buf, err := json.MarshalIndent(talks, "", "\t")
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println(string(buf))
}

// talks contains the crawled TED talks.
var talks []ted.Talk

// stash downloads a list of all TED talks.
func stash(url string) error {
	// Crawl page.
	if flagAll {
		fmt.Fprintln(os.Stdout, "crawling:", url)
	}
	doc, err := goquery.NewDocument(url)
	if err != nil {
		return err
	}

	// Dump talk information.
	var talk ted.Talk
	dump := func(i int, s *goquery.Selection) {
		switch i % 5 {
		case 0: // Date
			date, err := time.Parse("Jan 2006", s.Text())
			if err != nil {
				log.Println(err)
			}
			talk = ted.Talk{Date: date}
		case 1: // Event
			talk.Event = s.Text()
		case 2: // Title
			talk.Title = s.Find("a").Text()
		case 3: // Duration
			parts := strings.Split(s.Text(), ":")
			duration, err := time.ParseDuration(fmt.Sprintf("%sm%ss", parts[0], parts[1]))
			if err != nil {
				log.Println(err)
			}
			talk.Duration = duration
		case 4: // Download
			high, ok := s.Find("a").Last().Attr("href")
			if ok {
				talk.Download = high
				talks = append(talks, talk)
			} else {
				log.Printf("unable to locate high-definition download link for %q on page %v\n", talk.Title, url)
			}
		}
	}
	doc.Find("td").Each(dump)

	// Crawl next page recursively.
	if flagAll {
		next, ok := doc.Find(".next>a").First().Attr("href")
		if ok {
			return stash("http://www.ted.com" + next)
		}
	}
	return nil
}

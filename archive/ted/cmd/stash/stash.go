// stash is a tool which extends the stash of TED talks.
package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"strings"
	"time"

	"github.com/PuerkitoBio/goquery"
	"github.com/mewmew/playground/archive/ted"
)

func main() {
	// all specifies if all TED talk pages should be crawled (default: only crawl
	// first page).
	var all bool
	flag.BoolVar(&all, "all", false, "list all TED talks (default: show only first page)")
	flag.Parse()
	if err := stash("http://www.ted.com/talks/quick-list", all); err != nil {
		log.Fatalf("%+v", err)
	}
	buf, err := json.MarshalIndent(talks, "", "\t")
	if err != nil {
		log.Fatalf("%+v", err)
	}
	fmt.Println(string(buf))
}

// talks contains the crawled TED talks.
var talks []ted.Talk

// stash downloads a list of all TED talks.
func stash(url string, all bool) error {
	// Crawl page.
	if all {
		log.Println("crawling:", url)
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
			date, err := time.Parse("Jan 2006", strings.TrimSpace(s.Text()))
			if err != nil {
				log.Println(err)
			}
			talk = ted.Talk{Date: date}
		case 1: // Title and Link
			anchor := s.Find("a")
			talk.Title = anchor.Text()
			if link, ok := anchor.Attr("href"); ok {
				if strings.HasPrefix(link, "/") {
					link = "https://www.ted.com" + link
				}
				talk.Link = link
			}
		case 2: // Event
			talk.Event = strings.TrimSpace(s.Text())
		case 3: // Duration
			raw := strings.TrimSpace(s.Text())
			if strings.Contains(raw, ":") {
				parts := strings.Split(raw, ":")
				if len(parts) != 2 {
					log.Printf("invalid duration format `%v` for %q on page %v", raw, talk.Title, url)
				}
				raw = fmt.Sprintf("%sm%ss", parts[0], parts[1])
			} else if strings.Contains(raw, "h ") {
				raw = strings.Replace(raw, " ", "", -1)
			}
			duration, err := time.ParseDuration(raw)
			if err != nil {
				log.Printf("unable to parse duration %v; %v", raw, err)
			}
			talk.Duration = duration
		case 4: // Download
			high, ok := s.Find("a").Last().Attr("href")
			if ok {
				talk.Download = high
			} else {
				log.Printf("unable to locate high-definition download link for %q on page %v", talk.Title, url)
			}
			talks = append(talks, talk)
		}
	}
	doc.Find(".row .quick-list__row > div").Each(dump)

	// Crawl next page recursively.
	if all {
		next, ok := doc.Find("a.pagination__next").First().Attr("href")
		if ok {
			return stash("http://www.ted.com"+next, all)
		}
	}
	return nil
}

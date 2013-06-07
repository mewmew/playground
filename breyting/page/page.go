// Package page implements functions for monitoring pages for changes.
package page

import (
	"crypto/sha512"
	"encoding/hex"
	dbg "fmt"
	"hash"
	"log"
	"time"

	css "code.google.com/p/cascadia"
)

// DefaultTimeout is the default time interval to sleep in between page
// downloads.
const DefaultTimeout = 5 * time.Minute

// Timeout is the time interval to sleep in between page downloads.
var Timeout = DefaultTimeout

// PageKey corresponds to the raw URL and CSS selector of a page. It's primary
// use is in the active page map.
type PageKey struct {
	RawUrl string
	RawSel string
}

// active is a map containing all the active pages.
var active = map[PageKey]bool{}

// Page contains a hash of the content available on the URL that matches the CSS
// selector.
type Page struct {
	PageKey
	sel    css.Selector
	digest hash.Hash
}

// New returns a new page based on the provided URL and CSS selector.
func New(rawUrl, rawSel string) (p *Page, err error) {
	p = new(Page)
	p.RawUrl = rawUrl
	p.RawSel = rawSel
	if rawSel != "" {
		p.sel, err = css.Compile(rawSel)
		if err != nil {
			return nil, err
		}
	}
	p.digest = sha512.New()
	return p, nil
}

// Watch monitors a page for changes. It will continue as long as it is regarded
// as active, and will sleep between each iteration based on Timeout.
func (p *Page) Watch() {
	first := true
	for {
		if !IsActive(p.PageKey) {
			// Stop watching inactive page.
			return
		}
		if !first {
			time.Sleep(Timeout)
		}
		first = false
		err := p.download()
		if err != nil {
			log.Println(err)
			continue
		}
		dbg.Println(hex.Dump(p.digest.Sum(nil)))
		///p.diff()
		/// ### [ todo ] ###
		///   - locate changes.
		/// ### [/ todo ] ###

	}
}

// IsActive returns true if the page is regarded as active.
func IsActive(key PageKey) bool {
	_, ok := active[key]
	return ok
}

func SetActive(key PageKey) {
	active[key] = true
}

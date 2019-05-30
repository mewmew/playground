// Package wallhaven implements search and download functionality for
// wallhaven.cc.
package wallhaven

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"path/filepath"
	"strings"

	"github.com/PuerkitoBio/goquery"
	"github.com/pkg/errors"
)

// Search searches for wallpapers based on the given query and search options.
func Search(query string, options ...Option) (ids []ID, err error) {
	// Parse search options.
	values := make(url.Values)
	if len(query) != 0 {
		values.Add("q", query)
	}
	for _, option := range options {
		key := option.Key()
		val := option.Value()
		values.Add(key, val)
	}

	// Send search request.
	rawquery := values.Encode()
	rawurl := "http://wallhaven.cc/search?" + rawquery
	doc, err := goquery.NewDocument(rawurl)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	// Locate wallpaper IDs in response.
	//
	// Example response:
	//    <a class="preview" href="http://wallhaven.cc/w/oxl3xl">
	f := func(i int, s *goquery.Selection) {
		rawurl, ok := s.Attr("href")
		if !ok {
			return
		}
		pos := strings.Index(rawurl, "/w/")
		if pos == -1 {
			log.Printf("unable to locate wallpaper ID in %q", rawurl)
			return
		}
		id := rawurl[pos+len("/w/"):]
		ids = append(ids, ID(id))
	}
	doc.Find("a.preview").Each(f)
	return ids, nil
}

// ID represents the wallpaper ID of a specific wallpaper on wallhaven.cc.
type ID string

// Download downloads the wallpaper to the given directory, and returns the path
// to the downloaded file.
func (id ID) Download(dir string) (path string, err error) {
	download := func(ext string) (path string, err error) {
		filename := fmt.Sprintf("wallhaven-%s.%s", id, ext)
		path = filepath.Join(dir, filename)
		rawurl := fmt.Sprintf("https://w.wallhaven.cc/full/%s/%s", id[:2], filename)
		resp, err := http.Get(rawurl)
		if err != nil {
			return "", errors.WithStack(err)
		}
		defer resp.Body.Close()
		if resp.StatusCode != http.StatusOK {
			return "", errors.Errorf("invalid status code for %q; expected %d, got %d", rawurl, http.StatusOK, resp.StatusCode)
		}
		buf, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			return "", errors.WithStack(err)
		}
		if err := ioutil.WriteFile(path, buf, 0644); err != nil {
			return "", errors.WithStack(err)
		}
		return path, nil
	}

	// Try to download with jpg extension.
	if path, err := download("jpg"); err == nil {
		// Return early on success.
		return path, nil
	}

	// Fallback to download with png extension.
	path, err = download("png")
	if err != nil {
		return "", errors.WithStack(err)
	}
	return path, nil
}

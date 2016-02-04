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
	"strconv"

	"github.com/PuerkitoBio/goquery"
	"github.com/mewkiz/pkg/errutil"
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
	rawurl := "http://alpha.wallhaven.cc/search?" + rawquery
	doc, err := goquery.NewDocument(rawurl)
	if err != nil {
		return nil, errutil.Err(err)
	}

	// Locate wallpaper IDs in response.
	//
	// Example response:
	//    <figure id="thumb-109603" class="thumb thumb-sfw thumb-general" data-wallpaper-id="109603" style="width:300px;height:200px" >
	f := func(i int, s *goquery.Selection) {
		rawid, ok := s.Attr("data-wallpaper-id")
		if !ok {
			return
		}
		id, err := strconv.Atoi(rawid)
		if err != nil {
			log.Print(errutil.Err(err))
			return
		}
		ids = append(ids, ID(id))
	}
	doc.Find("figure.thumb").Each(f)

	return ids, nil
}

// ID represents the wallpaper ID of a specific wallpaper on wallhaven.cc.
type ID int

// Download downloads the wallpaper to the given directory, and returns the path
// to the downloaded file.
func (id ID) Download(dir string) (path string, err error) {
	download := func(ext string) (path string, err error) {
		filename := fmt.Sprintf("wallhaven-%d.%s", id, ext)
		path = filepath.Join(dir, filename)
		rawurl := "http://wallpapers.wallhaven.cc/wallpapers/full/" + filename
		resp, err := http.Get(rawurl)
		if err != nil {
			return "", errutil.Err(err)
		}
		defer resp.Body.Close()
		if resp.StatusCode != http.StatusOK {
			return "", errutil.Newf("invalid status code; expected %d, got %d", http.StatusOK, resp.StatusCode)
		}
		buf, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			return "", errutil.Err(err)
		}
		if err := ioutil.WriteFile(path, buf, 0644); err != nil {
			return "", errutil.Err(err)
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
		return "", errutil.Err(err)
	}
	return path, nil
}

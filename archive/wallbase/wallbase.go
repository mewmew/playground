// Package wallbase implements search and download functions for wallbase.cc.
package wallbase

import (
	"bytes"
	"errors"
	"fmt"
	"image"
	_ "image/jpeg" // enable jpeg decoding.
	_ "image/png"  // enable png decoding.
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/mewkiz/pkg/bytesutil"
	"github.com/mewkiz/pkg/httputil"
)

// Res represent the screen resolution used in search queries. If blank no
// screen resolution will be enforced.
//
// Example:
//    "1920x1080".
var Res string

// Search searches for wallpapers matching the provided query and returns their
// ids. The search result order is random.
func Search(query string) (ids []int, err error) {
	// Perform search query.
	url := fmt.Sprintf("http://wallbase.cc/search?q=%s&order=random", query)
	if len(Res) > 0 {
		url = fmt.Sprintf("%s&res=%s", url, Res)
	}
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Add("X-Requested-With", "XMLHttpRequest")
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	buf, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	// Locate wallpapers IDs in the response.
	for {
		pos := bytesutil.IndexAfter(buf, []byte(`"http://wallbase.cc/wallpaper/`))
		if pos == -1 {
			break
		}
		end := bytes.IndexByte(buf[pos:], '"')
		if end == -1 {
			return nil, errors.New("wallbase.Search: unmatched quote in wallpaper URL")
		}
		rawID := buf[pos : pos+end]
		buf = buf[pos+end:]
		id, err := strconv.Atoi(string(rawID))
		if err != nil {
			return nil, err
		}
		ids = append(ids, id)
	}

	return ids, nil
}

// Download downloads the wallpaper specified by id and returns it's content and
// file extension.
func Download(id int) (buf []byte, ext string, err error) {
	// Download the wallpaper page.
	url := fmt.Sprintf("http://wallbase.cc/wallpaper/%d", id)
	page, err := httputil.Get(url)
	if err != nil {
		return nil, "", err
	}

	// Locate the wallpaper image URL.
	pos := bytes.Index(page, []byte("http://wallpapers.wallbase.cc/"))
	if pos == -1 {
		return nil, "", fmt.Errorf("wallbase.Download: unable to locate wallpaper image URL for %d", id)
	}
	end := bytes.IndexByte(page[pos:], '"')
	if end == -1 {
		return nil, "", errors.New("wallbase.Download: unmatched quote in wallpaper URL")
	}
	wallURL := page[pos : pos+end]

	// Download the wallpaper image.
	buf, err = httputil.Get(string(wallURL))
	if err != nil {
		return nil, "", err
	}

	// Locate the file extension.
	_, ext, err = image.DecodeConfig(bytes.NewReader(buf))
	if err != nil {
		return nil, "", err
	}

	return buf, ext, nil
}

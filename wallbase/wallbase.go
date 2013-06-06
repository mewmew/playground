// Package wallbase implements search and download functions for wallbase.cc.
package wallbase

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"image"
	_ "image/jpeg"
	_ "image/png"
	"io/ioutil"
	"net/http"
	"strings"

	"github.com/mewkiz/pkg/httputil"
	"github.com/mewkiz/pkg/stringsutil"
)

// Wall is a wallpaper.
type Wall struct {
	// Buf is the image file content of the wallpaper. Use the Download method to
	// retrieve the data.
	Buf []byte
	// Ext is the file extension of the wallpaper.
	Ext string
	// Id is the unique identifier (at wallbase.cc) of the wallpaper.
	Id int `json: "id"`
}

// Search performs a search based on the provided query. The search result order
// is random.
func Search(query string) (walls []*Wall, err error) {
	rawUrl := "http://wallbase.cc/search"
	data := strings.NewReader(fmt.Sprintf("query=%s&orderby=random", query))
	req, err := http.NewRequest("POST", rawUrl, data)
	if err != nil {
		return nil, err
	}
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
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
	err = json.Unmarshal(buf, &walls)
	if err != nil {
		return nil, err
	}
	return walls, nil
}

// Download downloads the wallpaper and stores it's content in wall.Buf.
func (wall *Wall) Download() (err error) {
	if wall.Buf != nil {
		// already downloaded.
		return nil
	}
	imgUrl, err := wall.getImageUrl()
	if err != nil {
		return err
	}
	wall.Buf, err = httputil.Get(imgUrl)
	if err != nil {
		return err
	}
	_, wall.Ext, err = image.DecodeConfig(bytes.NewReader(wall.Buf))
	if err != nil {
		return err
	}
	return nil
}

// getImageUrl locates the image URL of a given wallpaper. The image URL is part
// of a javascript and it is base64 encoded.
func (wall *Wall) getImageUrl() (imgUrl string, err error) {
	rawUrl := fmt.Sprintf("http://wallbase.cc/wallpaper/%d", wall.Id)
	body, err := httputil.GetString(rawUrl)
	if err != nil {
		return "", err
	}
	// example:
	//    document.write('<img ... src="'+B('aHR0cDovL25zMjIzNTA2Lm92aC5uZXQvcm96bmUvYmZjMzIwNzM5ZGY4NzMwOWE2N2E1MTdjMTQ5MDIwODAvd2FsbHBhcGVyLTIzOTMyMTMuanBn')+'" />');
	start := stringsutil.IndexAfter(body, ` src="'+B('`)
	if start == -1 {
		return "", errors.New("wallbase.Download: image URL start position not found.")
	}
	imgUrlEnc := body[start:]
	end := strings.Index(imgUrlEnc, "'")
	if end == -1 {
		return "", errors.New("wallbase.Download: image URL end position not found.")
	}
	imgUrlEnc = imgUrlEnc[:end]
	buf, err := base64.StdEncoding.DecodeString(imgUrlEnc)
	if err != nil {
		return "", err
	}
	return string(buf), nil
}

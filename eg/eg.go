// Package eg implementes functions for communicating with the enigmagroup
// server.
package eg

import "crypto/tls"
import "errors"
import "io/ioutil"
import "net/http"
import "strings"

// Session variables used by eg.
var (
	FieldV4   string
	PhpSessid string
)

// HasSession returns true if an eg session has been set.
func HasSession() bool {
	if len(FieldV4) < 1 {
		return false
	}
	if len(PhpSessid) < 1 {
		return false
	}
	return true
}

// Get performs an HTTP GET on rawUrl, using the eg session.
func Get(rawUrl string) (buf []byte, err error) {
	return request("GET", rawUrl, "")
}

// Get performs an HTTP POST on rawUrl, using the eg session.
func Post(rawUrl string, data string) (buf []byte, err error) {
	return request("POST", rawUrl, data)
}

// request performs an HTTP request, using the eg session.
func request(method string, rawUrl string, data string) (buf []byte, err error) {
	if !HasSession() {
		return nil, errors.New("no session variables set.")
	}
	var req *http.Request
	if len(data) < 1 {
		req, err = http.NewRequest(method, rawUrl, nil)
		if err != nil {
			return nil, err
		}
	} else {
		req, err = http.NewRequest(method, rawUrl, strings.NewReader(data))
		if err != nil {
			return nil, err
		}
		req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
		req.Header.Set("Referer", rawUrl)
	}
	req.AddCookie(&http.Cookie{Name: "enigmafiedV4", Value: FieldV4})
	req.AddCookie(&http.Cookie{Name: "PHPSESSID", Value: PhpSessid})
	r, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer r.Body.Close()
	buf, err = ioutil.ReadAll(r.Body)
	if err != nil {
		return nil, err
	}
	return buf, nil
}

// client enabled the use of insecure SSL connections.
var client = &http.Client{
	Transport: &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	},
}

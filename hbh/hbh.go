package hbh

import "crypto/tls"
import "errors"
import "io/ioutil"
import "net/http"
import "strings"

var PhpSessid string
var FusionUser string

func HasSession() bool {
	if len(PhpSessid) < 1 {
		return false
	}
	if len(FusionUser) < 1 {
		return false
	}
	return true
}

func Get(rawUrl string) (text string, err error) {
	return Request("GET", rawUrl, "")
}

func Post(rawUrl string, data string) (text string, err error) {
	return Request("POST", rawUrl, data)
}

func Request(method string, rawUrl string, data string) (text string, err error) {
	if !HasSession() {
		return "", errors.New("no session variables set.")
	}
	var req *http.Request
	if len(data) < 1 {
		req, err = http.NewRequest(method, rawUrl, nil)
		if err != nil {
			return "", err
		}
	} else {
		req, err = http.NewRequest(method, rawUrl, strings.NewReader(data))
		if err != nil {
			return "", err
		}
		req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	}
	req.AddCookie(&http.Cookie{Name: "PHPSESSID", Value: PhpSessid})
	req.AddCookie(&http.Cookie{Name: "fusion_user", Value: FusionUser})
	r, err := client.Do(req)
	if err != nil {
		return "", err
	}
	defer r.Body.Close()
	buf, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return "", err
	}
	return string(buf), nil
}

var client = &http.Client{
	Transport: &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	},
}

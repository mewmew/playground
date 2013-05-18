package gs

import (
	"bytes"
	"crypto/md5"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"net/http/httputil"
	"strings"

	gouuid "code.google.com/p/go-uuid/uuid"
)

// When Verbose is true verbose output is enabled.
var Verbose bool

// UserAgent is the HTTP User-Agent to be used for all requests.
var UserAgent = "Mozilla/5.0 (X11; Linux x86_64; rv:21.0) Gecko/20100101 Firefox/21.0"

// init initiates the session by retrieving a new session cookie (PHPSESSID).
func (sess *Session) init() (err error) {
	// A HEAD request is enough to get a session cookie.
	req, err := http.NewRequest("HEAD", "http://grooveshark.com/", nil)
	if err != nil {
		return err
	}
	req.Header.Set("User-Agent", UserAgent)
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return err
	}
	resp.Body.Close()

	// Locate the session cookie.
	for _, cookie := range resp.Cookies() {
		if cookie.Name == "PHPSESSID" {
			if Verbose {
				log.Println("Session cookie (PHPSESSID):", cookie.Value)
			}
			sess.cookie = cookie
			return nil
		}
	}

	return errors.New("Session.init: unable to locate the PHPSESSID session cookie.")
}

type gsReqCommToken struct {
	SecretKey string `json:"secretKey"`
}

// Example response (success):
//
//    "result":"5194cf4a8e89f"
//
// Example response (failure):
//
//    "fault":
//    {
//       "code":0,
//       "message":"Error fetching token: did not recognize secret key"
//    }
type gsRespCommToken struct {
	Result string       `json:"result"`
	Err    *gsRespError `json:"fault"`
}

// initCommToken retrieves the communication token for the provided session.
func (sess *Session) initCommToken() (err error) {
	// secretKey is the md5sum of PHPSESSID.
	h := md5.New()
	io.WriteString(h, sess.cookie.Value)
	md5sum := h.Sum(nil)

	// Perform request.
	method := "getCommunicationToken"
	params := gsReqCommToken{
		SecretKey: fmt.Sprintf("%x", md5sum),
	}
	buf, err := sess.request(method, params)
	if err != nil {
		return err
	}

	// Unmarshal JSON response.
	var resp gsRespCommToken
	err = json.Unmarshal(buf, &resp)
	if err != nil {
		return err
	}
	if resp.Err != nil {
		return resp.Err
	}
	sess.commToken = resp.Result

	return nil
}

// methodClient maps request methods to their associated client.
var methodClient = map[string]*client{
	"addSongsToQueue":          jsqueue,
	"getStreamKeyFromSongIDEx": jsqueue,
	"markSongDownloadedEx":     jsqueue,
}

// client contains necessary information for making valid groovshark JSON
// requests.
type client struct {
	// Client name.
	Name string
	// Client revision.
	Rev string
	// Salt is used when generating request tokens.
	Salt string
	// HTTP referer.
	Referer string
}

// defaultClient is the client used for most requests.
var defaultClient = &client{
	Name: "htmlshark",
	Rev:  "20120312",
	Salt: "reallyHotSauce",
}

// jsqueue is the client used for streaming songs.
var jsqueue = &client{
	Name:    "jsqueue",
	Rev:     "20120312.08",
	Salt:    "circlesAndSquares",
	Referer: "http://grooveshark.com/JSQueue.swf?20120312.08",
}

var uuid = gouuid.NewRandom()

type gsReq struct {
	Header gsReqHeader `json:"header"`
	Method string      `json:"method"`
	Params interface{} `json:"parameters"`
}

type gsReqHeader struct {
	Client    string             `json:"client"`
	ClientRev string             `json:"clientRevision"`
	Privacy   int                `json:"privacy"`
	Country   gsReqHeaderCountry `json:"country"`
	Session   string             `json:"session"`
	UUID      string             `json:"uuid"`
}

type gsReqHeaderCountry struct {
	ID  int
	CC1 int
	CC2 int
	CC3 int
	CC4 int
	DMA int
	IPR int
}

type gsRespGeneric struct {
	Result interface{}  `json:"result"`
	Err    *gsRespError `json:"fault"`
}

type gsRespError struct {
	Code int    `json:"code"`
	Msg  string `json:"message"`
}

func (e *gsRespError) Error() string {
	return e.Msg
}

// request performs a JSON request to https://grooveshark.com/more.php.
//
// Example request:
//
//    POST https://grooveshark.com/more.php?getCommunicationToken
//
//    {
//       "header":
//       {
//          "client":"htmlshark",
//          "clientRevision":"20120830",
//          "privacy":0,
//          "country":
//          {
//             "ID":47,
//             "CC1":70368744177664,
//             "CC2":0,
//             "CC3":0,
//             "CC4":0,
//             "DMA":0,
//             "IPR":0
//          },
//          "uuid":"9FF30A69-29DD-4231-8736-5D97BDD4EB4B",
//          "session":"29671432f06faf5a7be5792247e840fc"
//       },
//       "method":"getCommunicationToken",
//       "parameters":
//       {
//          "secretKey":"bf383947cd93a1fbb1c7554d8cfc76db"
//       }
//    }
func (sess *Session) request(method string, params interface{}) (buf []byte, err error) {
	// Create JSON request message.
	client, ok := methodClient[method]
	if !ok {
		client = defaultClient
	}
	data := gsReq{
		Header: gsReqHeader{
			Client:    client.Name,
			ClientRev: client.Rev,
			Country: gsReqHeaderCountry{
				ID:  47,
				CC1: 70368744177664,
			},
			UUID:    strings.ToUpper(uuid.String()),
			Session: sess.cookie.Value,
		},
		Method: method,
		Params: params,
	}
	if sess.commToken != "" {
		// TODO(u): Create token for all requests except getCommunicationToken.
	}
	rawData, err := json.Marshal(data)
	if err != nil {
		return nil, err
	}

	// Send request.
	url := fmt.Sprintf("https://grooveshark.com/more.php?%s", method)
	req, err := http.NewRequest("POST", url, bytes.NewReader(rawData))
	if err != nil {
		return nil, err
	}
	req.Header.Set("User-Agent", UserAgent)
	if client.Referer != "" {
		req.Header.Set("Referer", client.Referer)
	}
	req.AddCookie(sess.cookie)
	if Verbose {
		log.Printf("request (%s):\n", method)
		dump, err := httputil.DumpRequestOut(req, true)
		if err != nil {
			return nil, err
		}
		log.Println(string(dump))
	}
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	// Read response.
	buf, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	if Verbose {
		log.Println("response:")
		log.Println(string(buf))
	}

	return buf, nil
}

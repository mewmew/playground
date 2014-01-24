package gs

import (
	"bytes"
	"crypto/md5"
	"crypto/sha1"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"math/rand"
	"net/http"
	"net/http/httputil"
	"strconv"
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
	req, err := http.NewRequest("GET", "http://grooveshark.com/preload.php?getCommunicationToken=1&hash=", nil)
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

	return errors.New("Session.init: unable to locate the PHPSESSID session cookie")
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

type gsReqUserID struct {
	Name string `json:"name"`
}

// Example response:
//
//    "result":
//    {
//       "type": "user",
//       "data":
//       {
//          "FavoriteSongCount": 2,
//          "LibrarySongCount": 3,
//          "PlaylistCount": 2
//       },
//       "user":
//       {
//          "UserID": 21229114,
//          "FName": "Alice",
//          "LName": "",
//          "Picture": null,
//          "TSAdded": "2013-05-16 10:59:15",
//          "City": null,
//          "State": null,
//          "Country": null,
//          "IsPremium": "0",
//          "Sex": "F",
//          "Flags": "0",
//          "About": null,
//          "Username": "Alice"
//       }
//    }
type gsRespUserID struct {
	Result gsRespUserIDResult `json:"result"`
	Err    *gsRespError       `json:"fault"`
}

type gsRespUserIDResult struct {
	User gsRespUserIDUser `json:"user"`
}

type gsRespUserIDUser struct {
	UserID int
}

//  userID returns the user id associated with the provided username.
func (sess *Session) userID(username string) (userID int, err error) {
	// Perform request.
	// TODO(u): use getUserIDFromUsername method instead.
	method := "getItemByPageName"
	params := gsReqUserID{
		Name: username,
	}
	buf, err := sess.request(method, params)
	if err != nil {
		return 0, err
	}

	// Unmarshal JSON response.
	var resp gsRespUserID
	err = json.Unmarshal(buf, &resp)
	if err != nil {
		return 0, err
	}
	if resp.Err != nil {
		return 0, resp.Err
	}

	return resp.Result.User.UserID, nil
}

type gsReqCollection struct {
	Page   string `json:"page"`
	UserID int    `json:"userID"`
}

type gsRespCollection struct {
	Result gsRespCollectionResult `json:"result"`
	Err    *gsRespError           `json:"fault"`
}

// Example response:
//
//    "result":
//    {
//       "Songs": [
//          {
//             "SongID": "28653841",
//             "Name": "Tiefblau",
//             "AlbumName": "Breathless",
//             "AlbumID": "5564043",
//             "Flags": "0",
//             "ArtistName": "Schiller",
//             "ArtistID": "1700932",
//             "Year": "2010",
//             "CoverArtFilename": "5564043.jpg",
//             "EstimateDuration": null,
//             "IsVerified": "1",
//             "IsLowBitrateAvailable": "0",
//             "Popularity": "1313500001",
//             "TrackNum": "2",
//             "TSAdded": "2013-05-16 11:01:27"
//          }
//       ],
//       "TSModified": 1368719620,
//       "hasMore": false
//    }
type gsRespCollectionResult struct {
	Songs   []*gsSong
	HasMore bool `json:"hasMore"`
}

type gsSong struct {
	Name       string
	ArtistName string
	AlbumName  string
	TrackNum   string
	SongID     string
	ArtistID   string
}

func (gsSong *gsSong) Song() (song *Song, err error) {
	song = &Song{
		Title:  gsSong.Name,
		Artist: gsSong.ArtistName,
		Album:  gsSong.AlbumName,
	}
	if gsSong.TrackNum != "" {
		song.TrackNum, err = strconv.Atoi(gsSong.TrackNum)
		if err != nil {
			return nil, err
		}
	}
	if gsSong.SongID != "" {
		song.id, err = strconv.Atoi(gsSong.SongID)
		if err != nil {
			return nil, err
		}
	}
	if gsSong.ArtistID != "" {
		song.artistID, err = strconv.Atoi(gsSong.ArtistID)
		if err != nil {
			return nil, err
		}
	}
	return song, nil
}

// collection returns a list of the songs in the provided user collection's
// page. hasMore is true if there are more pages available.
func (sess *Session) collection(userID int, page int) (songs []*gsSong, hasMore bool, err error) {
	// Perform request.
	method := "userGetSongsInLibrary"
	params := gsReqCollection{
		Page:   strconv.Itoa(page),
		UserID: userID,
	}
	buf, err := sess.request(method, params)
	if err != nil {
		return nil, false, err
	}

	// Unmarshal JSON response.
	var resp gsRespCollection
	err = json.Unmarshal(buf, &resp)
	if err != nil {
		return nil, false, err
	}
	if resp.Err != nil {
		return nil, false, resp.Err
	}

	return resp.Result.Songs, resp.Result.HasMore, nil
}

type gsReqFavorites struct {
	OfWhat string `json:"ofWhat"`
	UserID int    `json:"userID"`
}

// Example response:
//
//    "result": [
//       {
//          "Name": "Tiefblau",
//          "SongID": "28653841",
//          "Flags": "0",
//          "ArtistID": "1700932",
//          "ArtistName": "Schiller",
//          "AlbumID": "5564043",
//          "AlbumName": "Breathless",
//          "CoverArtFilename": "5564043.jpg",
//          "TSFavorited": "2013-05-16 11:01:33",
//          "IsLowBitrateAvailable": "0",
//          "EstimateDuration": null,
//          "IsVerified": "1",
//          "Popularity": "1313500001",
//          "TrackNum": "2"
//       }
//    ]
type gsRespFavorites struct {
	Result []*gsSong    `json:"result"`
	Err    *gsRespError `json:"fault"`
}

// favorites returns a list of the provided user's favorite songs.
func (sess *Session) favorites(userID int) (songs []*gsSong, err error) {
	// Perform request.
	method := "getFavorites"
	params := gsReqFavorites{
		OfWhat: "Songs",
		UserID: userID,
	}
	buf, err := sess.request(method, params)
	if err != nil {
		return nil, err
	}

	// Unmarshal JSON response.
	var resp gsRespFavorites
	err = json.Unmarshal(buf, &resp)
	if err != nil {
		return nil, err
	}
	if resp.Err != nil {
		return nil, resp.Err
	}

	return resp.Result, nil
}

type gsReqPlaylists struct {
	UserID int `json:"userID"`
}

// Example response:
//
//    "result":
//    {
//       "Playlists": [
//          {
//             "UUID": "5194f4e3e56c67243a000001",
//             "TSAdded": "2013-05-16 11:01:55",
//             "Picture": "1699951-2579234-4296074-5564043.jpg",
//             "TSModified": 1368716515,
//             "Name": "blip",
//             "PlaylistID": 86557331,
//             "SongCount": 4,
//             "UserID": 21229114
//          }
//       ]
//    }
type gsRespPlaylists struct {
	Result gsRespPlaylistsResult `json:"result"`
	Err    *gsRespError          `json:"fault"`
}

type gsRespPlaylistsResult struct {
	Playlists []*gsPlaylist
}

type gsPlaylist struct {
	Name       string
	PlaylistID int
}

// playlists  returns a list of the provided user's playlists.
func (sess *Session) playlists(userID int) (playlists []*gsPlaylist, err error) {
	// Perform request.
	method := "userGetPlaylists"
	params := gsReqPlaylists{
		UserID: userID,
	}
	buf, err := sess.request(method, params)
	if err != nil {
		return nil, err
	}

	// Unmarshal JSON response.
	var resp gsRespPlaylists
	err = json.Unmarshal(buf, &resp)
	if err != nil {
		return nil, err
	}
	if resp.Err != nil {
		return nil, resp.Err
	}

	return resp.Result.Playlists, nil
}

type gsReqPlaylistSongs struct {
	PlaylistID int `json:"playlistID"`
}

// Example response:
//
//    "result":
//    {
//       "Songs": [
//          {
//             "SongID": "28653841",
//             "Name": "Tiefblau",
//             "SongNameID": "12624501",
//             "AlbumID": "5564043",
//             "AlbumName": "Breathless",
//             "ArtistID": "1700932",
//             "ArtistName": "Schiller",
//             "AvgRating": null,
//             "IsVerified": "1",
//             "CoverArtFilename": "5564043.jpg",
//             "Year": "2010",
//             "UserRating": "0",
//             "EstimateDuration": null,
//             "Popularity": "1313500001",
//             "TrackNum": "2",
//             "IsLowBitrateAvailable": "0",
//             "Flags": "0",
//             "Sort": 0
//          }
//       ]
//    }
type gsRespPlaylistSongs struct {
	Result gsRespPlaylistSongsResult `json:"result"`
	Err    *gsRespError              `json:"fault"`
}

type gsRespPlaylistSongsResult struct {
	Songs []*gsSong
}

// playlistSongs  returns a list of all songs in the provided playlist.
func (sess *Session) playlistSongs(playlistID int) (songs []*gsSong, err error) {
	// Perform request.
	method := "playlistGetSongs"
	params := gsReqPlaylistSongs{
		PlaylistID: playlistID,
	}
	buf, err := sess.request(method, params)
	if err != nil {
		return nil, err
	}

	// Unmarshal JSON response.
	var resp gsRespPlaylistSongs
	err = json.Unmarshal(buf, &resp)
	if err != nil {
		return nil, err
	}
	if resp.Err != nil {
		return nil, resp.Err
	}

	return resp.Result.Songs, nil
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
	// Use get_keys.php from https://github.com/matt-/GrooveShark_PHP to update
	// Rev and Salt.
	Rev:  "20120830",
	Salt: "greenPlants",
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
	Token     string             `json:"token,omitempty"`
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
		// Generate request token.
		data.Header.Token = sess.genToken(method)
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

// genToken generates and returns a request token based on the communication
// token.
func (sess *Session) genToken(method string) (token string) {
	// Get a value between 0x100 and 0x1000000
	n := rand.Intn(0x1000000 - 0x100)
	n += 0x100
	rnd := fmt.Sprintf("%06x", n)

	// Create the request token.
	h := sha1.New()
	client, ok := methodClient[method]
	if !ok {
		client = defaultClient
	}
	plain := fmt.Sprintf("%s:%s:%s:%s", method, sess.commToken, client.Salt, rnd)
	io.WriteString(h, plain)
	token = fmt.Sprintf("%s%x", rnd, h.Sum(nil))
	if Verbose {
		log.Println("request token (plain):", plain)
		log.Println("request token:", token)
	}

	return token
}

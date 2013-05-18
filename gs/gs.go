// Package gs implements an unofficial API for grooveshark.com.
package gs

import (
	"errors"
	"net/http"
	"strconv"
)

// Session contains the cookies and token of a grooveshark session.
type Session struct {
	// Session cookie (PHPSESSID).
	cookie *http.Cookie
	// Communication token based on the user's session.
	commToken string
}

// NewSession creates an unauthenticated session.
func NewSession() (sess *Session, err error) {
	sess = new(Session)

	// Get session cookie (PHPSESSID).
	err = sess.init()
	if err != nil {
		return nil, err
	}

	// Get communication token based on the session cookie.
	err = sess.initCommToken()
	if err != nil {
		return nil, err
	}

	return sess, nil
}

// UserId returns the user id associated with the provided username.
func (sess *Session) UserId(username string) (userId int, err error) {
	return 0, errors.New("gs.UserId: not yet implemented.")
}

// UserSongs returns a list of all songs in the provided user's collection.
func (sess *Session) UserSongs(userId int) (songs []*Song, err error) {
	// Get one page at the time.
	for page := 0; ; page++ {
		gsSongs, hasMore, err := sess.collection(userId, page)
		if err != nil {
			return nil, err
		}
		for _, gsSong := range gsSongs {
			song := &Song{
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
				song.artistId, err = strconv.Atoi(gsSong.ArtistID)
				if err != nil {
					return nil, err
				}
			}
			songs = append(songs, song)
		}
		if !hasMore {
			break
		}
	}

	return songs, nil
}

// UserFavorites returns a list of the provided user's favorite songs.
func (sess *Session) UserFavorites(userId int) (songs []*Song, err error) {
	return nil, errors.New("gs.UserFavorites: not yet implemented.")
}

// UserPlaylists returns a list of the provided user's playlists.
func (sess *Session) UserPlaylists(userId int) (playlists []*Playlist, err error) {
	return nil, errors.New("gs.UserPlaylists: not yet implemented.")
}

// A Song is a music track with associated information.
type Song struct {
	// Song title.
	Title string
	// Artist of song.
	Artist string
	// Song album name.
	Album string
	// Album track number.
	TrackNum int
	// Song id.
	id int
	// Artist id.
	artistId int
}

// A Playlist is an ordered list of songs with an associated name.
type Playlist struct {
	// Playlist name.
	Name string
	// An ordered slice of songs.
	Songs []*Song
}

// Package gs implements an unofficial API for grooveshark.com.
package gs

import (
	"fmt"
	"net/http"
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

// UserID returns the user id associated with the provided username.
func (sess *Session) UserID(username string) (userID int, err error) {
	userID, err = sess.userID(username)
	if err != nil {
		return 0, err
	}
	if userID == 0 {
		return 0, fmt.Errorf("Session.UserID: unable to locate user id for %q", username)
	}
	return userID, nil
}

// UserSongs returns a list of all songs in the provided user's collection.
func (sess *Session) UserSongs(userID int) (songs []*Song, err error) {
	// Get one page at the time.
	for page := 0; ; page++ {
		gsSongs, hasMore, err := sess.collection(userID, page)
		if err != nil {
			return nil, err
		}
		for _, gsSong := range gsSongs {
			song, err := gsSong.Song()
			if err != nil {
				return nil, err
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
func (sess *Session) UserFavorites(userID int) (songs []*Song, err error) {
	gsSongs, err := sess.favorites(userID)
	if err != nil {
		return nil, err
	}
	for _, gsSong := range gsSongs {
		song, err := gsSong.Song()
		if err != nil {
			return nil, err
		}
		songs = append(songs, song)
	}

	return songs, nil
}

// UserPlaylists returns a list of the provided user's playlists.
func (sess *Session) UserPlaylists(userID int) (playlists []*Playlist, err error) {
	// Locate user playlists.
	gsPlaylists, err := sess.playlists(userID)
	if err != nil {
		return nil, err
	}
	for _, gsPlaylist := range gsPlaylists {
		playlist := &Playlist{
			Name: gsPlaylist.Name,
			id:   gsPlaylist.PlaylistID,
		}

		// Populate the playlist with it's associated songs.
		gsSongs, err := sess.playlistSongs(playlist.id)
		if err != nil {
			return nil, err
		}
		for _, gsSong := range gsSongs {
			song, err := gsSong.Song()
			if err != nil {
				return nil, err
			}
			playlist.Songs = append(playlist.Songs, song)
		}
		playlists = append(playlists, playlist)
	}

	return playlists, nil
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
	artistID int
}

func (song *Song) String() string {
	return fmt.Sprintf("%s - %s", song.Artist, song.Title)
}

// A Playlist is an ordered list of songs with an associated name.
type Playlist struct {
	// Playlist name.
	Name string
	// An ordered slice of songs.
	Songs []*Song
	// Playlist id.
	id int
}

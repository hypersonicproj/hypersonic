package library

import "time"

type Playlist struct {
	Name      string
	CreatedAt time.Time
}

type Id interface {
	Text() string
}

type AlbumId Id

type AlbumArtist struct {
	Name string
}

type TrackId Id

type Artist struct {
	Name string
}

package library

import (
	"fmt"
	"time"
)

func LoadAlbum(publisherName string, title string, releasedAt *time.Time, addedAt time.Time) Album {
	return Album{
		Title: title,
		AlbumArtist: AlbumArtist{
			Name: publisherName,
		},
		ReleasedAt: releasedAt,
		AddedAt:    addedAt,
	}
}

type Album struct {
	Title       string
	AlbumArtist AlbumArtist
	ReleasedAt  *time.Time
	AddedAt     time.Time
}

func (a Album) Id() AlbumId {
	return AlbumId(&id{
		text: fmt.Sprintf("%s-%s", a.AlbumArtist.Name, a.Title),
	})
}

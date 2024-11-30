package domain

import (
	"fmt"
	"time"
)

func LoadAlbum(publisherName string, title string, releasedAt *time.Time, addedAt time.Time) Album {
	return &album{
		Title: title,
		AlbumArtist: AlbumArtist{
			Name: publisherName,
		},
		ReleasedAt: releasedAt,
		AddedAt:    addedAt,
	}
}

type album struct {
	Title       string
	AlbumArtist AlbumArtist
	ReleasedAt  *time.Time
	AddedAt     time.Time
}

func (a *album) Id() AlbumId {
	return AlbumId(&id{
		text: fmt.Sprintf("%s-%s", a.AlbumArtist.Name, a.Title),
	})
}

func (a *album) Get() album {
	return *a
}

package domain

import (
	"fmt"
	"time"
)

func LoadAlbum(publisherName string, title string, addedAt time.Time) Album {
	return &album{
		Title: title,
		Publisher: Publisher{
			Name: publisherName,
		},
		AddedAt: addedAt,
	}
}

type album struct {
	Title     string
	Publisher Publisher
	AddedAt   time.Time
}

func (a *album) Id() AlbumId {
	return AlbumId(&id{
		text: fmt.Sprintf("%s-%s", a.Publisher.Name, a.Title),
	})
}

func (a *album) Get() album {
	return *a
}

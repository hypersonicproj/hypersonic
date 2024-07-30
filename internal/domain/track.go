package domain

import "fmt"

func LoadTrack(albumId AlbumId, title, artist string, trackNumber uint) Track {
	return &track{
		Title:            title,
		Artist:           Artist{Name: artist},
		AlbumTrackNumber: trackNumber,
		albumId:          albumId,
	}
}

type track struct {
	Title            string
	Artist           Artist
	AlbumTrackNumber uint
	albumId          AlbumId
}

func (t *track) Id() TrackId {
	return TrackId(&id{
		text: fmt.Sprintf("%s-%s", t.albumId.Text(), t.Title),
	})
}

func (a *track) Get() track {
	return *a
}

package library

import (
	"fmt"
	"time"
)

func LoadTrack(albumId Id, albumArtist AlbumArtist, title, artist string, genre *string, trackNumber *int, releasedAt *time.Time, addedAt time.Time) Track {
	return Track{
		Title:            title,
		Artist:           Artist{Name: artist},
		Genre:            genre,
		AlbumTrackNumber: trackNumber,
		AlbumId:          albumId,
		AlbumArtist:      albumArtist,
		ReleasedAt:       releasedAt,
		AddedAt:          addedAt,
	}
}

type Track struct {
	Title            string
	Artist           Artist
	Genre            *string
	AlbumTrackNumber *int
	AlbumId          AlbumId
	AlbumArtist      AlbumArtist
	ReleasedAt       *time.Time
	AddedAt          time.Time
}

func (t Track) Id() TrackId {
	return TrackId(&id{
		text: fmt.Sprintf("%s-%s", t.AlbumId.Text(), t.Title),
	})
}

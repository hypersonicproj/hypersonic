package id3v2

import (
	"fmt"
	"io"

	"github.com/bogem/id3v2"
)

type Tag struct {
	Title       string
	Artist      string
	Album       string
	AlbumArtist string
	Genre       string
	Year        string
	Track       string
}

func Read(r io.Reader) (Tag, error) {
	tag, err := id3v2.ParseReader(r, id3v2.Options{Parse: true})
	if err != nil {
		return Tag{}, fmt.Errorf("failed to open id3 tag: %w", err)
	}
	defer tag.Close()

	return Tag{
		Title:       tag.Title(),
		Artist:      tag.Artist(),
		Album:       tag.Album(),
		AlbumArtist: tag.GetTextFrame("TPE2").Text,
		Genre:       tag.Genre(),
		Year:        tag.Year(),
		Track:       tag.GetTextFrame("TRCK").Text,
	}, nil
}

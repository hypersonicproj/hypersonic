package id3v2

import (
	"fmt"
	audio_tag "hypersonic/internal/pkg/audio/tag"
	"io"

	"github.com/bogem/id3v2"
)

func Read(r io.Reader) (audio_tag.Tag, error) {
	id3v2Tag, err := id3v2.ParseReader(r, id3v2.Options{Parse: true})
	if err != nil {
		return audio_tag.Tag{}, fmt.Errorf("failed to open id3 tag: %w", err)
	}
	defer id3v2Tag.Close()

	return audio_tag.Tag{
		Title:       id3v2Tag.Title(),
		Artist:      id3v2Tag.Artist(),
		Album:       id3v2Tag.Album(),
		AlbumArtist: id3v2Tag.GetTextFrame("TPE2").Text,
		Genre:       id3v2Tag.Genre(),
		Year:        id3v2Tag.Year(),
		Release:     id3v2Tag.GetTextFrame("TDRL").Text,
		Track:       id3v2Tag.GetTextFrame("TRCK").Text,
	}, nil
}

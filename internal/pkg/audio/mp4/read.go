package mp4

import (
	"fmt"
	audio_tag "hypersonic/internal/pkg/audio/tag"
	"io/fs"
	"strconv"

	"github.com/tingtt/mp4tag"
)

func Read(f fs.File) (audio_tag.Tag, error) {
	reader, err := mp4tag.Reader(f)
	if err != nil {
		return audio_tag.Tag{}, fmt.Errorf("failed to open mp4 tag: %w", err)
	}
	mp4Tag, err := reader.Read()
	if err != nil {
		return audio_tag.Tag{}, fmt.Errorf("failed to open mp4 tag: %w", err)
	}

	return audio_tag.Tag{
		Title:       mp4Tag.Title,
		Artist:      mp4Tag.Artist,
		Album:       mp4Tag.Album,
		AlbumArtist: mp4Tag.AlbumArtist,
		Genre:       mp4tag.ResolveGenreName(mp4Tag.Genre),
		Year:        strconv.Itoa(int(mp4Tag.Year)),
		Track:       strconv.Itoa(int(mp4Tag.TrackNumber)),
	}, nil
}

package mp4

import (
	"fmt"
	audio_tag "hypersonic/internal/pkg/audio/tag"
	"io/fs"

	"github.com/tingtt/qtffilst"
)

func Read(f fs.File) (audio_tag.Tag, error) {
	reader, err := qtffilst.NewReader(f)
	if err != nil {
		return audio_tag.Tag{}, fmt.Errorf("failed to open mp4 tag: %w", err)
	}
	mp4Tag, err := reader.Read()
	if err != nil {
		return audio_tag.Tag{}, fmt.Errorf("failed to open mp4 tag: %w", err)
	}

	tag := audio_tag.Tag{}
	if mp4Tag.TitleC != nil {
		tag.Title = mp4Tag.TitleC.Text
	}
	if mp4Tag.Artist != nil {
		tag.Artist = mp4Tag.Artist.Text
	}
	if mp4Tag.AlbumC != nil {
		tag.Album = mp4Tag.AlbumC.Text
	}
	if mp4Tag.AlbumArtist != nil {
		tag.AlbumArtist = mp4Tag.AlbumArtist.Text
	}
	if mp4Tag.GenreC != nil {
		tag.Genre = mp4Tag.GenreC.Text
	}
	if mp4Tag.ContentCreateDate != nil {
		tag.Year = mp4Tag.ContentCreateDate.Text
	}
	if mp4Tag.ReleaseDate != nil {
		tag.Release = mp4Tag.ReleaseDate.Text
	}
	if mp4Tag.TrackNumber != nil {
		tag.Track = fmt.Sprintf("%d/%d", mp4Tag.TrackNumber.Number, mp4Tag.TrackNumber.Total)
	}

	return tag, nil
}

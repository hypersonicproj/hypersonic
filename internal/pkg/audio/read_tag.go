package audio

import (
	"errors"
	"hypersonic/internal/pkg/audio/id3v2"
	"hypersonic/internal/pkg/audio/mp4"
	audio_tag "hypersonic/internal/pkg/audio/tag"
	"io/fs"
	"path/filepath"
)

var (
	ErrUnsupportedFileFormat = errors.New("unsupported file format")
)

func ReadTag(file fs.File) (audio_tag.Tag, error) {
	stat, err := file.Stat()
	if err != nil {
		return audio_tag.Tag{}, err
	}

	switch filepath.Ext(stat.Name()) {
	case ".mp3":
		return id3v2.Read(file)
	case ".m4a":
		return mp4.Read(file)
	default:
		return audio_tag.Tag{}, ErrUnsupportedFileFormat
	}
}

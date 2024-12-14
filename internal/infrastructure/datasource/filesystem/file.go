package filesystem

import (
	"fmt"
	"hypersonic/internal/domain/library"
	"hypersonic/internal/pkg/audio"
	audio_tag "hypersonic/internal/pkg/audio/tag"
	"hypersonic/internal/usecase/search"
	"io/fs"
	"strconv"
	"strings"
	"time"
)

func loadFile(fsys fs.FS, trackPath string, entry fs.DirEntry) (File, error) {
	file, err := fsys.Open(trackPath)
	if err != nil {
		return File{}, fmt.Errorf("failed to open audio file (%s): %w", trackPath, err)
	}
	defer file.Close()

	tag, err := audio.ReadTag(file)
	if err != nil {
		return File{}, err
	}

	fileInfo, err := entry.Info()
	if err != nil {
		return File{}, fmt.Errorf("failed to read file info (%s): %w", trackPath, err)
	}

	return File{tag, fileInfo.ModTime()}, nil
}

type File struct {
	tag     audio_tag.Tag
	modTime time.Time
}

func (f *File) Album() library.Album {
	var releasedAt *time.Time
	if len(f.tag.Release) >= 10 {
		if _releasedAt, err := time.Parse(time.DateOnly, f.tag.Release[:10]); err == nil {
			releasedAt = &_releasedAt
		}
	}
	return library.LoadAlbum(f.tag.AlbumArtist, f.tag.Album, releasedAt, f.modTime)
}

func (f *File) Track() search.Track {
	album := f.Album()
	var genre *string
	if f.tag.Genre != "" {
		genre = &f.tag.Genre
	}
	var trackNumber *int
	if f.tag.Track != "" {
		i, err := strconv.Atoi(strings.Split(f.tag.Track, "/")[0])
		if err == nil {
			trackNumber = &i
		}
	}
	var releasedAt *time.Time
	if len(f.tag.Release) >= 10 {
		if _releasedAt, err := time.Parse(time.DateOnly, f.tag.Release[:10]); err == nil {
			releasedAt = &_releasedAt
		}
	}
	return search.Track{
		Number: 0,
		Track: library.LoadTrack(
			album.Id(), album.AlbumArtist,
			f.tag.Title, f.tag.Artist, genre, trackNumber,
			releasedAt, f.modTime,
		),
	}
}

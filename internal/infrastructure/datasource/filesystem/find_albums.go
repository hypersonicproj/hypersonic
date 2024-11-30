package filesystem

import (
	"errors"
	"fmt"
	"hypersonic/internal/domain"
	"hypersonic/internal/pkg/audio"
	audio_tag "hypersonic/internal/pkg/audio/tag"
	"hypersonic/internal/pkg/tree"
	"hypersonic/internal/usecase/search"
	"io/fs"
	"strings"
	"time"
)

func (s *filesystem) FindAlbumsReleaseDateDesc(option search.FindOption) ([]domain.Album, error) {
	var root *tree.Node[domain.Album]
	err := walkAlbumDir(s.read, func(album domain.Album) {
		albumReleaseDateAddedAtDescending := func(new, curr domain.Album) (isLeft bool) {
			_new := new.Get()
			_curr := curr.Get()
			if _new.ReleasedAt == nil {
				return _curr.ReleasedAt == nil
			} else if _curr.ReleasedAt == nil {
				return _new.ReleasedAt != nil
			}
			if /* same release date */
			_new.ReleasedAt.Format(time.DateOnly) == _curr.ReleasedAt.Format(time.DateOnly) {
				return _new.AddedAt.After(_curr.AddedAt)
			}
			return _new.ReleasedAt.After(*_curr.ReleasedAt)
		}
		root = tree.Insert(root, album, albumReleaseDateAddedAtDescending)
	})
	if err != nil {
		return nil, err
	}

	var albumList []domain.Album
	tree.InOrderTraversal(root, &albumList)
	return albumList, nil
}

func (s *filesystem) FindAlbumsNameAsc(option search.FindOption) ([]domain.Album, error) {
	var root *tree.Node[domain.Album]
	err := walkAlbumDir(s.read, func(album domain.Album) {
		albumNameAscending := func(new, curr domain.Album) (isLeft bool) {
			return new.Get().Title < curr.Get().Title
		}
		root = tree.Insert(root, album, albumNameAscending)
	})
	if err != nil {
		return nil, err
	}

	var albumList []domain.Album
	tree.InOrderTraversal(root, &albumList)
	return albumList, nil
}

var (
	ErrInvalidDateFormat = errors.New("invalid date format")
)

func walkAlbumDir(fsys fs.FS, yield func(domain.Album)) error {
	return fs.WalkDir(fsys, ".", func(path string, entry fs.DirEntry, err error) error {
		if err != nil {
			return fmt.Errorf("failed to filepath.WalkDir: %w", err)
		}
		if /* not album directory */ !entry.IsDir() || !strings.Contains(path, "/") {
			return nil // skip
		}
		albumDir := entry

		tag, err := inspectAlbumInDir(fsys, path)
		if errors.Is(err, errDirNotContainsAnyAudioFiles) {
			return nil // skip
		}
		if err != nil {
			return err
		}

		var releasedAtP *time.Time
		if releasedAt, err := time.Parse(time.DateOnly, tag.Release[:10]); err == nil {
			releasedAtP = &releasedAt
		}

		albumDirInfo, err := albumDir.Info()
		if err != nil {
			return fmt.Errorf("failed to fs.DirEntry.Info: %w", err)
		}

		album := domain.LoadAlbum(tag.AlbumArtist, tag.Album, releasedAtP, albumDirInfo.ModTime())
		yield(album)
		return nil
	})
}

var (
	errDirNotContainsAnyAudioFiles = errors.New("album dir not contains any audio files")
)

func inspectAlbumInDir(fsys fs.FS, albumDirPath string) (tag audio_tag.Tag, err error) {
	err = fs.WalkDir(fsys, albumDirPath, func(trackPath string, entry fs.DirEntry, err error) error {
		if err != nil {
			return fmt.Errorf("failed to filepath.WalkDir: %w", err)
		}
		if entry.IsDir() {
			return nil
		}

		file, err := fsys.Open(trackPath)
		if err != nil {
			return fmt.Errorf("failed to open audio file (%s): %w", trackPath, err)
		}
		defer file.Close()

		parsedTag, err := audio.ReadTag(file)
		if errors.Is(err, audio.ErrUnsupportedFileFormat) {
			return nil // Skip unsupported file
		}
		if err != nil {
			return err
		}
		tag = parsedTag
		return fs.SkipDir // Found the first audio file, no need to continue walking
	})
	if err != nil {
		return audio_tag.Tag{}, fmt.Errorf("failed to find audio file: %w", err)
	}
	if /* album dir not contains any audio files */ tag.Title == "" {
		return audio_tag.Tag{}, errDirNotContainsAnyAudioFiles
	}

	return tag, nil
}

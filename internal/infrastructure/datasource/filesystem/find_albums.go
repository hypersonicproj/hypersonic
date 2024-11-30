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
)

func (s *filesystem) FindAlbumsAddedDesc(option search.FindOption) ([]domain.Album, error) {
	var root *tree.Node[domain.Album]
	err := walkAlbumDir(s.read, func(album domain.Album) {
		albumAddedDescending := func(new, curr domain.Album) (isLeft bool) {
			return new.Get().AddedAt.After(curr.Get().AddedAt)
		}
		root = tree.Insert(root, album, albumAddedDescending)
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

func walkAlbumDir(fsys fs.FS, yield func(domain.Album)) error {
	return fs.WalkDir(fsys, ".", func(path string, entry fs.DirEntry, err error) error {
		if err != nil {
			return fmt.Errorf("failed to filepath.WalkDir: %w", err)
		}
		if !entry.IsDir() {
			return nil // skip
		}
		if !strings.Contains(path, "/") {
			return nil // skip
		}

		tag, err := inspectAlbumInDir(fsys, path)
		if errors.Is(err, errDirNotContainsAnyAudioFiles) {
			return nil // skip
		}
		if err != nil {
			return err
		}

		info, err := entry.Info()
		if err != nil {
			return fmt.Errorf("failed to fs.DirEntry.Info: %w", err)
		}

		album := domain.LoadAlbum(tag.AlbumArtist, tag.Album, info.ModTime())
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
		if errors.Is(err, audio.ErrUnupportedFileFormat) {
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

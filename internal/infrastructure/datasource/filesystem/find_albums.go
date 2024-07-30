package filesystem

import (
	"fmt"
	"hypersonic/internal/domain"
	"hypersonic/internal/pkg/tree"
	"hypersonic/internal/usecase/search"
	"io/fs"
	"path/filepath"
	"strings"

	"github.com/bogem/id3v2"
)

func albumAddedDescending(new, curr domain.Album) (isLeft bool) {
	return new.Get().AddedAt.After(curr.Get().AddedAt)
}

func (s *filesystem) FindAlbumsAddedDesc(option search.FindOption) ([]domain.Album, error) {
	var root *tree.Node[domain.Album]

	err := fs.WalkDir(s.instance, ".", func(path string, entry fs.DirEntry, err error) error {
		if err != nil {
			return fmt.Errorf("failed to filepath.WalkDir: %w", err)
		}
		if !entry.IsDir() {
			return nil // skip
		}
		if !strings.Contains(path, "/") {
			return nil // skip
		}

		albumTitle, err := inspectAlbumTitleInDir(s.instance, path)
		if err != nil {
			return err
		}
		if /* album dir not contains any audio files */ albumTitle == "" {
			return nil // skip
		}

		info, err := entry.Info()
		if err != nil {
			return fmt.Errorf("failed to fs.DirEntry.Info: %w", err)
		}

		album := domain.LoadAlbum( /* publisherName */ filepath.Base(path), albumTitle, info.ModTime())
		root = tree.Insert(root, album, albumAddedDescending)
		return nil
	})
	if err != nil {
		return nil, err
	}

	var albumList []domain.Album
	tree.InOrderTraversal(root, &albumList)
	return albumList, nil
}

func (s *filesystem) FindAlbumsNameAsc(option search.FindOption) ([]domain.Album, error) {
	panic("unimplemented")
}

func inspectAlbumTitleInDir(fsys fs.FS, albumDirPath string) (string, error) {
	var albumTitle string

	err := fs.WalkDir(fsys, albumDirPath, func(trackPath string, entry fs.DirEntry, err error) error {
		if err != nil {
			return fmt.Errorf("failed to filepath.WalkDir: %w", err)
		}
		if entry.IsDir() {
			return nil
		}

		if isSupportedAudioFile(trackPath) {
			file, err := fsys.Open(trackPath)
			if err != nil {
				return fmt.Errorf("failed to open audio file (%s): %w", trackPath, err)
			}
			tag, err := id3v2.ParseReader(file, id3v2.Options{Parse: true})
			if err != nil {
				return fmt.Errorf("failed to open id3 tag: %w", err)
			}
			defer tag.Close()

			albumTitle = tag.Album()

			return fs.SkipDir // Found the first audio file, no need to continue walking
		}
		return nil
	})
	if err != nil {
		return "", fmt.Errorf("failed to find audio file: %w", err)
	}

	return albumTitle, nil
}

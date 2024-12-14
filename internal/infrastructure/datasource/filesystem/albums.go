package filesystem

import (
	"errors"
	"fmt"
	"hypersonic/internal/usecase/search"
	"io/fs"
	"iter"
	"strings"
)

func newAlbumsRepository(fs fs.FS) search.AlbumsRepository {
	return &albums{fs}
}

type albums struct {
	fs fs.FS
}

// All implements search.AlbumsRepository.
func (a *albums) All() iter.Seq2[search.AlbumRepositoryLoader, error] {
	return func(yield func(search.AlbumRepositoryLoader, error) bool) {
		err := fs.WalkDir(a.fs, ".", func(path string, entry fs.DirEntry, err error) error {
			if err != nil {
				return fmt.Errorf("failed to fs.WalkDir: %w", err)
			}
			if /* not album directory */ !entry.IsDir() || !strings.Contains(path, "/") {
				return nil // skip
			}

			albumLoader, err := newAlbumLoader(a.fs, path)
			if err != nil {
				return err
			}

			if !yield(albumLoader, nil) {
				return fs.SkipAll // break
			}
			return nil
		})
		if err != nil && !errors.Is(err, fs.SkipAll) {
			yield(nil, err)
		}
	}
}

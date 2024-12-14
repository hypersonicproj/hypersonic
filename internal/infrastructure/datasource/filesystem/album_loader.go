package filesystem

import (
	"errors"
	"hypersonic/internal/domain/library"
	"hypersonic/internal/pkg/audio"
	"hypersonic/internal/usecase/search"
	"io/fs"
	"iter"
)

func newAlbumLoader(fsys fs.FS, albumDirPath string) (*albumLoader, error) {
	next, stop := iter.Pull2(func(yield func(File, error) bool) {
		err := fs.WalkDir(fsys, albumDirPath, func(trackPath string, entry fs.DirEntry, err error) error {
			if entry.IsDir() {
				return nil // continue (skip unsupported file)
			}

			file, err := loadFile(fsys, trackPath, entry)
			if errors.Is(err, audio.ErrUnsupportedFileFormat) {
				return nil // continue (skip unsupported file)
			}
			if err != nil {
				return err // break
			}

			_continue := yield(file, nil)
			if !_continue {
				return fs.SkipAll // break
			}
			return nil // continue
		})
		if err != nil && !errors.Is(err, fs.SkipAll) {
			yield(File{}, err)
		}
	})
	first, err, _ := next()
	return &albumLoader{first, next, stop}, err
}

type albumLoader struct {
	first File
	next  func() (File, error, bool)
	stop  func()
}

// Load implements search.AlbumRepositoryLoader.
func (a *albumLoader) Load() (library.Album, error) {
	return a.first.Album(), nil
}

// Tracks implements search.AlbumRepositoryLoader.
func (a *albumLoader) Tracks() ([]search.Track, error) {
	defer a.stop()

	tracks := []search.Track{a.first.Track()}

	for {
		file, err, _continue := a.next()
		if !_continue {
			break
		}
		if err != nil {
			return nil, err
		}
		tracks = append(tracks, file.Track())
	}

	return tracks, nil
}

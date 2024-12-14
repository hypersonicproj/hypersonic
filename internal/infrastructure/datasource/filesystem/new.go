package filesystem

import (
	"hypersonic/internal/usecase/search"
	"io/fs"
	"os"
	"path"
)

type AbsoluteDirPath string

// Example usage:
//
//	fs := os.DirFS("/path/to/music-library")
//	repository := filesystem.NewRepository(fs)
//
// The provided fs.FS needs to follow the structure:
//
//	```
//	├── AlbumArtist
//	│   ├── AlbumName
//	│   │   ├── *.mp3
//	│   │   ├── *.wav
//	│   │   ├── *.m4a
//	```
//
// This structure allows the repository to organize and search tracks
// based on artist and album hierarchy.
func NewRepository(baseDir string) search.Dependencies {
	fsR := os.DirFS(baseDir)
	fsW := newFSW(baseDir)
	fs := filesystem{read: fsR, write: fsW}
	return search.Dependencies{
		AlbumsRepository:    newAlbumsRepository(fs.read),
		PlaylistsRepository: newPlaylistsRepository(fs.read),
	}
}

type filesystem struct {
	read  fs.FS
	write fsW
}

type fsW interface {
	Open(name string) (*os.File, error)
}

func newFSW(baseDir string) fsW {
	return fsWritable{baseDir}
}

type fsWritable struct {
	baseDir string
}

func (f fsWritable) Open(name string) (*os.File, error) {
	return os.Open(path.Join(f.baseDir, name))
}

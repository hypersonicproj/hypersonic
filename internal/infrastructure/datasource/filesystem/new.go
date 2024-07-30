package filesystem

import (
	"hypersonic/internal/usecase/search"
	"io/fs"
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
func NewRepository(fs fs.FS) search.Repository {
	return &filesystem{instance: fs}
}

type filesystem struct {
	instance fs.FS
}

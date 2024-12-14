package search

import (
	"iter"
)

func newAlbums(repository AlbumsRepository) Albums {
	return &albums{repository}
}

type albums struct {
	repository AlbumsRepository
}

// All implements Albums.
func (a *albums) All() iter.Seq2[Album, error] {
	return func(yield func(Album, error) bool) {
		for albumLoader, err := range a.repository.All() {
			if err != nil {
				if !yield(Album{}, err) {
					return // break
				}
			}

			album, err := albumLoader.Load()

			if !yield(Album{album, albumLoader}, err) {
				return // break
			}
		}
	}
}

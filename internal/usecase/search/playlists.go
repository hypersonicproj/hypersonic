package search

import "iter"

func newPlaylists(repository PlaylistsRepository) Playlists {
	return &playlists{repository}
}

type playlists struct {
	repository PlaylistsRepository
}

// All implements Playlists.
func (a *playlists) All() iter.Seq2[Playlist, error] {
	return func(yield func(Playlist, error) bool) {
		for playlistLoader, err := range a.repository.All() {
			if err != nil {
				if !yield(Playlist{}, err) {
					return // break
				}
			}

			playlist, err := playlistLoader.Load()

			if !yield(Playlist{playlist, playlistLoader}, err) {
				return // break
			}
		}
	}
}

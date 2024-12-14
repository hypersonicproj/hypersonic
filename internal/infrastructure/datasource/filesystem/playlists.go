package filesystem

import (
	"hypersonic/internal/usecase/search"
	"io/fs"
	"iter"
)

func newPlaylistsRepository(fs fs.FS) search.PlaylistsRepository {
	return &playlists{fs}
}

type playlists struct {
	fs fs.FS
}

// All implements search.PlaylistsRepository.
func (p *playlists) All() iter.Seq2[search.PlaylistRepositoryLoader, error] {
	panic("unimplemented")
}

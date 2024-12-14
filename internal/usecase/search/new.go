package search

import (
	"hypersonic/internal/domain/library"
	"iter"
)

type Search struct {
	Albums
	Playlists
}

func New(deps Dependencies) Search {
	return Search{newAlbums(deps.AlbumsRepository), newPlaylists(deps.PlaylistsRepository)}
}

type Dependencies struct {
	AlbumsRepository
	PlaylistsRepository
}

type AlbumsRepository interface {
	All() iter.Seq2[AlbumRepositoryLoader, error]
}

type AlbumRepositoryLoader interface {
	Load() (library.Album, error)
	TrackInspector
}

type PlaylistsRepository interface {
	All() iter.Seq2[PlaylistRepositoryLoader, error]
}

type PlaylistRepositoryLoader interface {
	Load() (library.Playlist, error)
	TrackInspector
}

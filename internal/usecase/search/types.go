package search

import "hypersonic/internal/domain"

type Search interface {
	FindPlaylistsCreatedAsc(...FindOptionApplier) ([]domain.Playlist, error)

	FindAlbumsAddedDesc(...FindOptionApplier) ([]domain.Album, error)
	FindAlbumsNameAsc(...FindOptionApplier) ([]domain.Album, error)

	FindArtistsNameAsc(...FindOptionApplier) ([]domain.Artist, error)

	FindTracksNameAsc(...FindOptionApplier) ([]domain.Track, error)
	FindTracksAddedAsc(...FindOptionApplier) ([]domain.Track, error)

	FindTracksInAlbum(domain.Album, ...FindOptionApplier) ([]domain.Track, error)
	FindTracksInPlaylist(domain.Playlist, ...FindOptionApplier) ([]domain.Track, error)
}

type Repository interface {
	FindPlaylistsCreatedAsc(FindOption) ([]domain.Playlist, error)

	FindAlbumsAddedDesc(FindOption) ([]domain.Album, error)
	FindAlbumsNameAsc(FindOption) ([]domain.Album, error)

	FindArtistsNameAsc(FindOption) ([]domain.Artist, error)

	FindTracksNameAsc(FindOption) ([]domain.Track, error)
	FindTracksAddedAsc(FindOption) ([]domain.Track, error)

	FindTracksInAlbum(domain.Album, FindOption) ([]domain.Track, error)
	FindTracksInPlaylist(domain.Playlist, FindOption) ([]domain.Track, error)
}

package search

import (
	"hypersonic/internal/domain"
	"time"
)

type cacheMutex struct {
	PlaylistsCreatedAsc []domain.Playlist
	AlbumsAddedAsc      []domain.Album

	TracksInAlbum         map[string] /* album id */ []domain.Track
	TracksInAlbumLFUIndex []struct {
		AlbumId       string
		ReadCount     int
		ReadTimestamp time.Time
	}

	TracksInPlaylist         map[string] /* playlist name */ map[int]domain.Track
	TracksInPlaylistLFUIndex []struct {
		playlistName  string
		ReadCount     int
		ReadTimestamp time.Time
	}

	AlbumsNameAsc  []domain.Album
	ArtistsNameAsc []domain.Artist
	TracksNameAsc  []domain.Track
}

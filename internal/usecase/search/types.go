package search

import (
	"hypersonic/internal/domain/library"
	"iter"
)

type Albums interface {
	All() iter.Seq2[Album, error]
}

type Album struct {
	library.Album
	TrackInspector
}

type Playlists interface {
	All() iter.Seq2[Playlist, error]
}

type Playlist struct {
	library.Playlist
	TrackInspector
}

type Track struct {
	Number int
	library.Track
}

type TrackInspector interface {
	// Tracks creates iterator iterate tracks sorted by track number and title
	Tracks() ([]Track, error)
}

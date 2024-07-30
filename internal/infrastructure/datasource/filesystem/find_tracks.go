package filesystem

import (
	"hypersonic/internal/domain"
	"hypersonic/internal/usecase/search"
)

func (s *filesystem) FindTracksAddedAsc(option search.FindOption) ([]domain.Track, error) {
	panic("unimplemented")
}

func (s *filesystem) FindTracksNameAsc(option search.FindOption) ([]domain.Track, error) {
	panic("unimplemented")
}

func (s *filesystem) FindTracksInAlbum(album domain.Album, option search.FindOption) ([]domain.Track, error) {
	panic("unimplemented")
}

func (s *filesystem) FindTracksInPlaylist(playlist domain.Playlist, option search.FindOption) ([]domain.Track, error) {
	panic("unimplemented")
}

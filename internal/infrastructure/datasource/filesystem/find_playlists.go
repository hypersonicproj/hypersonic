package filesystem

import (
	"hypersonic/internal/domain"
	"hypersonic/internal/usecase/search"
)

func (s *filesystem) FindPlaylistsCreatedAsc(option search.FindOption) ([]domain.Playlist, error) {
	panic("unimplemented")
}

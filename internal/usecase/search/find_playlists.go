package search

import "hypersonic/internal/domain"

func (s *search) FindPlaylistsCreatedAsc(o ...FindOptionApplier) ([]domain.Playlist, error) {
	option := DefaultFindOption()
	for _, apply := range o {
		apply(&option)
	}

	// TODO: cache control
	return s.repository.FindPlaylistsCreatedAsc(option)
}

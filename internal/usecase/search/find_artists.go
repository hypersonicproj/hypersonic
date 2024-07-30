package search

import "hypersonic/internal/domain"

func (s *search) FindArtistsNameAsc(o ...FindOptionApplier) ([]domain.Artist, error) {
	option := DefaultFindOption()
	for _, apply := range o {
		apply(&option)
	}

	// TODO: cache control
	return s.repository.FindArtistsNameAsc(option)
}

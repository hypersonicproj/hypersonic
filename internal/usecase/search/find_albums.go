package search

import "hypersonic/internal/domain"

func (s *search) FindAlbumsAddedDesc(o ...FindOptionApplier) ([]domain.Album, error) {
	option := DefaultFindOption()
	for _, apply := range o {
		apply(&option)
	}

	// TODO: cache control
	return s.repository.FindAlbumsAddedDesc(option)
}

func (s *search) FindAlbumsNameAsc(o ...FindOptionApplier) ([]domain.Album, error) {
	option := DefaultFindOption()
	for _, apply := range o {
		apply(&option)
	}

	// TODO: cache control
	return s.repository.FindAlbumsNameAsc(option)
}

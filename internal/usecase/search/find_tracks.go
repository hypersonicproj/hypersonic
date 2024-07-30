package search

import "hypersonic/internal/domain"

func (s *search) FindTracksAddedAsc(o ...FindOptionApplier) ([]domain.Track, error) {
	option := DefaultFindOption()
	for _, apply := range o {
		apply(&option)
	}

	// TODO: cache control
	return s.repository.FindTracksAddedAsc(option)
}

func (s *search) FindTracksNameAsc(o ...FindOptionApplier) ([]domain.Track, error) {
	option := DefaultFindOption()
	for _, apply := range o {
		apply(&option)
	}

	// TODO: cache control
	return s.repository.FindTracksNameAsc(option)
}

func (s *search) FindTracksInAlbum(album domain.Album, o ...FindOptionApplier) ([]domain.Track, error) {
	option := DefaultFindOption()
	for _, apply := range o {
		apply(&option)
	}

	// TODO: cache control
	return s.repository.FindTracksInAlbum(album, option)
}

func (s *search) FindTracksInPlaylist(playlist domain.Playlist, o ...FindOptionApplier) ([]domain.Track, error) {
	option := DefaultFindOption()
	for _, apply := range o {
		apply(&option)
	}

	// TODO: cache control
	return s.repository.FindTracksInPlaylist(playlist, option)
}

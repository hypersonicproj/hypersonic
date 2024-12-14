package handler

import (
	"hypersonic/internal/interface-adapter/handler/graphql/graph/model"
	"hypersonic/internal/usecase/search"
	"time"
)

func newModelAlbum(album search.Album, embedTracks bool) (model.Album, error) {
	albumModel := model.Album{
		ID:      album.Id().Text(),
		Title:   album.Title,
		Artist:  album.AlbumArtist.Name,
		AddedAt: album.AddedAt.Format(time.RFC3339),
	}
	if album.ReleasedAt != nil {
		albumModel.ReleaseDate = album.ReleasedAt.Format(time.DateOnly)
		albumModel.Year = album.ReleasedAt.Year()
	}
	if embedTracks {
		tracks, err := album.Tracks()
		if err != nil {
			return model.Album{}, err
		}
		for _, track := range tracks {
			albumModel.Tracks = append(albumModel.Tracks, &model.Track{
				ID:          track.Id().Text(),
				Title:       track.Title,
				Artist:      track.Artist.Name,
				Genre:       track.Genre,
				Album:       albumModel.Title,
				AlbumArtist: albumModel.Artist,
				ReleaseDate: track.ReleasedAt.Format(time.DateOnly),
				Year:        track.ReleasedAt.Year(),
				TrackNumber: track.AlbumTrackNumber,
				AddedAt:     track.AddedAt.Format(time.RFC3339),
			})
		}
	}
	return albumModel, nil
}

func newModelPlaylist(playlist search.Playlist, embedTracks bool) (model.Playlist, error) {
	playlistModel := model.Playlist{
		Name:      playlist.Name,
		CreatedAt: playlist.CreatedAt.Format(time.RFC3339),
	}
	if embedTracks {
		tracks, err := playlist.Tracks()
		if err != nil {
			return model.Playlist{}, err
		}
		for _, track := range tracks {
			playlistModel.Tracks = append(playlistModel.Tracks, &model.Track{
				ID:          track.Id().Text(),
				Title:       track.Title,
				Artist:      track.Artist.Name,
				Genre:       track.Genre,
				Album:       track.AlbumId.Text(),
				AlbumArtist: track.Artist.Name,
				ReleaseDate: track.ReleasedAt.Format(time.DateOnly),
				Year:        track.ReleasedAt.Year(),
				TrackNumber: &track.Number,
				AddedAt:     track.AddedAt.Format(time.RFC3339),
			})
		}
	}
	return playlistModel, nil
}

package handler

import (
	"hypersonic/internal/interface-adapter/handler/graphql/graph/model"
	"hypersonic/internal/pkg/tree"
	"strings"

	"github.com/tingtt/options"
)

func (h *handler) Albums(sortFunc SortYield[*model.Album], optionAppliers ...QueryOptionApplier) ([]*model.Album, error) {
	query := options.CreateWithDefault(defaultQueryOption(), optionAppliers...)

	var root *tree.Node[*model.Album]
	for album, err := range h.search.Albums.All() {
		if err != nil {
			return nil, err
		}

		albumModel, err := newModelAlbum(album, query.embedTracks)
		if err != nil {
			return nil, err
		}

		if query.filterByName != nil {
			if strings.Contains(albumModel.Title, *query.filterByName) {
				root = tree.Insert(root, &albumModel, sortFunc)
			}
			continue
		}
		root = tree.Insert(root, &albumModel, sortFunc)
	}
	albums := []*model.Album{}
	tree.InOrderTraversal(root, &albums)
	return albums, nil
}

func SortAlbumByTitle(fallbacks ...SortYieldMaker[*model.Album]) SortYield[*model.Album] {
	return func(new, curr *model.Album) (isLeft bool) {
		if new.Title == "" {
			return curr.Title == ""
		}
		if curr.Title == "" {
			return new.Title != ""
		}
		if /* same added date */ new.Title == curr.Title {
			if fallbacks == nil {
				return false
			}
			return fallbacks[0](fallbacks[1:]...)(new, curr)
		}
		return new.Title < curr.Title
	}
}

func SortAlbumByReleaseDate(fallbacks ...SortYieldMaker[*model.Album]) SortYield[*model.Album] {
	return func(new, curr *model.Album) (isLeft bool) {
		if new.ReleaseDate == "" {
			return curr.ReleaseDate == ""
		}
		if curr.ReleaseDate == "" {
			return new.ReleaseDate != ""
		}
		if /* same release date */ new.ReleaseDate == curr.ReleaseDate {
			if fallbacks == nil {
				return false
			}
			return fallbacks[0](fallbacks[1:]...)(new, curr)
		}
		return new.ReleaseDate < curr.ReleaseDate
	}
}

func SortAlbumByArtist(fallbacks ...SortYieldMaker[*model.Album]) SortYield[*model.Album] {
	return func(new, curr *model.Album) (isLeft bool) {
		if new.Artist == "" {
			return curr.Artist == ""
		}
		if curr.Artist == "" {
			return new.Artist != ""
		}
		if /* same release date */ new.Artist == curr.Artist {
			if fallbacks == nil {
				return false
			}
			return fallbacks[0](fallbacks[1:]...)(new, curr)
		}
		return new.Artist < curr.Artist
	}
}

func SortAlbumByAddedAt(fallbacks ...SortYieldMaker[*model.Album]) SortYield[*model.Album] {
	return func(new, curr *model.Album) (isLeft bool) {
		if new.AddedAt == "" {
			return curr.AddedAt == ""
		}
		if curr.AddedAt == "" {
			return new.AddedAt != ""
		}
		if /* same added date */ new.AddedAt == curr.AddedAt {
			if fallbacks == nil {
				return false
			}
			return fallbacks[0](fallbacks[1:]...)(new, curr)
		}
		return new.AddedAt < curr.AddedAt
	}
}

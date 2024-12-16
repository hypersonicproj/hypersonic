package handler

import (
	"hypersonic/internal/interface-adapter/handler/graphql/graph/model"
	"hypersonic/internal/pkg/tree"
	"strings"

	"github.com/tingtt/options"
)

func (h *handler) Playlists(sortFunc SortYield[*model.Playlist], optionAppliers ...QueryOptionApplier) ([]*model.Playlist, error) {
	query := options.CreateWithDefault(defaultQueryOption(), optionAppliers...)

	var root *tree.Node[*model.Playlist]
	for playlist, err := range h.search.Playlists.All() {
		if err != nil {
			return nil, err
		}

		playlistModel, err := newModelPlaylist(playlist, query.embedTracks)
		if err != nil {
			return nil, err
		}

		if query.filterByName != nil {
			if strings.Contains(playlistModel.Name, *query.filterByName) {
				root = tree.Insert(root, &playlistModel, sortFunc)
			}
			continue
		}
		root = tree.Insert(root, &playlistModel, sortFunc)
	}
	playlists := []*model.Playlist{}
	tree.InOrderTraversal(root, &playlists)
	return playlists, nil
}

func SortPlaylistByCreatedAt(fallbacks ...SortYieldMaker[*model.Playlist]) SortYield[*model.Playlist] {
	return func(new, curr *model.Playlist) (isLeft bool) {
		if new.CreatedAt == "" {
			return curr.CreatedAt == ""
		}
		if curr.CreatedAt == "" {
			return new.CreatedAt != ""
		}
		if /* same created time */ new.CreatedAt == curr.CreatedAt {
			if fallbacks == nil {
				return false
			}
			return fallbacks[0](fallbacks[1:]...)(new, curr)
		}
		return new.CreatedAt < curr.CreatedAt
	}
}

func SortPlaylistByName(fallbacks ...SortYieldMaker[*model.Playlist]) SortYield[*model.Playlist] {
	return func(new, curr *model.Playlist) (isLeft bool) {
		if new.Name == "" {
			return curr.Name == ""
		}
		if curr.Name == "" {
			return new.Name != ""
		}
		if /* same name */ new.Name == curr.Name {
			if fallbacks == nil {
				return false
			}
			return fallbacks[0](fallbacks[1:]...)(new, curr)
		}
		return new.Name < curr.Name
	}
}

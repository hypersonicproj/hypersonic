package handler

import "hypersonic/internal/interface-adapter/handler/graphql/graph/model"

type Handler interface {
	Albums(sortFunc SortYield[*model.Album], optionAppliers ...QueryOptionApplier) ([]*model.Album, error)
	Playlists(sortFunc SortYield[*model.Playlist], optionAppliers ...QueryOptionApplier) ([]*model.Playlist, error)
}

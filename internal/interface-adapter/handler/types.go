package handler

import "hypersonic/internal/interface-adapter/handler/graphql/graph/model"

type Handler interface {
	Albums(sortFunc SortYield[*model.Album], optionAppliers ...queryOptionApplier) ([]*model.Album, error)
	Playlists(sortFunc SortYield[*model.Playlist], optionAppliers ...queryOptionApplier) ([]*model.Playlist, error)
}

package graph

import (
	"context"
	"hypersonic/internal/interface-adapter/handler"
	"hypersonic/internal/interface-adapter/handler/graphql/graph/model"

	"github.com/stretchr/testify/mock"
)

var _ handler.Handler = new(MockHandler)

type MockHandler struct {
	mock.Mock
}

// Albums implements handler.Handler.
func (m *MockHandler) Albums(sortFunc handler.SortYield[*model.Album], optionAppliers ...handler.QueryOptionApplier) ([]*model.Album, error) {
	args := m.Called(sortFunc, optionAppliers)
	return args.Get(0).([]*model.Album), args.Error(1)
}

// Playlists implements handler.Handler.
func (m *MockHandler) Playlists(sortFunc handler.SortYield[*model.Playlist], optionAppliers ...handler.QueryOptionApplier) ([]*model.Playlist, error) {
	panic("unimplemented")
}

var _ fieldCollector = new(MockFieldCollector)

type MockFieldCollector struct {
	mock.Mock
}

// Collect implements fieldCollector.
func (m *MockFieldCollector) Collect(ctx context.Context) []string {
	args := m.Called(ctx)
	return args.Get(0).([]string)
}

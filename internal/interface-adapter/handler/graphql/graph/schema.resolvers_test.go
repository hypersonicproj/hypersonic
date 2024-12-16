package graph

import (
	"context"
	"hypersonic/internal/interface-adapter/handler"
	"hypersonic/internal/interface-adapter/handler/graphql/graph/model"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type AlbumsTest struct {
	caseName               string
	in                     AlbumsTest_In
	fieldCollectorBehavior AlbumsTest_FieldCollectorBehavior
	handlerBehavior        AlbumsTest_HandlerBehavior
	out                    AlbumsTest_Out
}

type AlbumsTest_In struct {
	sort         *model.SortAlbumsBy
	order        *model.Order
	filterByName *string
}

type AlbumsTest_FieldCollectorBehavior struct {
	ctx context.Context
	out []string
}

type AlbumsTest_HandlerBehavior struct {
	Albums *AlbumsTest_HandlerBehaviorAlbums
	// Playlists(sortFunc handler.SortYield[*model.Playlist], optionAppliers ...handler.QueryOptionApplier) ([]*model.Playlist, error)
}

type AlbumsTest_HandlerBehaviorAlbums struct {
	sortFunc       handler.SortYield[*model.Album]
	optionAppliers []handler.QueryOptionApplier
	out            AlbumsTest_Out
}

type AlbumsTest_Out struct {
	albums []*model.Album
	err    error
}

var testalbums = []AlbumsTest{
	{
		caseName: "may return album list sorted by release date desc when any args not specified (default call)",
		in:       AlbumsTest_In{nil, nil, nil},
		fieldCollectorBehavior: AlbumsTest_FieldCollectorBehavior{
			ctx: context.Background(),
			out: []string{},
		},
		handlerBehavior: AlbumsTest_HandlerBehavior{
			Albums: &AlbumsTest_HandlerBehaviorAlbums{
				sortFunc:       handler.SortAlbumByReleaseDate(),
				optionAppliers: []handler.QueryOptionApplier{},
				out: AlbumsTest_Out{
					albums: []*model.Album{},
					err:    nil,
				},
			},
		},
		out: AlbumsTest_Out{
			albums: []*model.Album{},
			err:    nil,
		},
	},
}

func Test_queryResolver_Albums(t *testing.T) {
	t.Parallel()

	for _, testalbum := range testalbums {
		t.Run(testalbum.caseName, func(t *testing.T) {
			t.Parallel()

			mockFieldCollector := new(MockFieldCollector)
			mockFieldCollector.On("Collect", mock.Anything).Return(testalbum.fieldCollectorBehavior.out)

			mockHandler := new(MockHandler)
			if testalbum.handlerBehavior.Albums != nil {
				mockHandler.On("Albums", mock.Anything, mock.Anything).Return(
					testalbum.handlerBehavior.Albums.out.albums, testalbum.handlerBehavior.Albums.out.err,
				)
			}

			rootResolver := &Resolver{mockHandler, mockFieldCollector}
			queryResolver := rootResolver.Query()
			got, err := queryResolver.Albums(context.Background(),
				testalbum.in.sort, testalbum.in.order, testalbum.in.filterByName,
			)

			// assert mock calls
			mockHandler.AssertNumberOfCalls(t, "Albums", 1)
			mockHandler.AssertCalled(t, "Albums",
				mock.Anything, //TODO: research how to test func type
				testalbum.handlerBehavior.Albums.optionAppliers,
			)
			// assert out
			assert.Equal(t, testalbum.out.albums, got)
			assert.Equal(t, testalbum.out.err, err)
		})
	}
}

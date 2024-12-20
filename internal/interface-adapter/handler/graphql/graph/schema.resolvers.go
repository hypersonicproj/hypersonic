package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.60

import (
	"context"
	"fmt"
	"hypersonic/internal/interface-adapter/handler"
	"hypersonic/internal/interface-adapter/handler/graphql/graph/model"
	"slices"

	"github.com/99designs/gqlgen/graphql"
)

// UploadTrack is the resolver for the uploadTrack field.
func (r *mutationResolver) UploadTrack(ctx context.Context, track model.NewTrack) (*model.Track, error) {
	panic(fmt.Errorf("not implemented: UploadTrack - uploadTrack"))
}

// UploadAlbum is the resolver for the uploadAlbum field.
func (r *mutationResolver) UploadAlbum(ctx context.Context, album model.NewAlbum) (*model.Album, error) {
	panic(fmt.Errorf("not implemented: UploadAlbum - uploadAlbum"))
}

// Albums is the resolver for the Albums field.
func (r *queryResolver) Albums(ctx context.Context, sort *model.SortAlbumsBy, order *model.Order, filterByName *string) ([]*model.Album, error) {
	sortFunc := handler.SortAlbumByReleaseDate()
	if sort != nil {
		switch *sort {
		case model.SortAlbumsByTitle:
			sortFunc = handler.SortAlbumByTitle()
		case model.SortAlbumsByRelease:
			sortFunc = handler.SortAlbumByReleaseDate()
		case model.SortAlbumsByArtist:
			sortFunc = handler.SortAlbumByArtist()
		case model.SortAlbumsByAdded:
			sortFunc = handler.SortAlbumByAddedAt()
		}
	}
	if order != nil && *order == model.OrderDesc {
		sortFunc = handler.Desc(sortFunc)
	}

	options := []handler.QueryOptionApplier{}
	if filterByName != nil {
		options = append(options, handler.WithFilterByName(*filterByName))
	}
	if slices.Contains(r.fieldCollector.Collect(ctx), "tracks") {
		options = append(options, handler.WithEmbedTracks())
	}
	return r.handler.Albums(sortFunc, options...)
}

// Playlists is the resolver for the Playlists field.
func (r *queryResolver) Playlists(ctx context.Context, sort *model.SortPlaylistsBy, order *model.Order, filterByName *string) ([]*model.Playlist, error) {
	sortFunc := handler.SortPlaylistByCreatedAt()
	if sort != nil {
		switch *sort {
		case model.SortPlaylistsByCreated:
			sortFunc = handler.SortPlaylistByCreatedAt()
		case model.SortPlaylistsByName:
			sortFunc = handler.SortPlaylistByName()
		}
	}
	if order != nil && *order == model.OrderDesc {
		sortFunc = handler.Desc(sortFunc)
	}

	options := []handler.QueryOptionApplier{}
	if filterByName != nil {
		options = append(options, handler.WithFilterByName(*filterByName))
	}
	if slices.Contains(graphql.CollectAllFields(ctx), "tracks") {
		options = append(options, handler.WithEmbedTracks())
	}
	return r.handler.Playlists(sortFunc, options...)
}

// Mutation returns MutationResolver implementation.
func (r *Resolver) Mutation() MutationResolver { return &mutationResolver{r} }

// Query returns QueryResolver implementation.
func (r *Resolver) Query() QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }

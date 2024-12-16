package graph

import (
	"context"
	"hypersonic/internal/interface-adapter/handler"
)

//go:generate go run github.com/99designs/gqlgen generate

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

func NewResolver(handler handler.Handler) ResolverRoot {
	return &Resolver{handler, &fieldCollectorImpl{}}
}

type Resolver struct {
	handler        handler.Handler
	fieldCollector fieldCollector
}

type fieldCollector interface {
	Collect(ctx context.Context) []string
}

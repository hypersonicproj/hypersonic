package graph

import (
	"hypersonic/internal/usecase/search"
)

//go:generate go run github.com/99designs/gqlgen generate

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

func NewResolver(deps Dependencies) ResolverRoot {
	return &Resolver{deps}
}

type Resolver struct {
	deps Dependencies
}

type Dependencies struct {
	Search search.Search
}

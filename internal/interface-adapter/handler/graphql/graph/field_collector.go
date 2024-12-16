package graph

import (
	"context"

	"github.com/99designs/gqlgen/graphql"
)

type fieldCollectorImpl struct{}

// Collect implements fieldCollector.
func (f *fieldCollectorImpl) Collect(ctx context.Context) []string {
	return graphql.CollectAllFields(ctx)
}

package graphql

import (
	"hypersonic/internal/interface-adapter/handler/graphql/graph"
	"net/http"

	"github.com/99designs/gqlgen/graphql/handler"
)

func NewHandler(deps graph.Dependencies) http.Handler {
	return handler.NewDefaultServer(
		graph.NewExecutableSchema(graph.Config{
			Resolvers: graph.NewResolver(deps),
		}),
	)
}

package graphql

import (
	"hypersonic/internal/interface-adapter/handler"
	"hypersonic/internal/interface-adapter/handler/graphql/graph"
	"hypersonic/internal/usecase/search"
	"net/http"

	gqlhandler "github.com/99designs/gqlgen/graphql/handler"
)

func NewHandler(deps Dependencies) http.Handler {
	return gqlhandler.NewDefaultServer(
		graph.NewExecutableSchema(graph.Config{
			Resolvers: graph.NewResolver(handler.New(deps.Search)),
		}),
	)
}

type Dependencies struct {
	Search search.Search
}

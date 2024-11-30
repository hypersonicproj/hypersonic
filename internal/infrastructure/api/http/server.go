package http

import (
	"hypersonic/internal/interface-adapter/handler/graphql"
	"hypersonic/internal/interface-adapter/handler/graphql/graph"
	"net/http"

	"github.com/99designs/gqlgen/graphql/playground"
	"golang.org/x/net/http2"
	"golang.org/x/net/http2/h2c"
)

type Server struct {
	addr string
	deps graph.Dependencies
}

func NewServer(addr string, deps graph.Dependencies) *Server {
	return &Server{addr, deps}
}

func (s *Server) Serve() error {
	mux := http.NewServeMux()

	// `/hypersonic.v1graphql.MusicLibrary/`
	mux.Handle("/hypersonic.v1graphql.MusicLibrary/", graphql.NewHandler(s.deps))
	// `/hypersonic.v1graphql.MusicLibrary/playground`
	mux.Handle("/hypersonic.v1graphql.MusicLibrary/playground", playground.Handler("GraphQL playground", "/hypersonic.v1graphql.MusicLibrary/"))

	return http.ListenAndServe(s.addr, h2c.NewHandler(mux, &http2.Server{}))
}
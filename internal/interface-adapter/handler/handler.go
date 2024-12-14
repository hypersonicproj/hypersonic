package handler

import "hypersonic/internal/usecase/search"

func New(search search.Search) Handler {
	return &handler{search}
}

type handler struct {
	search search.Search
}

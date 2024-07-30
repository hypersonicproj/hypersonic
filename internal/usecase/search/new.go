package search

func New(r Repository) Search {
	return &search{r, cacheMutex{}}
}

type search struct {
	repository Repository
	cacheMutex cacheMutex
}

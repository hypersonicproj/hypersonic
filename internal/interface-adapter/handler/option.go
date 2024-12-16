package handler

import "github.com/tingtt/options"

func defaultQueryOption() queryOption {
	return *options.Create[queryOption]()
}

type queryOption struct {
	offset       *int
	limit        *int
	filterByName *string
	embedTracks  bool
}

type QueryOptionApplier = options.Applier[queryOption]

func WithOffset(offset int) QueryOptionApplier {
	return func(o *queryOption) {
		o.offset = &offset
	}
}

func WithLimit(limit int) QueryOptionApplier {
	return func(o *queryOption) {
		o.limit = &limit
	}
}

func WithFilterByName(name string) QueryOptionApplier {
	return func(o *queryOption) {
		o.filterByName = &name
	}
}

func WithEmbedTracks() QueryOptionApplier {
	return func(o *queryOption) {
		o.embedTracks = true
	}
}

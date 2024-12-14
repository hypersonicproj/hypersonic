package handler

import "github.com/tingtt/options"

func defaultQueryOption() queryOption {
	return *options.Create[queryOption]()
}

type queryOption struct {
	embedTracks bool
}

type queryOptionApplier = options.Applier[queryOption]

func WithEmbedTracks() queryOptionApplier {
	return func(o *queryOption) {
		o.embedTracks = true
	}
}

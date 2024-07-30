package search

import "time"

func DefaultFindOption() FindOption {
	deadline := time.After(30 * time.Second)
	return FindOption{
		TimeoutDeadline: &deadline,
	}
}

type FindOption struct {
	Offset          *int
	Limit           *int
	SearchByName    *string
	TimeoutDeadline *<-chan time.Time
}

type FindOptionApplier func(*FindOption)

func WithOffset(offset int) FindOptionApplier {
	return func(o *FindOption) {
		o.Offset = &offset
	}
}

func WithLimit(limit int) FindOptionApplier {
	return func(o *FindOption) {
		o.Limit = &limit
	}
}

func WithSearchByName(name string) FindOptionApplier {
	return func(o *FindOption) {
		o.SearchByName = &name
	}
}

func WithTimeout(deadline <-chan time.Time) FindOptionApplier {
	return func(o *FindOption) {
		o.TimeoutDeadline = &deadline
	}
}

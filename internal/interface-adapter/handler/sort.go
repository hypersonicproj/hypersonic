package handler

type SortYield[T any] func(new, curr T) (isLeft bool)
type SortYieldMaker[T any] func(fallbacks ...SortYieldMaker[T]) SortYield[T]

func Desc[T any](yield SortYield[T]) SortYield[T] {
	return func(new, curr T) (isLeft bool) {
		return !yield(new, curr)
	}
}

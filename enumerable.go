package js_array_methods


type Enumerable[T any] struct {
	items []T
}

func NewEnumerable[T any](items []T) Enumerable[T] {
	return Enumerable[T]{items: items}
}

func (e *Enumerable[T]) Map(fn func(*T)) {
	for i := range e.items {
		fn(&(e.items[i]))
	}
}

func (e *Enumerable[T]) Filter(fn func(*T) bool) {
	copy_items := make([]T, 0)

	for i := range e.items {
		if fn(&(e.items[i])) {
			copy_items = append(copy_items, e.items[i])
		}
	}
	e.items = copy_items
}

func (e *Enumerable[T]) Reduce(fn func(*T, *T) T, starting_point T) T {
	prev := starting_point
	curr := starting_point

	for _, item := range e.items {
		curr = item
		prev = fn(&prev, &curr)
	}

	return prev
}


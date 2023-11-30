package js_array_methods

type Enumerable[T any] struct {
	items []T
}

func ToEnumerable[T any](items []T) Enumerable[T] {
	return Enumerable[T]{items: items}
}

func (e *Enumerable[T]) ToList() []T {
	return e.items
}

// Call an anonymous function across each element in your list of items to change values.
func (e *Enumerable[T]) Map(fn func(*T)) *Enumerable[T] {
	for i := range e.items {
		fn(&(e.items[i]))
	}
	return e
}

// From C# works similarly to Map but instead returns a new value
// Call an anonymous function across each element in your list of items to change values.
func (e *Enumerable[T]) Select(fn func(T) interface{}) (new_list []any) {
	for i := range e.items {
		new_list = append(new_list, fn(e.items[i]))
	}
	return new_list
}

// Call an anonymous function across each element in your list of items to filter out unwanted items.
func (e *Enumerable[T]) Filter(fn func(*T) bool) *Enumerable[T] {
	copy_items := make([]T, 0)

	for i := range e.items {
		if fn(&(e.items[i])) {
			copy_items = append(copy_items, e.items[i])
		}
	}
	e.items = copy_items
	return e
}

// Squash your list of items into the same type using a criteria specified in a callback method using a starting value.
//
// - Outputs a value unlike the rest of the functional methods in this package
func (e *Enumerable[T]) Reduce(fn func(*T, *T) T, starting_point T) T {
	prev := starting_point
	for _, item := range e.items {
		prev = fn(&prev, &item)
	}

	return prev
}

// Squash your list of items into a single integer using a criteria specified in a callback method using a starting value.
//
// - Outputs a value unlike the rest of the functional methods in this package
func (e *Enumerable[T]) ReduceInt(fn func(*int, *T) int, starting_point int) int {
	prev := starting_point
	for _, item := range e.items {
		prev = fn(&prev, &item)
	}

	return prev
}

// Squash your list of items into a single string using a criteria specified in a callback method using a starting value.
//
// - Outputs a value unlike the rest of the functional methods in this package
func (e *Enumerable[T]) ReduceString(fn func(*string, *T) string, starting_point string) string {
	prev := starting_point
	for _, item := range e.items {
		prev = fn(&prev, &item)
	}

	return prev
}

// returns the number of items in the enumerable
func (e *Enumerable[T]) Count() int {
	return len(e.items)
}

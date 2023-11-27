# JavaScript array methods in golang 🤯

This package is just for fun and should not seriously be used in production.

This includes

| Method Name   | What it does                                                                                                          |
| ------------- | --------------------------------------------------------------------------------------------------------------------- |
| `.Map()`    | will perform the callback function action on each member of the `enumerable` object                                 |
| `.Filter()` | will remove items in the `enumerable` object that do not conform to the filters passed callback                     |
| `.Reduce()` | this will allow for an enumerable object to collapse into a single value using a callback method and a starting value |

## Example usage

Import the package

```go
import (
	js_array_methods "github.com/callummclu/js-array-methods-in-go"
)
```

Create the Enumerable

```go
nums := js_array_methods.NewEnumerable([]int{1,2,3,4,5})
```

#### `.Map()`

Call method on enumerable

```go
nums.Map(func(x *int){
	*(x) * 2
})
```

Example output

```go
[ 2 4 6 8 10 ]
```

#### `.Filter()`

Call method on enumerable

```go
nums.Filter(func(x *int) bool {
	return *(x) % 2 == 0
})
```

Example output

```go
[2 4]
```

#### `.Reduce()`

Call method on enumerable

```go
nums.Reduce(func(prev, curr *int){
	return *prev + *curr
}, 0)
```

Example output

```go
15
```

All of these method should also function with any other data type not just int

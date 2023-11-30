# JavaScript array and LINQ methods in golang 🤯

This package is just for fun and should not seriously be used in production.

This includes

| Method Name         | What it does                                                                                                          |
| ------------------- | --------------------------------------------------------------------------------------------------------------------- |
| `.Map()`          | will perform the callback function action on each member of the `enumerable` object                                 |
| `.Filter()`       | will remove items in the `enumerable` object that do not conform to the filters passed callback                     |
| `.Reduce()`       | this will allow for an enumerable object to collapse into a single value using a callback method and a starting value |
| `.ReduceInt()`    | Due to go not allowing generics on struct methods this allows for custom structs to be reduced into an integer value  |
| `.ReduceString()` | Again due to go not allowing generics on struct methods this allows for customer structs to be reduced into a string  |
| `.Select()`       | Straight out of `C#`'s LINQ this method allows for mapping an Enumerable list into a different type                 |
| `.ToList()`       | As is sounds, converts the enumerable type into a standard slice                                                      |
| `.Count()`        | unneccessary but stops you having to do `len(my_enum.ToList())`, grabs the number of items in the enumerable        |

## Example usage

Import the package

```go
import (
	js_array_methods "github.com/callummclu/js-array-methods-in-go"
)
```

Create the Enumerable

```go
nums := js_array_methods.ToEnumerable([]int{1,2,3,4,5})
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
total := nums.Reduce(func(prev, curr *int){
	return *prev + *curr
}, 0)
```

Example output

```go
15
```

#### `.ReduceInt()`

Call method on enumerable

```go
type Person struct{
	age int
}

people := ToEnumerable([]Person{{age:18},{age:19},{age:20},{age:21},{age:22}})

average_age := people.ReduceInt(func(prev *int, curr *Person) int{
	return *prev + curr.age
},0) / people.Count()
```

Example Output

```go
20
```

#### `.ReduceString()`

Call method on enumerable

```go
type Person struct{
	name string
}

people := ToEnumerable([]Person{{name:"gene"},{name:"oliver"},{name:"lewis"},{name:"andrew"},{name:"norman"},{name:"geoff"}})

team_name := people.ReduceString(func(prev *string, curr *Person) string{
	return *prev + curr.name[0:1]
},"")
```

Example Output

```go
"golang"
```

#### `.Select()`

Call method on enumerable

```go
type Person struct{
	age int
}

type Adult struct{
	is_above_18 bool
	age int
}

people := ToEnumerable([]Person{{age:10},{age:20},{age:18},{age:12}})

adults := people.Filter(func(p *Person) bool{
	return p.age >= 18
}).Select(func(p Person) interface{}{
	return Adult{is_above_18:true, age:p.age}
})
```

Example Output

```go
{true 20} {true 18}
```

#### `.ToList()`

Call method on enumerable

```go
nums := ToEnumerable([]int{1,2,3,4}})

nums.ToList()
```

Example output

```go
[ 1 2 3 4 ] // as a standard slice
```

#### `.Count()`

Call method on enumerable

```go
nums := ToEnumerable([]int{1,2,3,4}})

nums.Count()
```

Example output

```go
4
```


All of these method should also function with any other data type not just int

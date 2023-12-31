# Pythonic

Simple and easy to use collecting of `Python` functionality that is not easily available in `Go`.

# Install 

Simply run `go get github.com/westwardharbor0/pythonic` in your project root.

For using start typing `pythonic.`.

# Functions

- `IN(slice, value) bool` - Returns if element is present in slice.
- `IndexOf(slice, value) int` - Returns index of the value in slice. Not found is -1.
- `Set(slice, value) slice` - Deletes duplicates from slice and returns slice.
- `Delete(slice, value) slice` - Deletes the item by value from slice and returns slice.
- `Counter(slice) map` - Counts occurrences of items in string slice.

Will be adding more.

# Example

```go
package main

import "github.com/westwardharbor0/pythonic"

func main() {
	testSlice := []int{1, 9, 10}
	println(pythonic.IN(testSlice, 1)) // -> true
}
```

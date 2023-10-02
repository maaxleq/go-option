# go-option

The `option` package provides a generic implementation of an Option type in Go, representing a value that may or may not be present. It is designed to help developers write clean, concise, and error-free code by avoiding nil references and explicit error checks.

## Installation

To install the `option` package, you can use `go get`:

```sh
go get github.com/maaxleq/go-option
```

## Usage

Here’s a quick overview of how to use the `option` package:

```go
package main

import (
	"fmt"
	"github.com/maaxleq/go-option/option"
)

func main() {
	val := 5
	opt := option.Some(val)

	opt.IfPresent(func(v int) {
		fmt.Println(v) // Will print 5
	})
}
```

### Creating Options

- **Some(value T):** Creates an Option containing the given value.
- **None[T any]():** Creates an empty Option for the type T.
- **NewOptionFromValueOrError[T any](val T, err error):** Creates an Option based on a value and an error. If the error is non-nil, it returns an empty Option. If the error is nil, it returns an Option containing the provided value.

### Working with Options

- **IsPresent():** Returns true if the Option contains a value, false otherwise.
- **IfPresent(consumer func(T)):** Performs the given action with the value inside the Option if it is present, doing nothing if the Option is empty.
- **Get():** Retrieves the value contained in the Option or returns an error if the Option is empty.
- **OrElse(defaultValue T):** Returns the value contained in the Option if present, otherwise it returns the provided default value.
- **OrElseGet(provider func() T):** Returns the value contained in the Option if present, otherwise it computes and returns a value using the provided provider function.
- **OrElseError(err error):** Returns the value contained in the Option if present, otherwise it returns the provided error.
- **Filter(predicate func(T) bool):** Returns an Option containing the value of the original Option if it is present and matches the predicate, otherwise it returns an empty Option.

## Documentation

For more detailed documentation, please visit [pkg.go.dev](https://pkg.go.dev/github.com/maaxleq/go-option).

## License

This project is licensed under the MIT license

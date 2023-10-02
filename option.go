// Package option provides a generic implementation of an Option type in Go,
// representing a value that may or may not be present.
package option

// ErrNoElement represents an error returned when trying to access
// an element from an empty Option.
type ErrNoElement struct{}

// Error implements the error interface for ErrNoElement, returning
// a descriptive error message.
func (ErrNoElement) Error() string {
	return "option: no element is present"
}

// zeroValueOfAny returns the zero value of the type T.
func zeroValueOfAny[T any]() T {
	return *new(T)
}

// Option represents a container that may or may not hold a value of a given type.
type Option[T any] struct {
	val *T
}

// Some returns an Option containing the given value.
func Some[T any](value T) Option[T] {
	return Option[T]{&value}
}

// None returns an empty Option for the type T.
func None[T any]() Option[T] {
	return Option[T]{nil}
}

// NewOptionFromPtr creates a new Option from a pointer to a value of type T.
// If the pointer is nil, the returned Option will be empty.
func NewOptionFromPtr[T any](ptr *T) Option[T] {
	return Option[T]{ptr}
}

// NewOptionFromValueOrError creates a new Option based on a value and an error.
// If the error is non-nil, it returns an empty Option. If the error is nil,
// it returns an Option containing the provided value.
func NewOptionFromValueOrError[T any](val T, err error) Option[T] {
	if err != nil {
		return None[T]()
	}
	return Some(val)
}

// Get retrieves the value contained in the Option or returns an error
// if the Option is empty.
func (opt *Option[T]) Get() (T, error) {
	if opt.val == nil {
		return zeroValueOfAny[T](), ErrNoElement{}
	}
	return *opt.val, nil
}

// IsPresent returns true if the Option contains a value, false otherwise.
func (opt *Option[T]) IsPresent() bool {
	return opt.val != nil
}

// OrElse returns the value contained in the Option if present,
// otherwise it returns the provided default value.
func (opt *Option[T]) OrElse(defaultValue T) T {
	if opt.val == nil {
		return defaultValue
	}
	return *opt.val
}

// OrElseGet returns the value contained in the Option if present,
// otherwise it computes and returns a value using the provided provider function.
func (opt *Option[T]) OrElseGet(provider func() T) T {
	if opt.val == nil {
		return provider()
	}
	return *opt.val
}

// OrElseError returns the value contained in the Option if present,
// otherwise it returns the provided error.
func (opt *Option[T]) OrElseError(err error) (T, error) {
	if opt.val == nil {
		return zeroValueOfAny[T](), err
	}
	return *opt.val, nil
}

// Filter returns an Option containing the value of the original Option
// if it is present and matches the predicate, otherwise it returns an empty Option.
func (opt *Option[T]) Filter(predicate func(T) bool) Option[T] {
	if opt.val != nil && predicate(*opt.val) {
		return Some[T](*opt.val)
	}
	return None[T]()
}

// IfPresent performs the given action with the value inside the Option
// if it is present, doing nothing if the Option is empty.
// The provided consumer function takes the value of the Option as a parameter.
func (opt *Option[T]) IfPresent(consumer func(T)) {
	if opt.val != nil {
		consumer(*opt.val)
	}
}

package option_test

import (
	"errors"
	"testing"

	"github.com/maaxleq/go-option"
)

func TestSome(t *testing.T) {
	val := 5
	opt := option.Some(val)

	if !opt.IsPresent() {
		t.Errorf("Expected Option to be present")
	}

	if got, _ := opt.Get(); got != val {
		t.Errorf("Expected %d, got %d", val, got)
	}
}

func TestNone(t *testing.T) {
	opt := option.None[int]()

	if opt.IsPresent() {
		t.Errorf("Expected Option to be empty")
	}

	if _, err := opt.Get(); err == nil {
		t.Errorf("Expected error, got nil")
	}
}

func TestNewOptionFromPtr(t *testing.T) {
	val := 5
	opt := option.NewOptionFromPtr(&val)

	if !opt.IsPresent() {
		t.Errorf("Expected Option to be present")
	}

	if got, _ := opt.Get(); got != val {
		t.Errorf("Expected %d, got %d", val, got)
	}

	opt = option.NewOptionFromPtr[int](nil)

	if opt.IsPresent() {
		t.Errorf("Expected Option to be empty")
	}
}

func TestNewOptionFromValueOrError(t *testing.T) {
	val := 5
	err := errors.New("some error")

	opt := option.NewOptionFromValueOrError(val, err)
	if opt.IsPresent() {
		t.Errorf("Expected Option to be empty")
	}

	opt = option.NewOptionFromValueOrError(val, nil)
	if !opt.IsPresent() {
		t.Errorf("Expected Option to be present")
	}

	if got, _ := opt.Get(); got != val {
		t.Errorf("Expected %d, got %d", val, got)
	}
}

func TestOrElse(t *testing.T) {
	val := 5
	defaultVal := 10
	opt := option.Some(val)

	if got := opt.OrElse(defaultVal); got != val {
		t.Errorf("Expected %d, got %d", val, got)
	}

	opt = option.None[int]()

	if got := opt.OrElse(defaultVal); got != defaultVal {
		t.Errorf("Expected %d, got %d", defaultVal, got)
	}
}

func TestOrElseGet(t *testing.T) {
	val := 5
	defaultVal := 10
	opt := option.Some(val)

	if got := opt.OrElseGet(func() int { return defaultVal }); got != val {
		t.Errorf("Expected %d, got %d", val, got)
	}

	opt = option.None[int]()

	if got := opt.OrElseGet(func() int { return defaultVal }); got != defaultVal {
		t.Errorf("Expected %d, got %d", defaultVal, got)
	}
}

func TestOrElseError(t *testing.T) {
	val := 5
	err := errors.New("error")
	opt := option.Some(val)

	if got, e := opt.OrElseError(err); got != val || e != nil {
		t.Errorf("Expected %d and nil error, got %d and %v", val, got, e)
	}

	opt = option.None[int]()

	if _, e := opt.OrElseError(err); e == nil {
		t.Errorf("Expected error, got nil")
	}
}

func TestFilter(t *testing.T) {
	val := 5
	opt := option.Some(val)

	opt = opt.Filter(func(v int) bool { return v == 5 })

	if !opt.IsPresent() {
		t.Errorf("Expected Option to be present")
	}

	opt = opt.Filter(func(v int) bool { return v != 5 })

	if opt.IsPresent() {
		t.Errorf("Expected Option to be empty")
	}
}

func TestIfPresent(t *testing.T) {
	val := 5
	opt := option.Some(val)

	called := false
	opt.IfPresent(func(v int) {
		called = true
		if v != val {
			t.Errorf("Expected %d, got %d", val, v)
		}
	})

	if !called {
		t.Errorf("Expected consumer function to be called")
	}

	opt = option.None[int]()
	called = false
	opt.IfPresent(func(v int) {
		called = true
	})

	if called {
		t.Errorf("Expected consumer function to not be called")
	}
}

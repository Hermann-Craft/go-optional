package optional

import (
	"errors"
	"fmt"
	"testing"
)

func TestOptionalEmpty(t *testing.T) {
	opt := Empty[int]()
	if opt.IsPresent() {
		t.Errorf("Expected empty optional, but value was present")
	}
	if !opt.IsEmpty() {
		t.Errorf("Expected optional to be empty, but it was not")
	}
}

func TestOptionalOf(t *testing.T) {
	val := 42
	opt := Of(val)
	if !opt.IsPresent() {
		t.Errorf("Expected value to be present, but it was not")
	}
	if opt.Get() != val {
		t.Errorf("Expected value %d, but got %d", val, opt.Get())
	}
}

func TestOptionalOfWithNil(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Errorf("Expected panic for nil value, but did not panic")
		}
	}()
	_ = Of[*int](nil) // Should panic
}

func TestOptionalOfNullable(t *testing.T) {
	val := 42
	opt := OfNullable(&val)
	if !opt.IsPresent() {
		t.Errorf("Expected value to be present, but it was not")
	}
	if opt.Get() != val {
		t.Errorf("Expected value %d, but got %d", val, opt.Get())
	}

	optEmpty := OfNullable[*int](nil)
	if optEmpty.IsPresent() {
		t.Errorf("Expected no value, but got a present value")
	}
}

func TestOptionalGet(t *testing.T) {
	opt := Of(42)
	if opt.Get() != 42 {
		t.Errorf("Expected value 42, but got %d", opt.Get())
	}

	defer func() {
		if r := recover(); r == nil {
			t.Errorf("Expected panic for empty optional, but did not panic")
		}
	}()
	empty := Empty[int]()
	_ = empty.Get() // Should panic
}

func TestOptionalIfPresent(t *testing.T) {
	opt := Of(42)
	called := false
	opt.IfPresent(func(val int) {
		called = true
		if val != 42 {
			t.Errorf("Expected value 42, but got %d", val)
		}
	})
	if !called {
		t.Errorf("Expected IfPresent action to be called, but it was not")
	}

	empty := Empty[int]()
	empty.IfPresent(func(val int) {
		t.Errorf("Action should not be called for empty optional")
	})
}

func TestOptionalIfPresentOrElse(t *testing.T) {
	opt := Of(42)
	called := false
	emptyCalled := false
	opt.IfPresentOrElse(
		func(val int) {
			called = true
			if val != 42 {
				t.Errorf("Expected value 42, but got %d", val)
			}
		},
		func() {
			emptyCalled = true
		},
	)
	if !called {
		t.Errorf("Expected IfPresentOrElse action to be called, but it was not")
	}
	if emptyCalled {
		t.Errorf("Empty action should not be called when value is present")
	}

	empty := Empty[int]()
	emptyCalled = false
	empty.IfPresentOrElse(
		func(val int) {
			t.Errorf("Present action should not be called for empty optional")
		},
		func() {
			emptyCalled = true
		},
	)
	if !emptyCalled {
		t.Errorf("Expected empty action to be called, but it was not")
	}
}

func TestOptionalOrElse(t *testing.T) {
	opt := Empty[int]()
	val := opt.OrElse(100)
	if val != 100 {
		t.Errorf("Expected default value 100, but got %d", val)
	}

	opt = Of(42)
	val = opt.OrElse(100)
	if val != 42 {
		t.Errorf("Expected value 42, but got %d", val)
	}
}

func TestOptionalOrElseGet(t *testing.T) {
	opt := Empty[int]()
	val := opt.OrElseGet(func() int {
		return 100
	})
	if val != 100 {
		t.Errorf("Expected computed value 100, but got %d", val)
	}

	opt = Of(42)
	val = opt.OrElseGet(func() int {
		return 100
	})
	if val != 42 {
		t.Errorf("Expected value 42, but got %d", val)
	}
}

func TestOptionalOrElseThrow(t *testing.T) {
	opt := Of(42)
	val := opt.OrElseThrow(errors.New("error"))
	if val != 42 {
		t.Errorf("Expected value 42, but got %d", val)
	}

	defer func() {
		if r := recover(); r == nil {
			t.Errorf("Expected panic for empty optional, but did not panic")
		}
	}()
	empty := Empty[int]()
	_ = empty.OrElseThrow(errors.New("error")) // Should panic
}

func TestOptionalMap(t *testing.T) {
	opt := Of(42)
	mapped := Map(opt, func(val int) string {
		return fmt.Sprintf("Value: %d", val)
	})
	if !mapped.IsPresent() || mapped.Get() != "Value: 42" {
		t.Errorf("Expected mapped value 'Value: 42', but got %v", mapped.Get())
	}

	empty := Empty[int]()
	mappedEmpty := Map(empty, func(val int) string {
		return "This should not be called"
	})
	if mappedEmpty.IsPresent() {
		t.Errorf("Expected mapped optional to be empty, but it was not")
	}
}

func TestOptionalFlatMap(t *testing.T) {
	opt := Of(42)
	flatMapped := FlatMap(opt, func(val int) Optional[string] {
		return Of(fmt.Sprintf("Value is %d", val))
	})
	if !flatMapped.IsPresent() || flatMapped.Get() != "Value is 42" {
		t.Errorf("Expected flat-mapped value 'Value is 42', but got %v", flatMapped.Get())
	}

	empty := Empty[int]()
	flatMappedEmpty := FlatMap(empty, func(val int) Optional[string] {
		return Of("This should not be called")
	})
	if flatMappedEmpty.IsPresent() {
		t.Errorf("Expected flat-mapped optional to be empty, but it was not")
	}
}

func TestOptionalString(t *testing.T) {
	opt := Of(42)
	if opt.String() != "Optional[42]" {
		t.Errorf("Expected string 'Optional[42]', but got %s", opt.String())
	}

	empty := Empty[int]()
	if empty.String() != "Optional.empty" {
		t.Errorf("Expected string 'Optional.empty', but got %s", empty.String())
	}
}

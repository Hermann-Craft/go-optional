package optional

import (
	"fmt"
	"reflect"
)

// Optional represents a container that may or may not hold a value.
type Optional[T any] struct {
	value *T
}

// Empty creates an empty Optional instance.
func Empty[T any]() Optional[T] {
	return Optional[T]{value: nil}
}

// Of creates an Optional containing a non-nil value.
// It panics if the value is nil to ensure explicit non-null usage.
func Of[T any](value T) Optional[T] {
	// Vérifie explicitement si le type est un pointeur et si la valeur est nil
	if isNil(value) {
		panic("Optional.Of: value cannot be nil")
	}
	return Optional[T]{value: &value}
}

// isNil checks if a generic value is nil.
func isNil[T any](value T) bool {
	// Uses reflection to check if the value is a nil pointer.
	v := reflect.ValueOf(value)
	return v.Kind() == reflect.Ptr && v.IsNil()
}

// OfNullable creates an Optional containing the value if it is non-nil, otherwise an empty Optional.
// Supports cases where the input value is nil.
func OfNullable[T any](value *T) Optional[T] {
	return Optional[T]{value: value}
}

// IsPresent returns true if the Optional contains a value.
func (o Optional[T]) IsPresent() bool {
	return o.value != nil
}

// IsEmpty returns true if the Optional does not contain a value.
func (o Optional[T]) IsEmpty() bool {
	return o.value == nil
}

// Get returns the value if present, otherwise it panics.
func (o Optional[T]) Get() T {
	if o.IsEmpty() {
		panic("Optional.Get: no value present")
	}
	return *o.value
}

// IfPresent performs the given action with the value if it is present.
func (o Optional[T]) IfPresent(action func(T)) {
	if o.IsPresent() {
		action(*o.value)
	}
}

// IfPresentOrElse performs the given action with the value if it is present,
// otherwise performs the given empty action.
func (o Optional[T]) IfPresentOrElse(action func(T), emptyAction func()) {
	if o.IsPresent() {
		action(*o.value)
	} else {
		emptyAction()
	}
}

// OrElse returns the value if present, otherwise returns the provided default value.
func (o Optional[T]) OrElse(other T) T {
	if o.IsPresent() {
		return *o.value
	}
	return other
}

// OrElseGet returns the value if present, otherwise computes it using the given supplier.
func (o Optional[T]) OrElseGet(supplier func() T) T {
	if o.IsPresent() {
		return *o.value
	}
	return supplier()
}

// OrElseThrow returns the value if present, otherwise it panics with the provided error.
func (o Optional[T]) OrElseThrow(err error) T {
	if o.IsPresent() {
		return *o.value
	}
	panic(err)
}

// Map applies the given function to the value if present and returns an Optional describing the result.
func Map[T, U any](opt Optional[T], mapper func(T) U) Optional[U] {
	if opt.IsEmpty() {
		return Empty[U]()
	}
	return Of(mapper(opt.Get()))
}

// FlatMap applies the given function to the value if present and returns the result directly.
func FlatMap[T, U any](opt Optional[T], mapper func(T) Optional[U]) Optional[U] {
	if opt.IsEmpty() {
		return Empty[U]()
	}
	return mapper(opt.Get())
}

// String returns a string representation of the Optional.
func (o Optional[T]) String() string {
	if o.IsPresent() {
		return fmt.Sprintf("Optional[%v]", *o.value)
	}
	return "Optional.empty"
}

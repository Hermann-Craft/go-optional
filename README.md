# Optional<T> for Go

`Optional<T>` is a generic implementation in Go inspired by Java's `Optional`. It provides a type-safe and expressive way to handle values that might be absent, avoiding the use of `nil` directly and reducing the risk of null pointer errors.

## Features

- Create `Optional` instances with or without a value.
- Check if a value is present or absent.
- Access the value safely, with alternatives for default values or lazy computation.
- Apply transformations to the value using `Map` or `FlatMap`.
- Perform actions conditionally if a value is present or absent.
- Simple and readable API inspired by functional programming.

---

## Installation

```bash
go get github.com/hermann-craft/go-optional
```
---

## Usage

### Import the Package

```go
import "github.com/hermann-craft/go-optional"
```

### Create an Optional Instance

#### With a Non-Empty Value
```go
opt := optional.Of(42) // Create an Optional containing 42.
fmt.Println(opt)       // Output: Optional[42]
```

#### With a Potentially Nil Value
```go
var value *int = nil
opt := optional.OfNullable(value) // Creates an empty Optional if value is nil.
fmt.Println(opt)                  // Output: Optional.empty
```

#### Empty Optional
```go
empty := optional.Empty[int]()
fmt.Println(empty) // Output: Optional.empty
```

---

### Check Presence of a Value

#### Check if Value is Present
```go
if opt.IsPresent() {
    fmt.Println("Value is present")
}
```

#### Check if Optional is Empty
```go
if empty.IsEmpty() {
    fmt.Println("No value is present")
}
```

---

### Access the Value

#### Get the Value
```go
value := opt.Get() // Panics if the Optional is empty.
fmt.Println(value) // Output: 42
```

#### Provide a Default Value
```go
defaultValue := empty.OrElse(100) // Returns 100 if Optional is empty.
fmt.Println(defaultValue)         // Output: 100
```

#### Compute a Default Value
```go
computedValue := empty.OrElseGet(func() int {
    return 200
})
fmt.Println(computedValue) // Output: 200
```

#### Throw a Custom Error if Empty
```go
value := empty.OrElseThrow(errors.New("Value is required")) // Panics with the custom error.
```

---

### Perform Conditional Actions

#### If Value is Present
```go
opt.IfPresent(func(val int) {
    fmt.Printf("Value: %d\n", val) // Output: Value: 42
})
```

#### If Present or Else
```go
opt.IfPresentOrElse(
    func(val int) {
        fmt.Printf("Value: %d\n", val)
    },
    func() {
        fmt.Println("No value present")
    },
)
```

---

### Transform the Value

#### Map to a New Value
```go
opt := Of(42)
mapped := Map(opt, func(val int) string {
    return fmt.Sprintf("Value: %d", val)
})
fmt.Println(mapped) // Output: Optional[Value: 42]
```

#### FlatMap to Another Optional
```go
opt := Of(42)
flatMapped := FlatMap(opt, func(val int) Optional[string] {
    if val > 10 {
        return Of(fmt.Sprintf("Value is %d", val))
    }
    return Empty[string]()
})
fmt.Println(flatMapped) // Output: Optional[Value is 42]
```

---

### Examples

#### Example: Safe Division
```go
func safeDivide(a, b int) optional.Optional[float64] {
    if b == 0 {
        return optional.Empty[float64]()
    }
    return optional.Of(float64(a) / float64(b))
}

result := safeDivide(10, 2)
result.IfPresent(func(val float64) {
    fmt.Printf("Result: %.2f\n", val) // Output: Result: 5.00
})

noResult := safeDivide(10, 0)
noResult.IfPresentOrElse(
    func(val float64) {
        fmt.Printf("Result: %.2f\n", val)
    },
    func() {
        fmt.Println("Cannot divide by zero") // Output: Cannot divide by zero
    },
)
```

#### Example: Default Fallback
```go
opt := optional.Empty[string]()
value := opt.OrElse("default value")
fmt.Println(value) // Output: default value
```

---

## API Reference

### Creation

- `Empty[T]()` - Returns an empty `Optional`.
- `Of[T](value T)` - Returns an `Optional` with the given value, panics if the value is `nil`.
- `OfNullable[T](value *T)` - Returns an `Optional` with the given value or an empty `Optional` if `nil`.

### Inspection

- `IsPresent() bool` - Returns `true` if a value is present.
- `IsEmpty() bool` - Returns `true` if no value is present.

### Access

- `Get() T` - Returns the value if present, panics if empty.
- `OrElse(other T) T` - Returns the value if present, otherwise `other`.
- `OrElseGet(supplier func() T) T` - Returns the value if present, otherwise computes it using the supplier.
- `OrElseThrow(err error) T` - Returns the value if present, otherwise panics with the provided error.

### Actions

- `IfPresent(action func(T))` - Executes the action if a value is present.
- `IfPresentOrElse(action func(T), emptyAction func())` - Executes `action` if a value is present, otherwise executes `emptyAction`.

### Transformation

- `Map(mapper func(T) U) Optional[U]` - Applies the mapping function to the value and returns a new `Optional`.
- `FlatMap(mapper func(T) Optional[U]) Optional[U]` - Applies the mapping function and returns the resulting `Optional` directly.

---

## Contributing

Contributions are welcome! Please open an issue or submit a pull request for improvements or new features.

---

## License

This project is licensed under the MIT License. See the [LICENSE](LICENSE) file for details.
```

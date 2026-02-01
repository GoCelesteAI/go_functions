# Go Functions

Code examples from **Go Tutorial for Beginners #11 - Functions**

## Contents

- `func_basics.go` - Function declaration and parameters
- `func_returns.go` - Return values and multiple returns
- `func_named.go` - Named return values
- `func_files/` - Functions organized in separate files
  - `main.go` - Main entry point
  - `utils.go` - Utility functions

## Running the Examples

```bash
# Single file examples
go run func_basics.go
go run func_returns.go
go run func_named.go

# Multiple file example
cd func_files
go run *.go
```

## Key Concepts

### Function with Parameters
```go
func greet(name string) {
    fmt.Println("Hello,", name)
}

func greetFull(firstName, lastName string) {
    fmt.Printf("Hello, %s %s!\n", firstName, lastName)
}
```

### Return Values
```go
func add(a, b int) int {
    return a + b
}
```

### Multiple Return Values
```go
func divmod(a, b int) (int, int) {
    return a / b, a % b
}

// Common pattern: value + error
func safeDivide(a, b int) (int, error) {
    if b == 0 {
        return 0, fmt.Errorf("cannot divide by zero")
    }
    return a / b, nil
}
```

### Named Return Values
```go
func rectangle(w, h float64) (area, perimeter float64) {
    area = w * h
    perimeter = 2 * (w + h)
    return // naked return
}
```

### Functions in Separate Files
```go
// utils.go
package main

func Add(a, b int) int {
    return a + b
}

// main.go
package main

func main() {
    sum := Add(10, 20) // Uses function from utils.go
}
```

## Watch the Tutorial

[Go Tutorial for Beginners #11 - Functions](https://youtube.com/@CelesteAI)

## License

MIT

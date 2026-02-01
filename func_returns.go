package main

import "fmt"

// Function with single return value
func add(a, b int) int {
  return a + b
}

// Function with multiple return values
func divmod(a, b int) (int, int) {
  quotient := a / b
  remainder := a % b
  return quotient, remainder
}

// Common pattern: return value and error
func safeDivide(a, b int) (int, error) {
  if b == 0 {
    return 0, fmt.Errorf("cannot divide by zero")
  }
  return a / b, nil
}

func main() {
  // Single return value
  sum := add(10, 20)
  fmt.Println("10 + 20 =", sum)

  // Multiple return values
  q, r := divmod(17, 5)
  fmt.Printf("17 / 5 = %d remainder %d\n", q, r)

  // Ignore one return value with underscore
  quotient, _ := divmod(20, 3)
  fmt.Println("20 / 3 =", quotient)

  // Error handling pattern
  result, err := safeDivide(10, 2)
  if err != nil {
    fmt.Println("Error:", err)
  } else {
    fmt.Println("10 / 2 =", result)
  }

  // Handle error case
  result, err = safeDivide(10, 0)
  if err != nil {
    fmt.Println("Error:", err)
  } else {
    fmt.Println("Result:", result)
  }
}

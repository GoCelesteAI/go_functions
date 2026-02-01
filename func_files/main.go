package main

import "fmt"

func main() {
  // Using functions from utils.go
  // Both files share the same package

  fmt.Println("=== Math Operations ===")
  sum := Add(10, 20)
  fmt.Println("Add(10, 20) =", sum)

  diff := Subtract(50, 30)
  fmt.Println("Subtract(50, 30) =", diff)

  product := Multiply(7, 8)
  fmt.Println("Multiply(7, 8) =", product)

  fmt.Println("\n=== String Operations ===")
  greeting := FormatGreeting("Alice")
  fmt.Println(greeting)

  upper := ToUpperCase("hello world")
  fmt.Println("ToUpperCase:", upper)

  fmt.Println("\n=== Using Helper ===")
  result := CalculateTotal(100, 0.15)
  fmt.Printf("Total with 15%% tax: $%.2f\n", result)
}

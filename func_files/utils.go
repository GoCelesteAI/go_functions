package main

import "strings"

// Math operations
// Functions in the same package can call each other freely

func Add(a, b int) int {
	return a + b
}

func Subtract(a, b int) int {
	return a - b
}

func Multiply(a, b int) int {
	return a * b
}

// String operations

func FormatGreeting(name string) string {
	return "Hello, " + name + "! Welcome!"
}

func ToUpperCase(s string) string {
	return strings.ToUpper(s)
}

// Business logic

func CalculateTotal(price, taxRate float64) float64 {
	tax := price * taxRate
	return price + tax
}

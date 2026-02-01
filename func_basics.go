package main

import "fmt"

// Basic function with no parameters
func sayHello() {
  fmt.Println("Hello, World!")
}

// Function with one parameter
func greet(name string) {
  fmt.Println("Hello,", name)
}

// Function with multiple parameters (same type)
func greetFull(firstName, lastName string) {
  fmt.Printf("Hello, %s %s!\n", firstName, lastName)
}

// Function with different parameter types
func printInfo(name string, age int) {
  fmt.Printf("%s is %d years old\n", name, age)
}

func main() {
  // Call function with no parameters
  sayHello()

  // Call function with parameter
  greet("Alice")
  greet("Bob")

  // Call function with multiple parameters
  greetFull("John", "Doe")

  // Call function with different types
  printInfo("Carol", 25)
}

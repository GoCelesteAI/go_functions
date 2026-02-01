package main

import "fmt"

// Named return values
func rectangle(width, height float64) (area, perimeter float64) {
  area = width * height
  perimeter = 2 * (width + height)
  return // naked return
}

// Named returns with explicit return
func circle(radius float64) (area, circumference float64) {
  area = 3.14159 * radius * radius
  circumference = 2 * 3.14159 * radius
  return area, circumference
}

// Named returns as documentation
func minMax(numbers []int) (min, max int) {
  if len(numbers) == 0 {
    return 0, 0
  }

  min = numbers[0]
  max = numbers[0]

  for _, n := range numbers {
    if n < min {
      min = n
    }
    if n > max {
      max = n
    }
  }
  return
}

func main() {
  // Named returns
  a, p := rectangle(10, 5)
  fmt.Printf("Rectangle: area=%.2f, perimeter=%.2f\n", a, p)

  // Named returns with explicit return
  circArea, circ := circle(5)
  fmt.Printf("Circle: area=%.2f, circumference=%.2f\n", circArea, circ)

  // Named returns as documentation
  numbers := []int{3, 1, 4, 1, 5, 9, 2, 6}
  minimum, maximum := minMax(numbers)
  fmt.Printf("Numbers: %v\n", numbers)
  fmt.Printf("Min: %d, Max: %d\n", minimum, maximum)
}

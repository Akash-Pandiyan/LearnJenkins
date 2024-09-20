package main

import (
	"fmt"
)

func main() {
	// Print a simple message
	fmt.Println("Hello, Go!")
	
	// Add two numbers
	a := 5
	b := 3
	sum := a + b
	fmt.Printf("Sum of %d and %d is: %d\n", a, b, sum)

	// Call a function that returns the square of a number
	num := 4
	fmt.Printf("Square of %d is: %d\n", num, square(num))
}

// Function to return the square of a number
func square(n int) int {
	return n * n
}

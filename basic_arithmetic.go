package main

import (
	"fmt"
)

// Main function that demonstrates basic arithmetic operations
func main() {
	fmt.Println("Hello, World!")
	fmt.Println("The addition of 5 and 3 is:", add(5, 3))
	fmt.Println("The subtraction of 5 and 3 is:", subtract(5, 3))
	fmt.Println("The maximum of 5 and 3 is:", max(5, 3))
	fmt.Println("Is 7 a prime number?", isPrime(7))
	fmt.Println("Is 10 an even number?", isEven(10))
}

// Function to add two integers
func add(a int, b int) int {
	return a + b
}

// Function to subtract two integers
func subtract(a int, b int) int {
	return a - b
}

// Function to find the maximum of two integers
func max(a int, b int) int {
	if a > b {
		return a
	}
	return b
}

func isPrime(n int) bool {
	if n <= 1 {
		return false
	}
	for i := 2; i*i <= n; i++ {
		if n%i == 0 {
			return false
		}
	}
	return true
}

func isEven(n int) bool {
	return n%2 == 0
}

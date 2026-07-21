package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func Add(a int, b int) int {
	return a + b
}

func Max(a int, b int) int {
	if a > b {
		return a
	}
	return b
}

func IsPrime(n int) bool {
	if n < 2 {
		return false
	}
	for i := 2; i*i <= n; i++ {
		if n%i == 0 {
			return false
		}
	}
	return true
}

func ReverseString(s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}

// readInt prompts, reads a line, and parses it as an int.
func readInt(reader *bufio.Reader, prompt string) int {
	fmt.Print(prompt)
	line, _ := reader.ReadString('\n')
	n, _ := strconv.Atoi(strings.TrimSpace(line))
	return n
}

func main() {
	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Println("\n--- Menu ---")
		fmt.Println("1. Add two numbers")
		fmt.Println("2. Find the max of two numbers")
		fmt.Println("3. Check if a number is prime")
		fmt.Println("4. Reverse a string")
		fmt.Println("Type 'exit' to quit")
		fmt.Print("Choose an option: ")

		choice, _ := reader.ReadString('\n')
		choice = strings.TrimSpace(choice)

		switch choice {
		case "exit":
			fmt.Println("Goodbye!")
			return
		case "1":
			a := readInt(reader, "Enter first number: ")
			b := readInt(reader, "Enter second number: ")
			fmt.Println("Result:", Add(a, b))
		case "2":
			a := readInt(reader, "Enter first number: ")
			b := readInt(reader, "Enter second number: ")
			fmt.Println("Result:", Max(a, b))
		case "3":
			n := readInt(reader, "Enter a number: ")
			fmt.Println("Is prime:", IsPrime(n))
		case "4":
			fmt.Print("Enter a string: ")
			s, _ := reader.ReadString('\n')
			fmt.Println("Reversed:", ReverseString(strings.TrimSpace(s)))
		default:
			fmt.Println("Not a valid option, try again.")
		}
	}
}

package main

import (
    "fmt"
    "github.com/google/uuid"
)

// Defines a function that takes two integers and returns their sum
func add(x int, y int) int {
    return x + y
}

func main() {
    result := add(5, 7)
    fmt.Println("5 + 7 =", result)
	// Generate and print a UUID
    fmt.Println("Your unique ID is:", uuid.New().String())
}

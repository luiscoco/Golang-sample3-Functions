# Golang: Functions

![image](https://github.com/luiscoco/Golang-sample3-Functions/assets/32194879/9e22c3c7-fb56-42f3-818b-61a07372e16a)

## 1. Source code

```go
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
```

## 2. How to run the application

```
go run functions.go
```

package main

import (
	"errors" // for errors.New
	"fmt"
	"math" // for math.Abs
)

const (
	delta = 0.0001 // Delta value to stop iteration in Newton's method.
)

// TODO: Implement the sqrt function.

func main() {
	value, _ := sqrt(400.0)
	fmt.Printf("The square root of %f is %f\n", 400.0, value)
}

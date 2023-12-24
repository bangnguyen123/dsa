package main

import (
	"fmt"
)

// gets the power series of integer a and returns tuple of square of a
// and cube of a
func powerSeries(a int) (int, int, error) {
	square := a * a
	cube := square * a
	return square, cube, nil
}

func main() {
	var square int
	var cube int
	square, cube, _ = powerSeries(3)
	fmt.Println("Square ", square, "Cube", cube)
}

package main

import "fmt"

// 1. Basic function
func add(x int, y int) int {
	return x + y
}

// 2. Multiple return values
func swap(x, y string) (string, string) {
	return y, x
}

// 3. Named return values
// The return variables are treated as variables defined at the top of the function.
func split(sum int) (x, y int) {
	x = sum * 4 / 9
	y = sum - x
	return // "Naked" return
}

// 4. Variadic functions
// Accepts any number of trailing arguments.
func sum(nums ...int) int {
	total := 0
	for _, num := range nums {
		total += num
	}
	return total
}

func main() {
	// Basic
	fmt.Println("Add 42 + 13:", add(42, 13))

	// Multiple returns
	a, b := swap("hello", "world")
	fmt.Println("Swapped:", a, b)

	// Named returns
	x, y := split(17)
	fmt.Println("Split 17:", x, y)

	// Variadic
	fmt.Println("Sum of 1, 2, 3:", sum(1, 2, 3))
	fmt.Println("Sum of 10, 20, 30, 40:", sum(10, 20, 30, 40))
}

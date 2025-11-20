package main

import "fmt"

func main() {
	i := 42

	// 1. Creating a pointer
	// & generates a pointer to its operand.
	p := &i
	fmt.Println("Value of i:", i)
	fmt.Println("Address of i (p):", p)

	// 2. Dereferencing
	// * reads the value at the pointer's address.
	fmt.Println("Value at pointer p (*p):", *p)

	// 3. Changing value via pointer
	*p = 21 // This changes 'i' because 'p' points to 'i'
	fmt.Println("New value of i:", i)

	// 4. Pointers in functions
	num := 10
	fmt.Println("\nBefore double:", num)
	double(num) // Passes a copy
	fmt.Println("After double (no change):", num)

	doubleWithPointer(&num) // Passes the address
	fmt.Println("After doubleWithPointer:", num)
}

// Receives a copy of the value
func double(n int) {
	n = n * 2
}

// Receives a pointer to the original value
func doubleWithPointer(n *int) {
	*n = *n * 2
}

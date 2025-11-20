package main

import "fmt"

func main() {
	// 1. Arrays
	// Fixed size. Cannot be resized.
	var arr [3]int
	arr[0] = 10
	arr[1] = 20
	arr[2] = 30
	fmt.Println("Array:", arr)

	// 2. Slices
	// Dynamic size. Built on top of arrays.
	// This is what you will use 99% of the time.
	slice := []string{"Apple", "Banana", "Cherry"}
	fmt.Println("Slice:", slice)

	// Appending to a slice
	slice = append(slice, "Date")
	fmt.Println("Slice after append:", slice)

	// Slicing a slice (getting a subset)
	// [inclusive:exclusive]
	subset := slice[1:3] // "Banana", "Cherry"
	fmt.Println("Subset [1:3]:", subset)

	// 3. Maps
	// Key-value pairs.
	ages := make(map[string]int)
	ages["Alice"] = 25
	ages["Bob"] = 30

	fmt.Println("Map:", ages)
	fmt.Println("Alice's age:", ages["Alice"])

	// Deleting from a map
	delete(ages, "Bob")
	fmt.Println("Map after delete:", ages)

	// 4. Range Loop
	// Iterating over a slice
	fmt.Println("\nIterating over slice:")
	for index, value := range slice {
		fmt.Printf("%d: %s\n", index, value)
	}

	// Iterating over a map
	fmt.Println("\nIterating over map:")
	for key, value := range ages {
		fmt.Printf("%s is %d years old\n", key, value)
	}
}

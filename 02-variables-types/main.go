package main

import "fmt"

func main() {
	// 1. Declaring variables with var
	var name string = "Gopher"
	var age int = 10
	var isCool bool = true

	fmt.Println("Name:", name)
	fmt.Println("Age:", age)
	fmt.Println("Is Cool:", isCool)

	// 2. Short variable declaration (Type inference)
	// This is the most common way to declare variables inside functions.
	country := "Canada"
	year := 2023

	fmt.Printf("Country: %s, Year: %d\n", country, year)

	// 3. Constants
	// Constants cannot be changed once defined.
	const pi = 3.14159
	fmt.Println("Pi:", pi)

	// 4. Zero values
	// Variables declared without a value get a "zero value".
	var emptyString string
	var emptyInt int
	fmt.Printf("Zero values -> String: '%s', Int: %d\n", emptyString, emptyInt)
}

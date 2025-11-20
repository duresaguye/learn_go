package main

import "fmt"

func main() {
	// 1. If / Else
	age := 18
	if age >= 18 {
		fmt.Println("You are an adult.")
	} else if age >= 13 {
		fmt.Println("You are a teenager.")
	} else {
		fmt.Println("You are a child.")
	}

	// 2. For Loops
	// Go only has 'for'. No 'while' or 'do-while'.
	
	// Standard loop
	fmt.Println("Counting to 5:")
	for i := 1; i <= 5; i++ {
		fmt.Print(i, " ")
	}
	fmt.Println()

	// While-style loop
	fmt.Println("Countdown:")
	count := 3
	for count > 0 {
		fmt.Println(count)
		count--
	}

	// 3. Switch Statements
	day := "Saturday"
	switch day {
	case "Monday", "Tuesday", "Wednesday", "Thursday", "Friday":
		fmt.Println("It's a weekday.")
	case "Saturday", "Sunday":
		fmt.Println("It's the weekend!")
	default:
		fmt.Println("Invalid day.")
	}
}

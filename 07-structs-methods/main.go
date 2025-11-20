package main

import "fmt"

// 1. Defining a Struct
type Person struct {
	FirstName string
	LastName  string
	Age       int
}

// 2. Methods
// A method is a function with a special receiver argument.

// Value Receiver: Receives a copy. Cannot modify the original struct.
func (p Person) FullName() string {
	return p.FirstName + " " + p.LastName
}

// Pointer Receiver: Receives a pointer. Can modify the original struct.
// Use this when you need to modify the struct or if the struct is large.
func (p *Person) HaveBirthday() {
	p.Age++
}

// 3. Embedded Structs (Composition)
type Employee struct {
	Person   // Employee "is a" Person (kind of)
	JobTitle string
	Salary   int
}

func main() {
	// Creating a struct
	p1 := Person{
		FirstName: "John",
		LastName:  "Doe",
		Age:       30,
	}

	fmt.Println("Person:", p1)
	fmt.Println("Full Name:", p1.FullName())

	// Calling a pointer receiver method
	// Go automatically interprets p1.HaveBirthday() as (&p1).HaveBirthday()
	p1.HaveBirthday()
	fmt.Println("New Age:", p1.Age)

	// Using embedded structs
	e1 := Employee{
		Person: Person{
			FirstName: "Jane",
			LastName:  "Smith",
			Age:       28,
		},
		JobTitle: "Engineer",
		Salary:   100000,
	}

	// We can access Person fields directly on Employee
	fmt.Println("\nEmployee:", e1.FullName())
	fmt.Println("Job:", e1.JobTitle)
}

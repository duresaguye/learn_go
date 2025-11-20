# Lesson 8: Interfaces

Defining behavior, not data.

## Key Concepts

1.  **Interface**: A set of method signatures.
    *   `type Name interface { Method() ReturnType }`
2.  **Implicit Implementation**: There is no `implements` keyword in Go. If a type has all the methods declared in an interface, it *automatically* satisfies that interface. This is called "Duck Typing" (if it quacks like a duck...).
3.  **Polymorphism**: You can write functions that accept an interface, allowing them to work with many different types.

## How to Run

```bash
go run main.go
```

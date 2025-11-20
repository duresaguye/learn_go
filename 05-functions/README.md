# Lesson 5: Functions

Building blocks of your program.

## Key Concepts

1.  **Declaration**: `func name(params) returnType { ... }`
2.  **Multiple Returns**: Go functions can return multiple values. This is often used to return a result and an error (more on errors later).
    *   `func foo() (int, string)`
3.  **Named Returns**: You can name your return values in the function signature. A "naked" `return` statement returns those named values.
4.  **Variadic Functions**: Functions that accept a variable number of arguments (like `fmt.Println`). Use `...type` as the last parameter.

## How to Run

```bash
go run main.go
```

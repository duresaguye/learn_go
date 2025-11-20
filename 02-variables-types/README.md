# Lesson 2: Variables and Types

In this lesson, we explore how to store data in Go.

## Key Concepts

1.  **`var` keyword**: The standard way to declare variables. You can specify the type explicitly (e.g., `var x int`).
2.  **Short Declaration `:=`**: Inside functions, you can use `:=` to declare and initialize a variable in one step. Go infers the type for you.
    *   *Note*: You cannot use `:=` outside of functions.
3.  **Basic Types**:
    *   `string`: Text data.
    *   `int`: Whole numbers.
    *   `bool`: True or false.
    *   `float64`: Decimal numbers.
4.  **Constants**: Declared with `const`. Their value cannot be changed.
5.  **Zero Values**: If you declare a variable but don't assign a value, Go gives it a default "zero" value (e.g., `0` for numbers, `""` for strings, `false` for bools).

## How to Run

```bash
go run main.go
```

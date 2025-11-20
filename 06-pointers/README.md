# Lesson 6: Pointers

Understanding memory addresses and references.

## Key Concepts

1.  **What is a Pointer?**: It's just a variable that stores the *memory address* of another variable, instead of the value itself.
2.  **`&` Operator**: The "address of" operator. `&x` gives you the memory address of `x`.
3.  **`*` Operator**: The "dereference" operator. `*p` gives you the value stored at the address `p`.
4.  **Why use them?**
    *   **Efficiency**: Passing a large struct to a function makes a copy. Passing a pointer just copies the small memory address.
    *   **Mutability**: If you want a function to change the original variable, you must pass a pointer.

## How to Run

```bash
go run main.go
```

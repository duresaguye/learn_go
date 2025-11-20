# Lesson 7: Structs and Methods

Creating your own data types.

## Key Concepts

1.  **Structs**: A collection of fields. It's how you define custom data structures.
    *   `type Name struct { Field Type }`
2.  **Methods**: Functions attached to a specific type.
    *   `func (receiver Type) MethodName() { ... }`
3.  **Receivers**:
    *   **Value Receiver** (`p Person`): Gets a copy. Safe, but cannot change the original.
    *   **Pointer Receiver** (`p *Person`): Gets a reference. Can change the original. Efficient for large structs.
4.  **Embedding**: Go doesn't have inheritance (subclasses). Instead, it has composition via embedding. You can include one struct inside another.

## How to Run

```bash
go run main.go
```

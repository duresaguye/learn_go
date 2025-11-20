# Project: Library Management System

A system to manage books using Structs and Methods.

## How it Works

We define a `Book` struct and a `Library` struct. The `Library` struct has methods attached to it (like `AddBook`, `BorrowBook`) to manage the collection.

## Commands

*   `add <Title> by <Author>`: Add a new book.
    *   Example: `add Harry Potter by J.K. Rowling`
*   `borrow <ID>`: Borrow a book by its ID.
*   `return <ID>`: Return a book by its ID.
*   `list`: Show all books and their status.
*   `quit`: Exit.

## Concepts Used

*   **Structs**: Modeling real-world objects (`Book`, `Library`).
*   **Pointer Receivers**: Methods that modify the `Library` state (`*Library`).
*   **Slices of Structs**: Storing multiple books.

## How to Run

```bash
go run main.go
```

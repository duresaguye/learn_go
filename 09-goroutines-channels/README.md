# Lesson 9: Goroutines and Channels

Doing multiple things at once.

## Key Concepts

1.  **Goroutines**: Functions that run concurrently with other functions.
    *   Just add `go` before a function call: `go myFunction()`.
    *   They are extremely lightweight (you can run thousands of them).
2.  **Channels**: The pipes that connect goroutines.
    *   `ch := make(chan int)`
    *   **Send**: `ch <- v`
    *   **Receive**: `v := <-ch`
    *   By default, sends and receives block until the other side is ready. This allows goroutines to synchronize without explicit locks.
3.  **Buffered Channels**: `make(chan int, 100)`. Sends only block when the buffer is full. Receives only block when the buffer is empty.

## How to Run

```bash
go run main.go
```

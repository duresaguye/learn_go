# Lesson 4: Arrays, Slices, and Maps

Organizing data into collections.

## Key Concepts

1.  **Arrays**: Fixed size. `[5]int`.
    *   Rarely used directly in Go because they are inflexible.
2.  **Slices**: Dynamic size. `[]int`.
    *   The standard way to work with lists of data.
    *   Use `append(slice, value)` to add items.
    *   Use `slice[start:end]` to get a sub-section.
3.  **Maps**: Key-Value pairs. `map[string]int`.
    *   Like dictionaries in Python or objects in JavaScript.
    *   Use `delete(map, key)` to remove items.
4.  **Range**: The keyword used to loop over slices and maps.
    *   `for index, value := range slice { ... }`
    *   `for key, value := range map { ... }`

## How to Run

```bash
go run main.go
```

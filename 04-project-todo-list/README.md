# Project: Simple Todo List CLI

Manage your tasks with a command-line interface.

## How it Works

This program runs in a loop, waiting for your commands. It stores tasks in a Go **slice**.

## Commands

*   `add <task>`: Adds a new task to the list.
    *   Example: `add Buy milk`
*   `list`: Shows all current tasks.
*   `quit`: Exits the program.

## Concepts Used

*   **Slices**: Storing the list of tasks (`[]string`).
*   **Infinite Loop**: `for { ... }` to keep the program running.
*   **String Manipulation**: `strings.SplitN` to separate the command from the task description.

## How to Run

```bash
go run main.go
```

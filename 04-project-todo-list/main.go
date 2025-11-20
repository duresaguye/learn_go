package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	var todos []string
	reader := bufio.NewReader(os.Stdin)

	fmt.Println("Simple Todo List")
	fmt.Println("Commands: add [task], list, quit")

	for {
		fmt.Print("> ")
		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input)

		parts := strings.SplitN(input, " ", 2)
		command := parts[0]

		switch command {
		case "add":
			if len(parts) < 2 {
				fmt.Println("Error: Please specify a task to add.")
				continue
			}
			task := parts[1]
			todos = append(todos, task)
			fmt.Println("Added task:", task)

		case "list":
			if len(todos) == 0 {
				fmt.Println("Your todo list is empty.")
			} else {
				fmt.Println("Your Todos:")
				for i, task := range todos {
					fmt.Printf("%d. %s\n", i+1, task)
				}
			}

		case "quit":
			fmt.Println("Goodbye!")
			return

		default:
			fmt.Println("Unknown command. Try 'add', 'list', or 'quit'.")
		}
	}
}

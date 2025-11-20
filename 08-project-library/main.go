package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// Book struct
type Book struct {
	ID         int
	Title      string
	Author     string
	IsBorrowed bool
}

// Library struct
type Library struct {
	Books []Book
}

// AddBook adds a new book to the library
func (l *Library) AddBook(title, author string) {
	id := len(l.Books) + 1
	newBook := Book{
		ID:         id,
		Title:      title,
		Author:     author,
		IsBorrowed: false,
	}
	l.Books = append(l.Books, newBook)
	fmt.Printf("Added: %s by %s (ID: %d)\n", title, author, id)
}

// BorrowBook marks a book as borrowed
func (l *Library) BorrowBook(id int) {
	for i := range l.Books {
		if l.Books[i].ID == id {
			if l.Books[i].IsBorrowed {
				fmt.Println("Sorry, that book is already borrowed.")
				return
			}
			l.Books[i].IsBorrowed = true
			fmt.Println("You borrowed:", l.Books[i].Title)
			return
		}
	}
	fmt.Println("Book not found.")
}

// ReturnBook marks a book as returned
func (l *Library) ReturnBook(id int) {
	for i := range l.Books {
		if l.Books[i].ID == id {
			if !l.Books[i].IsBorrowed {
				fmt.Println("This book was not borrowed.")
				return
			}
			l.Books[i].IsBorrowed = false
			fmt.Println("You returned:", l.Books[i].Title)
			return
		}
	}
	fmt.Println("Book not found.")
}

// ListBooks prints all books
func (l *Library) ListBooks() {
	fmt.Println("\n--- Library Catalog ---")
	if len(l.Books) == 0 {
		fmt.Println("No books in library.")
		return
	}
	for _, b := range l.Books {
		status := "Available"
		if b.IsBorrowed {
			status = "Borrowed"
		}
		fmt.Printf("[%d] %s by %s (%s)\n", b.ID, b.Title, b.Author, status)
	}
	fmt.Println("-----------------------")
}

func main() {
	library := Library{}
	reader := bufio.NewReader(os.Stdin)

	// Pre-populate some books
	library.AddBook("The Go Programming Language", "Alan A. A. Donovan")
	library.AddBook("Clean Code", "Robert C. Martin")

	fmt.Println("\nWelcome to the Library System!")
	fmt.Println("Commands: add, borrow, return, list, quit")

	for {
		fmt.Print("\n> ")
		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input)
		parts := strings.SplitN(input, " ", 2)
		command := parts[0]

		switch command {
		case "add":
			if len(parts) < 2 {
				fmt.Println("Usage: add <Title> by <Author>")
				continue
			}
			// Simple parsing: split by " by "
			details := strings.Split(parts[1], " by ")
			if len(details) != 2 {
				fmt.Println("Usage: add <Title> by <Author>")
				continue
			}
			library.AddBook(details[0], details[1])

		case "borrow":
			if len(parts) < 2 {
				fmt.Println("Usage: borrow <ID>")
				continue
			}
			id, _ := strconv.Atoi(parts[1])
			library.BorrowBook(id)

		case "return":
			if len(parts) < 2 {
				fmt.Println("Usage: return <ID>")
				continue
			}
			id, _ := strconv.Atoi(parts[1])
			library.ReturnBook(id)

		case "list":
			library.ListBooks()

		case "quit":
			fmt.Println("Goodbye!")
			return

		default:
			fmt.Println("Unknown command.")
		}
	}
}

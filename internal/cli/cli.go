package cli

import (
	"fmt"
	"os"
	"strconv"
	"time"

	"github.com/JoaoRafa19/go-intensivo/internal/service"
)

type BookCLI struct {
	service *service.BookService
}

func NewBookCli(service *service.BookService) *BookCLI {
	return &BookCLI{
		service: service,
	}
}

func (cli *BookCLI) Run() {

	if len(os.Args) < 2 {
		fmt.Println("Usage: books <comand> [arguments]")
		return
	}
	command := os.Args[1]

	switch command {
	case "search":
		if len(os.Args) < 3 {
			fmt.Println("Usage: books search <book title>")
			return
		}
		bookName := os.Args[2]
		cli.searchBooks(bookName)
	case "simulate":
		if len(os.Args) < 3 {
			fmt.Println("Usage: books simulate <book_id> <book_id> ...")
			return
		}

		bookIds := os.Args[2:]
		cli.simulateReading(bookIds)

	}
}

func (cli *BookCLI) simulateReading(bookIDs []string) {
	var bookIDS []int
	for _, idString := range bookIDs {
		id, err := strconv.Atoi(idString)
		if err != nil {
			fmt.Println("invalid book id:", idString)
			continue
		}
		bookIDS = append(bookIDS, id)
	}
	responses := cli.service.SimulateMultipleReadings(bookIDS, 1*time.Second)
	for _ , response := range responses {
		fmt.Println(response)
	}
}

func (cli *BookCLI) searchBooks(name string) {
	books, err := cli.service.SearchBooksByName(name)

	if err != nil {
		fmt.Println("Error search")
		return
	}

	if len(books) == 0 {
		fmt.Println("No books")
		return
	}

	fmt.Printf("%d books found with name '%s'\n", len(books), name)
	for _, book := range books {
		fmt.Printf("ID: %d, Title: %s, Author: %s, Genre: %s\n", book.ID, book.Title, book.Author, book.Genre)
	}
}

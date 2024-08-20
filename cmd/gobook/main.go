package main

import (
	"database/sql"
	"net/http"
	"os"

	"github.com/JoaoRafa19/go-intensivo/internal/cli"
	"github.com/JoaoRafa19/go-intensivo/internal/service"
	"github.com/JoaoRafa19/go-intensivo/internal/web"
	_ "github.com/mattn/go-sqlite3"
)

func main() {
	db, err := sql.Open("sqlite3", "./books.db")

	if err != nil {
		panic(err)
	}

	defer db.Close()



	BookService := service.NewBookService(db)
	bookHandlers := web.NewBookHandlers(BookService)

	if len(os.Args) > 1 {
		bookCli := cli.NewBookCli(BookService)
		bookCli.Run()
		return
	}

	router := http.NewServeMux()
	router.HandleFunc("GET /books", bookHandlers.GetBooks)
	router.HandleFunc("GET /books/{id}", bookHandlers.GetBooksByID)
	router.HandleFunc("POST /books", bookHandlers.CreateBook)
	router.HandleFunc("PUT /books/{id}", bookHandlers.UpdateBook)
	router.HandleFunc("DELETE /books/{id}", bookHandlers.DeleteBook)

	http.ListenAndServe(":8080", router)
}
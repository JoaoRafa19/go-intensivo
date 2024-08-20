package service

import (
	"database/sql"
	"fmt"
	"time"
)

type Book struct {
	ID     int
	Title  string
	Author string
	Genre  string
}

type BookService struct {
	db *sql.DB
}

func NewBookService(db *sql.DB) *BookService {
	return &BookService{
		db: db,
	}
}

func (s *BookService) CreateBook(book *Book) error {
	querie := "insert into books (title, author, genre) values (?,?,?)"
	result, err := s.db.Exec(querie, book.Title, book.Author, book.Genre)
	if err != nil {
		return err
	}

	lastInsertID, err := result.LastInsertId()
	if err != nil {
		return err
	}
	book.ID = int(lastInsertID)

	return nil
}

func (s *BookService) GetBooks() ([]Book, error) {
	query := "select * from books"
	rows, err := s.db.Query(query)
	if err != nil {
		return nil, err
	}
	var books []Book

	for rows.Next() {
		var book Book
		err := rows.Scan(&book.ID, &book.Title, &book.Author, &book.Genre)
		if err != nil {
			return nil, err
		}

		books = append(books, book)
	}

	return books, nil
}

func (s *BookService) GetBookByID(id int) (*Book, error) {
	query := "select * from books where id=?"
	row := s.db.QueryRow(query, id)

	var book Book
	err := row.Scan(&book.ID, &book.Title, &book.Author, &book.Genre)
	if err != nil {
		return nil, err
	}

	return &book, nil
}

func (s *BookService) UpdateBook(book *Book) error {
	query := "update books set title=?, author=?, genre=? where id=?"
	_ , err := s.db.Exec(query, book.Title, book.Author, book.Genre, book.ID)
	return err
}

func (s *BookService) DeleteBook(id int) error {
	query := "delete from books where id=?"
	_ , err := s.db.Exec(query, id)
	return err
}



func (s *BookService) SimulateReading(bookID int, duration time.Duration, results chan <- string) {
	book, err := s.GetBookByID(bookID)
	if err != nil ||book == nil {
		results <- fmt.Sprintf("Book %d not found", bookID)
		return
	}
	time.Sleep(duration)

	results <- fmt.Sprintf("book %s read", book.Title)
}

func (s * BookService) SimulateMultipleReadings(bookIDs[]int, duration time.Duration) []string {
	result := make(chan string, len(bookIDs))
	for _, id := range bookIDs {
		go func (bookID int){
			s.SimulateReading(bookID, duration, result)
		}(id)
	}

	var responses[]string 
	for range bookIDs {
		responses = append(responses, <- result) // trava enquanto nao recebe um resultado
	}
	close(result)
	return responses
}

func (s *BookService) SearchBooksByName(name string) ([]Book, error ){
	query := "select id, title, author, genre from books where title like ?"
	rows, err := s.db.Query(query, "%"+name+"%")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var books [] Book
	for rows.Next() {
		var book Book
		err := rows.Scan(&book.ID, &book.Title, &book.Author, &book.Genre)
		if err != nil {
			return nil, err
		}
		books = append(books, book)
	}
	return books, nil
}
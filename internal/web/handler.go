package web

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/JoaoRafa19/go-intensivo/internal/service"
)

type BookHandlers struct {
	service *service.BookService
}

func NewBookHandlers(service *service.BookService) *BookHandlers{
	return &BookHandlers{
		service: service,
	}
}

func (h *BookHandlers) GetBooks(w http.ResponseWriter, r *http.Request) {
	books, err := h.service.GetBooks()

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(`failed to get books`))
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(books)
}

func (h *BookHandlers) CreateBook(w http.ResponseWriter, r *http.Request) {
	var book service.Book
	defer r.Body.Close()
	err := json.NewDecoder(r.Body).Decode(&book)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(`invalid request `))
		return
	}

	err = h.service.CreateBook(&book)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(`failed create`))
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(book)

}

// lida com a requisicao GET /books/{id}
func (h *BookHandlers) GetBooksByID(w http.ResponseWriter, r *http.Request) {
	idStr := r.PathValue("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(`invalid request`))
		return
	}

	book, err := h.service.GetBookByID(id)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(`invalid request`))
		return
	}

	if book == nil {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte(http.StatusText(http.StatusNotFound)))
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(book)
}

func (h *BookHandlers) UpdateBook(w http.ResponseWriter, r *http.Request) {
	idStr := r.PathValue("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(`invalid request`))
		return
	}
	defer r.Body.Close()
	var newBook service.Book
	if err := json.NewDecoder(r.Body).Decode(&newBook); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(`invalid request`))
		return
	}

	newBook.ID = id 

	if err := h.service.UpdateBook(&newBook); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(`invalid request`))
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(newBook)
}

//handle DELETE /books/{id}
func (h *BookHandlers) DeleteBook(w http.ResponseWriter, r *http.Request) {
	idStr := r.PathValue("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(http.StatusText(http.StatusBadRequest)))
		return
	}

	if err := h.service.DeleteBook(id); err!=nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(http.StatusText(http.StatusInternalServerError)))
		return
	}
	w.WriteHeader(http.StatusNoContent)

}

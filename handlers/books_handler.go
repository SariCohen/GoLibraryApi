package handlers

import (
	"encoding/json"
	"net/http"
	"github.com/gorilla/mux"
	"fmt"
	"go-project/models"
)

var Books = make(map[string]models.Book)
var id int

func GetBooks(w http.ResponseWriter, r *http.Request) {
	bookList := make([]models.Book, 0, len(Books))
	for _, book := range Books {
		bookList = append(bookList, book)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(bookList)
}

func GetBookById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	bookID := vars["id"]

	book, exists := Books[bookID]
	if !exists {
		http.Error(w, "Book not found", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(book)
}

func AddBook(w http.ResponseWriter, r *http.Request) {
	var book models.Book

	err := json.NewDecoder(r.Body).Decode(&book)
	if err != nil {
        http.Error(w, "Invalid request body", http.StatusBadRequest)
        return
    }

    if book.Title == "" || book.Author == "" || book.PublishedYear == 0 {
        http.Error(w, "Missing required fields", http.StatusBadRequest)
        return
    }

	id++
	bookID := fmt.Sprintf("%d", id)
	Books[bookID] = book

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]string{"message": "Book created successfully", "id": bookID})
}

func DeleteBook(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	bookID := vars["id"]

	_, exists := Books[bookID]
	if !exists {
		http.Error(w, "Book not found", http.StatusNotFound)
		return
	}

	delete(Books, bookID)

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{"message": "Book deleted successfully"})
}

func SearchBooks(w http.ResponseWriter, r *http.Request) {

	author := r.URL.Query().Get("author")
	fmt.Println(author)

	var results []models.Book

	for _, book := range Books {
		if author == "" || book.Author == author {
			results = append(results, book)
		}
	}

	if len(results) == 0 {
		http.Error(w, "No books found", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	if err := json.NewEncoder(w).Encode(results); err != nil {
		http.Error(w, "Error encoding response", http.StatusInternalServerError)
	}
}

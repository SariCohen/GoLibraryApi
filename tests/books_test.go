package main

import (
	"bytes"
	"encoding/json"
	"go-project/handlers"
	"go-project/models"
	"net/http"
	"net/http/httptest"
	"testing"
	"github.com/gorilla/mux"
)


func TestAddBook(t *testing.T) {
	book := models.Book{
		Title:         "Comix",
		Author:        "Dan Brown",
		PublishedYear: 2024,
	}
	body, _ := json.Marshal(book)

	req, err := http.NewRequest("POST", "/books", bytes.NewBuffer(body))
	if err != nil {
		t.Fatal(err)
	}
	req.Header.Set("Content-Type", "application/json")

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(handlers.AddBook)
	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusCreated {
		t.Errorf("expected status %d but got %d", http.StatusCreated, status)
	}
	
	expected := `"message":"Book created successfully"`
	if !bytes.Contains(rr.Body.Bytes(), []byte(expected)) {
		t.Errorf("expected response to contain %s but got %s", expected, rr.Body.String())
	}
}

func TestGetBooks(t *testing.T) {
	req, err := http.NewRequest("GET", "/books", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(handlers.GetBooks)
	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("expected status %d but got %d", http.StatusOK, status)
	}

	expected := `"title":"Comix"`
	if !bytes.Contains(rr.Body.Bytes(), []byte(expected)) {
		t.Errorf("expected response to contain %s but got %s", expected, rr.Body.String())
	}
}

func TestGetBookById(t *testing.T) {

	r := mux.NewRouter()
	r.HandleFunc("/books/{id:1}", handlers.GetBookById)

	req, err := http.NewRequest("GET", "/books/1", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := r
	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("expected status %d but got %d", http.StatusOK, status)
	}

	expected := `"title":"Comix"`
	if !bytes.Contains(rr.Body.Bytes(), []byte(expected)) {
		t.Errorf("expected response to contain %s but got %s", expected, rr.Body.String())
	}
}


func TestDeleteBook(t *testing.T) {

	r := mux.NewRouter()
	r.HandleFunc("/books/{id:1}", handlers.DeleteBook)

	req, err := http.NewRequest("DELETE", "/books/1", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := r
	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("expected status %d but got %d", http.StatusOK, status)
	}
}


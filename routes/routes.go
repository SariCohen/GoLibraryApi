package routes

import (
	"github.com/gorilla/mux"
	"go-project/handlers"
)



func InitializeRoutes() *mux.Router  {
	router := mux.NewRouter()
	router.HandleFunc("/books", handlers.GetBooks).Methods("GET")
	router.HandleFunc("/books/search", handlers.SearchBooks).Methods("GET")
	router.HandleFunc("/books/{id}", handlers.GetBookById).Methods("GET")
	router.HandleFunc("/books", handlers.AddBook).Methods("POST")
	router.HandleFunc("/books/{id}", handlers.DeleteBook).Methods("DELETE")

	return router;
}

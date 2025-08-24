package main

import (
	"log"
	"net/http"
	"go-project/routes"
)

func main() {
	r := routes.InitializeRoutes()
	log.Println("Server is running on port 8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}

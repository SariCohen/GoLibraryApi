# Go Library API
## 📌 Overview
A simple REST API built with Go and Gorilla Mux to manage a library system. The API allows adding, 
retrieving, and deleting books stored in memory, with basic error handling and optional search functionality.
## ▶️ How to Use This Project
To use and run this project on your local machine, follow the instructions below.
### 📥 Clone the repository
```bash
git clone https://github.com/SariCohen/GoLibraryApi.git
```
```bash
cd GoProject
```
## 🚀 How to Run the Project
### ✅ Prerequisites
Go 1.21+ installed

### 🖥️ Run the Server

```bash
go mod tidy
go run main.go
```

Open the API in your browser or via curl at:

```
http://localhost:8080
```

### 🧪 Run Tests

```bash
go test ./tests -cover
```

## 🧩 Features

- `POST /books` – Add a new book
- `GET /books` – Retrieve all books (with optional filter by author)
- `GET /books/{id}` – Get book by ID
- `DELETE /books/{id}` – Delete a book by ID
- Error handling with proper status codes
- Unit tests with Go’s testing package

## 📁 Project Structure

```
GO-PROJECT/
├── handlers/
│   └── books_handler.go
├── models/
│   └── Book.go
├── routes/
│   └── routes.go
├── tests/
│   └── books_test.go
├── tmp/
├── main.go
├── go.mod
├── go.sum
└── .air.conf
```

## 🛠 Technologies Used

- Go 1.21+
- Gorilla Mux
- In-memory storage
- Go testing framework

## 🧹 Best Practices Followed

- Clean package structure (handlers, models, routes, tests)
- Separation of concerns
- Consistent error handling
- Unit tests for core functionality

## 👩‍💻 Author

Sara Cohen  
26.08.2025

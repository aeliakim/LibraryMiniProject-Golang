package handlers

import (
	"LibraryMiniProject-Golang/models"
	"encoding/json"
	"net/http"
	"os"
)

func GetBooks(w http.ResponseWriter, r *http.Request) {
	filePath, err := os.Open("data/books.json")
	if err != nil {
		http.Error(w, "Failed to open books file", http.StatusInternalServerError)
		return
	}
	defer filePath.Close()

	var books []models.Book
	if err = json.NewDecoder(filePath).Decode(&books); err != nil {
		http.Error(w, "Failed to parse books", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(books)
}

func AddBook(w http.ResponseWriter, r *http.Request) {
	var newBook models.Book
	if err := json.NewDecoder(r.Body).Decode(&newBook); err != nil {
		http.Error(w, "Failed to parse request body", http.StatusBadRequest)
		return
	}

	// Load existing books
	var books []models.Book
	filePath, err := os.Open("data/books.json")
	if err == nil {
		defer filePath.Close()
		json.NewDecoder(filePath).Decode(&books)
	}

	// Assign a new ID to the new book
	newBook.ID = len(books) + 1
	books = append(books, newBook)

	// Save back to JSON file
	filePath, err = os.Create("data/books.json")
	if err != nil {
		http.Error(w, "Failed to create books file", http.StatusInternalServerError)
		return
	}
	defer filePath.Close()

	json.NewEncoder(filePath).Encode(books)
	w.WriteHeader(http.StatusCreated)
}

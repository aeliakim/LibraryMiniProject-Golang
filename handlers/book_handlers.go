package handlers

import (
	"LibraryMiniProject-Golang/models"
	"encoding/json"
	"net/http"
	"os"
	"strconv"

	"github.com/gorilla/mux"
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
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Content-Type", "application/json")

	// Parse JSON request body
	var book models.Book
	if err := json.NewDecoder(r.Body).Decode(&book); err != nil {
		http.Error(w, "Invalid request body: "+err.Error(), http.StatusBadRequest)
		return
	}

	// Validate required fields
	if book.Title == "" || book.Author == "" || book.ISBN == "" || book.Year == 0 {
		http.Error(w, "All fields are required", http.StatusBadRequest)
		return
	}

	// Read existing books
	books := []models.Book{}
	file, err := os.Open("data/books.json")
	if err == nil {
		defer file.Close()
		if err := json.NewDecoder(file).Decode(&books); err != nil {
			http.Error(w, "Failed to read existing books", http.StatusInternalServerError)
			return
		}
	}

	// Assign new ID
	book.ID = len(books) + 1
	books = append(books, book)

	// Write updated books back to file
	file, err = os.Create("data/books.json")
	if err != nil {
		http.Error(w, "Failed to create books file", http.StatusInternalServerError)
		return
	}
	defer file.Close()

	if err := json.NewEncoder(file).Encode(books); err != nil {
		http.Error(w, "Failed to save book", http.StatusInternalServerError)
		return
	}

	// Return success response with the new book
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(book)
}

func DeleteBook(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	idStr := vars["id"]
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid book ID", http.StatusBadRequest)
		return
	}

	// Read existing books
	file, err := os.Open("data/books.json")
	if err != nil {
		http.Error(w, "Failed to open books file", http.StatusInternalServerError)
		return
	}
	defer file.Close()

	var books []models.Book
	if err = json.NewDecoder(file).Decode(&books); err != nil {
		http.Error(w, "Failed to parse books", http.StatusInternalServerError)
		return
	}

	// Find and remove the book
	found := false
	newBooks := []models.Book{}
	for _, book := range books {
		if book.ID != id {
			newBooks = append(newBooks, book)
		} else {
			found = true
		}
	}

	if !found {
		http.Error(w, "Book not found", http.StatusNotFound)
		return
	}

	// Write updated books back to file
	file, err = os.Create("data/books.json")
	if err != nil {
		http.Error(w, "Failed to update books file", http.StatusInternalServerError)
		return
	}
	defer file.Close()

	if err = json.NewEncoder(file).Encode(newBooks); err != nil {
		http.Error(w, "Failed to save updated books", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}

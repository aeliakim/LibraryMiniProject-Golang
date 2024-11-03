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
	var book models.Book
	if err := r.ParseForm(); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	book.Title = r.FormValue("title")
	book.Author = r.FormValue("author")
	book.ISBN = r.FormValue("isbn")
	year, err := strconv.Atoi(r.FormValue("year"))
	if err != nil {
		http.Error(w, "Invalid year", http.StatusBadRequest)
		return
	}
	book.Year = year

	// Read and update the books.json file
	books := []models.Book{}
	file, err := os.Open("data/books.json")
	if err == nil {
		defer file.Close()
		json.NewDecoder(file).Decode(&books)
	}

	book.ID = len(books) + 1
	books = append(books, book)

	file, err = os.Create("data/books.json")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer file.Close()

	json.NewEncoder(file).Encode(books)
	w.WriteHeader(http.StatusCreated)
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

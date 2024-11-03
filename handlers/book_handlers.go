package handlers

import (
	"LibraryMiniProject-Golang/models"
	"encoding/json"
	"net/http"
	"os"
	"strconv"
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
	http.Redirect(w, r, "/", http.StatusSeeOther)
}

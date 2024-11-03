package models

type Book struct {
	ID     int    `json:"id"`     // unique identifier for the book
	Title  string `json:"title"`  // title of the book
	Author string `json:"author"` // author of the book
	Year   int    `json:"year"`   // publication year of the book
	ISBN   string `json:"isbn"`   // ISBN of the book
}

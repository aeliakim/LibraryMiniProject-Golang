# Library Mini Project

A basic web-based library management system built with Go (Golang) that allows users to add, view, and delete books. The application stores book data in a JSON file and provides a simple web interface.

## Features

- Add a new book with details: title, author, ISBN, and publication year
- View a list of all books in a table format
- Delete a book from the list

## How It Works

1. **Frontend**: HTML and CSS provide a simple and user-friendly interface.
2. **Backend**: Go handles API requests for adding, displaying, and deleting books.
3. **Data Storage**: The list of books is stored locally in a JSON file (`books.json`) in the `data` directory.

## Setup and Local Deployment

Follow these steps to set up and deploy the project locally.

### Prerequisites

- **Go** installed on your machine (version 1.16 or higher recommended).
- **Git** (for cloning the repository).

### Installation

1. **Clone the Repository**:

   ````bash
   git clone https://github.com/yourusername/LibraryMiniProject-Golang.git
   cd LibraryMiniProject-Golang```

   ````

2. **Set Up the data file**:

   - Change the name of the data file from `books-example.json` to `books.json`

3. **Run the application**:

   ```bash
   go run main.go
   ```

4. **Access the application**:
   - Open your web browser and navigate to `http://localhost:8080`

## Project Structure

- `main.go`: Entry point of the application, sets up routes and initializes the server.
- `handlers/`: Contains the Go handlers for serving HTML and handling API requests.
- `models/`: Defines the Book struct that represents the book model.
- `templates/`: Contains `index.html` for the frontend interface.
- `static/`: Holds CSS files and other static assets.
- `data/`: Directory where `books.json` stores the library's data.

## Usage Instructions

- Add a Book: Fill in the book details (title, author, ISBN, and publication year) and click "Add Book" to add a new entry to the list.
- View Books: The book list is displayed in a table format on the main page.
- Delete a Book: Click the "Delete" button next to a book to remove it from the list.

## Additional Notes

- The application is designed to run locally and does not persist data across different environments or sessions.
- API Endpoints:
  - GET /api/books: Fetches the list of all books.
  - POST /books: Adds a new book entry.
  - DELETE /books/{id}: Deletes a book by ID.

## Contributing

Contributions are welcome! Feel free to submit a pull request or open an issue if you have suggestions.

## License

This project is licensed under the MIT License.

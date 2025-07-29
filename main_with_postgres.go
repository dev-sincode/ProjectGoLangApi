// main.go
package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"strconv" // Import for string to int conversion

	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq" // PostgreSQL driver
)

// Book represents data about a book.
type Book struct {
	ID     int     `json:"id"`
	Title  string  `json:"title"`
	Author string  `json:"author"`
	Price  float64 `json:"price"`
}

var db *sql.DB // Global variable for database connection

func initDB() {
	// Database connection string. Replace with your actual credentials.
	// Example: "user=postgres password=your_password dbname=gin_api_db sslmode=disable"
	connStr := "user=postgres password=super dbname=postgres sslmode=disable" // IMPORTANT: Replace your_password
	var err error
	db, err = sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}

	// Ping the database to verify connection
	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Successfully connected to PostgreSQL!")

	// Create books table if it doesn't exist
	createTableSQL := `
	CREATE TABLE IF NOT EXISTS books (
		id SERIAL PRIMARY KEY,
		title VARCHAR(255) NOT NULL,
		author VARCHAR(255) NOT NULL,
		price NUMERIC(10, 2) NOT NULL
	);`
	_, err = db.Exec(createTableSQL)
	if err != nil {
		log.Fatalf("Error creating table: %v", err)
	}
	fmt.Println("Books table ensured.")
}

func main() {
	initDB() // Initialize database connection
	defer db.Close() // Close database connection when main function exits

	router := gin.Default()

	// Define API routes
	router.GET("/books", getBooks)
	router.POST("/books", postBooks)
	router.GET("/books/:id", getBookByID)
	router.PUT("/books/:id", updateBookByID)
	router.DELETE("/books/:id", deleteBookByID)

	router.Run("localhost:8080")
}

// getBooks responds with the list of all books from the database as JSON.
func getBooks(c *gin.Context) {
	rows, err := db.Query("SELECT id, title, author, price FROM books")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	defer rows.Close()

	var books []Book
	for rows.Next() {
		var book Book
		if err := rows.Scan(&book.ID, &book.Title, &book.Author, &book.Price); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		books = append(books, book)
	}

	if err := rows.Err(); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.IndentedJSON(http.StatusOK, books)
}

// postBooks adds a book from JSON received in the request body to the database.
func postBooks(c *gin.Context) {
	var newBook Book

	// Call BindJSON to bind the received JSON to newBook.
	if err := c.BindJSON(&newBook); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Insert the new book into the database
	sqlStatement := `INSERT INTO books (title, author, price) VALUES ($1, $2, $3) RETURNING id`
	err := db.QueryRow(sqlStatement, newBook.Title, newBook.Author, newBook.Price).Scan(&newBook.ID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.IndentedJSON(http.StatusCreated, newBook)
}



// getBookByID locates the book whose ID matches the id
// parameter sent by the client, then returns that book as a response from the database.
func getBookByID(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr) // Convert string ID from URL to int
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid book ID"})
		return
	}

	var book Book
	row := db.QueryRow("SELECT id, title, author, price FROM books WHERE id = $1", id)
	err = row.Scan(&book.ID, &book.Title, &book.Author, &book.Price)

	if err == sql.ErrNoRows {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "book not found"})
		return
	} else if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.IndentedJSON(http.StatusOK, book)
}





// updateBookByID updates an existing book by ID in the database.
func updateBookByID(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr) // Convert string ID from URL to int
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid book ID"})
		return
	}

	var updatedBook Book
	if err := c.BindJSON(&updatedBook); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	updatedBook.ID = id // Ensure the ID from the URL is used for the update

	sqlStatement := `UPDATE books SET title = $1, author = $2, price = $3 WHERE id = $4`
	result, err := db.Exec(sqlStatement, updatedBook.Title, updatedBook.Author, updatedBook.Price, updatedBook.ID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	if rowsAffected == 0 {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "book not found"})
		return
	}

	c.IndentedJSON(http.StatusOK, updatedBook)
}







// main.go
package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// Book represents data about a book.
type Book struct {
	ID     string `json:"id"`
	Title  string `json:"title"`
	Author string `json:"author"`
	Price  float64 `json:"price"`
}

// books slice to seed book data.
var books = []Book{
	{ID: "1", Title: "The Lord of the Rings", Author: "J.R.R. Tolkien", Price: 25.99},
	{ID: "2", Title: "Pride and Prejudice", Author: "Jane Austen", Price: 12.50},
	{ID: "3", Title: "1984", Author: "George Orwell", Price: 15.00},
}




// getBooks responds with the list of all books as JSON.
func getBooks(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, books)
}

// postBooks adds a book from JSON received in the request body.
func postBooks(c *gin.Context) {
	var newBook Book

	// Call BindJSON to bind the received JSON to newBook.
	if err := c.BindJSON(&newBook); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Add the new book to the slice.
	books = append(books, newBook)
	c.IndentedJSON(http.StatusCreated, newBook)
}

// getBookByID locates the book whose ID matches the id
// parameter sent by the client, then returns that book as a response.
func getBookByID(c *gin.Context) {
	id := c.Param("id")

	// Loop over the list of books, looking for a book whose ID value matches the parameter.
	for _, a := range books {
		if a.ID == id {
			c.IndentedJSON(http.StatusOK, a)
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "book not found"})
}



// updateBookByID updates an existing book by ID.
func updateBookByID(c *gin.Context) {
	id := c.Param("id")
	var updatedBook Book

	if err := c.BindJSON(&updatedBook); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	for i, book := range books {
		if book.ID == id {
			// Update the book in the slice
			books[i] = updatedBook
			c.IndentedJSON(http.StatusOK, updatedBook)
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "book not found"})
}


// deleteBookByID deletes a book by ID.
func deleteBookByID(c *gin.Context) {
	id := c.Param("id")

	for i, book := range books {
		if book.ID == id {
			// Remove the book from the slice
			books = append(books[:i], books[i+1:]...)
			c.IndentedJSON(http.StatusOK, gin.H{"message": "book deleted successfully"})
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "book not found"})
}






func main() {
	router := gin.Default()
	// We will add routes and handlers here in subsequent steps.
	// Define API routes
	router.GET("/books", getBooks)
	router.POST("/books", postBooks)
	router.GET("/books/:id", getBookByID)
	router.PUT("/books/:id", updateBookByID)
	router.DELETE("/books/:id", deleteBookByID)

	
	router.Run("localhost:8080")
}


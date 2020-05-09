package controllers

import (
	"context"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type Book struct {
	ID        string    `json:"id,omitempty" bson:"id,omitempty"`
	Name      string    `json:"name" bson:"name"`
	Isbn      string    `json:"isbn" bson:"isbn"`
	Author    *Author   `json:"author" bson:"author"`
	CreatedAt time.Time `json:"created_at" bson:"created_at"`
	UpdatedAt time.Time `json:"updated_at" bson:"updated_at"`
}

type Author struct {
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`
}

// DATABASE INSTANCE
var collection *mongo.Collection

func BookCollection(c *mongo.Database) {
	collection = c.Collection("books")
}

func GetBooks(c *gin.Context) {
	books := []Book{}

	ctx, _ := context.WithTimeout(context.Background(), 30*time.Second)

	// bson = binary json, specific of mongodb --> what json doesn't support
	//bson.M{} -> m=map
	//bson.D{} slice
	cursor, err := collection.Find(ctx, bson.M{})
	if err != nil {
		log.Printf("Error while getting Books, %v\n", err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  http.StatusInternalServerError,
			"message": "Something went wrong",
		})
		return
	}
	// Iterate through the returned cursor.
	for cursor.Next(context.TODO()) {
		var book Book
		cursor.Decode(&book)
		books = append(books, book)
	}
	c.JSON(http.StatusOK, gin.H{
		"status":  http.StatusOK,
		"message": "Book List",
		"data":    books,
	})
	return
}

func CreateBook(c *gin.Context) {
	var book Book

	c.BindJSON(&book)

	name := book.Name
	isbn := book.Isbn
	author := book.Author
	id := uuid.New().String()

	b := Book{
		ID:        id,
		Name:      name,
		Isbn:      isbn,
		Author:    author,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	_, err := collection.InsertOne(context.TODO(), b)

	if err != nil {
		log.Printf("Error inserting new book into db, %x\n", err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  http.StatusInternalServerError,
			"message": "Something went wrong",
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"status":  http.StatusCreated,
		"message": "Book created Successfully",
	})
	return
}

func GetBook(c *gin.Context) {
	bookID := c.Param("bookId")

	book := Book{}

	err := collection.FindOne(context.TODO(), bson.M{"id": bookID}).Decode(&book)
	if err != nil {
		log.Printf("Error getting the book :%x\n", err)
		c.JSON(http.StatusNotFound, gin.H{
			"status":  http.StatusNotFound,
			"message": "Book not found",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"status":  http.StatusOK,
		"message": "Book",
		"data":    book,
	})
	return
}

func EditBook(c *gin.Context) {

	bookId := c.Param("bookId")

	var book Book

	c.BindJSON(&book)

	name := book.Name
	isbn := book.Isbn
	author := book.Author

	newData := bson.M{
		"$set": bson.M{
			"name":       name,
			"isbn":       isbn,
			"author":     author,
			"updated_at": time.Now(),
		},
	}
	_, err := collection.UpdateOne(context.TODO(), bson.M{"id": bookId}, newData)

	if err != nil {
		log.Printf("Error, Reason: %v\n", err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  500,
			"message": "Something went wrong",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"status":  200,
		"message": "Book Edited Successfully",
	})
	return
}

func DeleteBook(c *gin.Context) {

	bookId := c.Param("bookId")

	_, err := collection.DeleteOne(context.TODO(), bson.M{"id": bookId})

	if err != nil {
		log.Printf("Error deleting book : %x\n", err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  http.StatusInternalServerError,
			"message": "Something went wrong",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"status":  http.StatusOK,
		"message": "Book deleted successfully",
	})
	return
}

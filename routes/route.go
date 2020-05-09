package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/shraddha0602/golang-mongodb/controllers"
)

func Routes(router *gin.Engine) {
	router.GET("/", welcome)

	router.GET("/books", controllers.GetBooks)
	router.POST("/book", controllers.CreateBook)
	router.GET("/book/:bookId", controllers.GetBook)
	router.PUT("/book/:bookId", controllers.EditBook)
	router.DELETE("/book/:bookId", controllers.DeleteBook)
}

func welcome(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"status":  200,
		"message": "Welcome to API golang-mongodb",
	})
	return
}

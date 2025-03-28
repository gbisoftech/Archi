package routers

import (
	"main/internal/controllers/auth"
	"main/internal/controllers/book"
	"main/internal/middleware"
	"net/http"

	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	router := gin.Default()

	router.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "Hello, World!"})
	})
	router.POST("/login", auth.Login)
	router.POST("/register", auth.Register)

	protected := router.Group("/")
	protected.Use(middleware.JwtAuthMiddleware())

	protected.POST("/books", book.CreateBook)
	protected.GET("/books", book.GetBooks)
	protected.GET("/books/:id", book.GetBook)
	protected.PUT("/books/:id", book.UpdateBook)
	protected.DELETE("/books/:id", book.DeleteBook)
	return router
}

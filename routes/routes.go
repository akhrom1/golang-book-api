package routes

import (
	"golang-book-api/controllers"
	"golang-book-api/middlewares"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes(r *gin.Engine) {
	api := r.Group("/api")
	api.Use(middlewares.BasicAuth())

	// Categories
	api.GET("/categories", controllers.GetCategories)
	api.POST("/categories", controllers.CreateCategory)
	api.GET("/categories/:id", controllers.GetCategoryByID)
	api.DELETE("/categories/:id", controllers.DeleteCategory)
	api.GET("/categories/:id/books", controllers.GetBooksByCategory)

	// Books
	api.GET("/books", controllers.GetBooks)
	api.POST("/books", controllers.CreateBook)
	api.GET("/books/:id", controllers.GetBookByID)
	api.DELETE("/books/:id", controllers.DeleteBook)
}

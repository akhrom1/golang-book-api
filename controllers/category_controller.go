package controllers

import (
	"database/sql"
	"net/http"

	"golang-book-api/config"
	"golang-book-api/models"

	"github.com/gin-gonic/gin"
)

func GetCategories(c *gin.Context) {
	rows, err := config.DB.Query("SELECT id, name, created_at, created_by, modified_at, modified_by FROM categories")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	defer rows.Close()

	var categories []models.Category
	for rows.Next() {
		var cat models.Category
		rows.Scan(&cat.ID, &cat.Name, &cat.CreatedAt, &cat.CreatedBy, &cat.ModifiedAt, &cat.ModifiedBy)
		categories = append(categories, cat)
	}

	c.JSON(http.StatusOK, categories)
}

func CreateCategory(c *gin.Context) {
	var req struct {
		Name string `json:"name"`
	}

	if err := c.ShouldBindJSON(&req); err != nil || req.Name == "" {
		c.JSON(http.StatusBadRequest, gin.H{"message": "name is required"})
		return
	}

	_, err := config.DB.Exec(
		"INSERT INTO categories(name, created_by) VALUES($1, $2)",
		req.Name, "system",
	)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "category created"})
}

func GetCategoryByID(c *gin.Context) {
	id := c.Param("id")

	var cat models.Category
	err := config.DB.QueryRow(
		"SELECT id, name, created_at, created_by, modified_at, modified_by FROM categories WHERE id=$1",
		id,
	).Scan(&cat.ID, &cat.Name, &cat.CreatedAt, &cat.CreatedBy, &cat.ModifiedAt, &cat.ModifiedBy)

	if err == sql.ErrNoRows {
		c.JSON(http.StatusNotFound, gin.H{"message": "category not found"})
		return
	}

	c.JSON(http.StatusOK, cat)
}

func DeleteCategory(c *gin.Context) {
	id := c.Param("id")

	res, err := config.DB.Exec("DELETE FROM categories WHERE id=$1", id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	aff, _ := res.RowsAffected()
	if aff == 0 {
		c.JSON(http.StatusNotFound, gin.H{"message": "category not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "category deleted"})
}

func GetBooksByCategory(c *gin.Context) {
	id := c.Param("id")

	var exists int
	err := config.DB.QueryRow("SELECT id FROM categories WHERE id=$1", id).Scan(&exists)
	if err == sql.ErrNoRows {
		c.JSON(http.StatusNotFound, gin.H{"message": "category not found"})
		return
	}

	rows, err := config.DB.Query("SELECT id, title, description, image_url, release_year, price, total_page, thickness, category_id FROM books WHERE category_id=$1", id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	defer rows.Close()

	var books []models.Book
	for rows.Next() {
		var b models.Book
		rows.Scan(&b.ID, &b.Title, &b.Description, &b.ImageURL, &b.ReleaseYear, &b.Price, &b.TotalPage, &b.Thickness, &b.CategoryID)
		books = append(books, b)
	}

	c.JSON(http.StatusOK, books)
}

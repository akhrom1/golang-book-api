package controllers

import (
	"database/sql"
	"net/http"

	"golang-book-api/config"
	"golang-book-api/models"

	"github.com/gin-gonic/gin"
)

func GetBooks(c *gin.Context) {
	rows, err := config.DB.Query("SELECT id, title, description, image_url, release_year, price, total_page, thickness, category_id FROM books")
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

func CreateBook(c *gin.Context) {
	var req struct {
		Title       string `json:"title"`
		Description string `json:"description"`
		ImageURL    string `json:"image_url"`
		ReleaseYear int    `json:"release_year"`
		Price       int    `json:"price"`
		TotalPage   int    `json:"total_page"`
		CategoryID  int    `json:"category_id"`
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "invalid payload"})
		return
	}

	if req.ReleaseYear < 1980 || req.ReleaseYear > 2024 {
		c.JSON(http.StatusBadRequest, gin.H{"message": "release_year must be between 1980 and 2024"})
		return
	}

	thickness := "tipis"
	if req.TotalPage > 100 {
		thickness = "tebal"
	}

	_, err := config.DB.Exec(`
		INSERT INTO books
		(title, description, image_url, release_year, price, total_page, thickness, category_id, created_by)
		VALUES ($1,$2,$3,$4,$5,$6,$7,$8,$9)
	`,
		req.Title, req.Description, req.ImageURL, req.ReleaseYear,
		req.Price, req.TotalPage, thickness, req.CategoryID, "system",
	)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "book created"})
}

func GetBookByID(c *gin.Context) {
	id := c.Param("id")

	var b models.Book
	err := config.DB.QueryRow(
		"SELECT id, title, description, image_url, release_year, price, total_page, thickness, category_id FROM books WHERE id=$1",
		id,
	).Scan(&b.ID, &b.Title, &b.Description, &b.ImageURL, &b.ReleaseYear, &b.Price, &b.TotalPage, &b.Thickness, &b.CategoryID)

	if err == sql.ErrNoRows {
		c.JSON(http.StatusNotFound, gin.H{"message": "book not found"})
		return
	}

	c.JSON(http.StatusOK, b)
}

func DeleteBook(c *gin.Context) {
	id := c.Param("id")

	res, err := config.DB.Exec("DELETE FROM books WHERE id=$1", id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	aff, _ := res.RowsAffected()
	if aff == 0 {
		c.JSON(http.StatusNotFound, gin.H{"message": "book not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "book deleted"})
}

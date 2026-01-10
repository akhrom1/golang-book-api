package main

import (
	"golang-book-api/config"
	"golang-book-api/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	config.InitDB()
	config.RunMigration(config.DB)

	r := gin.Default()
	routes.RegisterRoutes(r)

	r.Run(":8080")
}

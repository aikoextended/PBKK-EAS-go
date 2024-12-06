package main

import (
	"gin-gorm/models"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	router.LoadHTMLGlob("templates/*")

	router.GET("/anime", func(c *gin.Context) {
		var animes []models.Anime
		// Fetch anime data from the database
		// db.Find(&animes)

		// For demonstration purposes, we'll use dummy data
		animes = []models.Anime{
			{ID: 1, Title: "Naruto", Genre: models.Genre{Name: "Action"}, Review: "Great anime!", Episodes: 220, CreatedAt: time.Now(), UpdatedAt: time.Now()},
			{ID: 2, Title: "One Piece", Genre: models.Genre{Name: "Adventure"}, Review: "Amazing story!", Episodes: 1000, CreatedAt: time.Now(), UpdatedAt: time.Now()},
		}

		c.HTML(http.StatusOK, "anime.html", animes)
	})

	router.Run(":8080")
}

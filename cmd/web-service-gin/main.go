package main

import (
	"fmt"
	"net/http"

	"github.com/abhishekghosh-in/web-service-gin/api/handler"
	"github.com/abhishekghosh-in/web-service-gin/internal/database"
	"github.com/gin-gonic/gin"
)

func main() {
	// Initializing in-memory mock database.
	database.Init()

	// Configuring router.
	router := gin.Default()
	router.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "Welcome to Gin Music API.",
		})
	})
	router.GET("/albums", handler.ReturnGetAlbumsHandler())
	router.GET("/albums/:id", handler.ReturnGetAlbumByIDHandler())
	router.POST("/albums", handler.ReturnPostAlbumsHandler())

	// Starting backend server.
	port := "8080"
	host := "localhost"
	addr := fmt.Sprintf("%s:%s", host, port)
	fmt.Println("Starting server...")
	router.Run(addr)
}

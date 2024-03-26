package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/abhishekghosh-in/web-service-gin/api/handler"
	"github.com/abhishekghosh-in/web-service-gin/internal/database"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	// Loading environment variables.
	godotenv.Load(".env")
	// Connecting with DB.
	dbConnection, err := database.Init()
	if err != nil {
		log.Fatal("DB connection failed.")
	}
	// Configuring router.
	router := gin.Default()
	router.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "Welcome to Gin Music API.",
		})
	})
	router.GET("/albums", handler.ReturnGetAlbumsHandler(&dbConnection))
	router.GET("/albums/:id", handler.ReturnGetAlbumByIDHandler(&dbConnection))
	router.POST("/albums", handler.ReturnPostAlbumsHandler(&dbConnection))

	// Starting backend server.
	port := "8080"
	host := "localhost"
	addr := fmt.Sprintf("%s:%s", host, port)
	fmt.Println("Starting server...")
	router.Run(addr)
}

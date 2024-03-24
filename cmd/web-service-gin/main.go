package main

import (
	"fmt"
	"net/http"

	"github.com/abhishekghosh-in/web-service-gin/api/handler"
	"github.com/abhishekghosh-in/web-service-gin/internal/database"
	"github.com/abhishekghosh-in/web-service-gin/internal/utils"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	// Loading environment variables.
	godotenv.Load(".env")
	// Initializing in-memory mock database.
	database.Init()

	mongoDbURI := utils.EnvVarOrFallback("MONGODB_URI", "mongodb://localhost:27017")
	_ = mongoDbURI
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

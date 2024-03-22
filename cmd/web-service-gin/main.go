package main

import (
	"fmt"

	"github.com/abhishekghosh-in/web-service-gin/internal/database"
	"github.com/gin-gonic/gin"
)

func main() {
	// Initializing in-memory mock database.
	database.Init()

	// Configuring router.
	router := gin.Default()

	// Starting backend server.
	port := "8080"
	addr := fmt.Sprintf("localhost:%s", port)
	fmt.Println("Starting server...")
	router.Run(addr)
}

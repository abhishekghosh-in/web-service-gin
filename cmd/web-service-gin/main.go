package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func main() {
	port := "8080"
	addr := fmt.Sprintf("localhost:%s", port)
	fmt.Println("Starting server...")
	router := gin.Default()
	router.Run(addr)
}

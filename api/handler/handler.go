package handler

import (
	"net/http"

	"github.com/abhishekghosh-in/web-service-gin/internal/database"
	"github.com/abhishekghosh-in/web-service-gin/internal/models"
	"github.com/gin-gonic/gin"
)

func ReturnGetAlbumsHandler() func(*gin.Context) {
	return getAlbums
}

func ReturnGetAlbumByIDHandler() func(*gin.Context) {
	return getAlbumByID
}

func ReturnPostAlbumsHandler() func(*gin.Context) {
	return postAlbums
}

// GetAlbums responds with the list of all albums as JSON.
func getAlbums(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, database.GetAllAlbums())
}

// getAlbumByID locates the album whose ID value matches the id
// parameter sent by the client, then returns that album as a response.
func getAlbumByID(c *gin.Context) {
	id := c.Param("id")
	var album *models.Album = database.GetSpecificAlbum(id)
	if album == nil {
		// Album not found in database with the given ID.
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "album not found"})
		return
	}
	c.IndentedJSON(http.StatusOK, *album)
}

// PostAlbums adds an album from JSON received in the request body.
func postAlbums(c *gin.Context) {
	var newAlbum models.Album

	// Call BindJSON to bind the received JSON to newAlbum.
	if err := c.BindJSON(&newAlbum); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "request body not matching with album configuration."})
		return
	}

	// Add the new album to the slice.
	database.AppendNewAlbum(&newAlbum)
	// Adding it to the response body.
	c.IndentedJSON(http.StatusCreated, newAlbum)
}

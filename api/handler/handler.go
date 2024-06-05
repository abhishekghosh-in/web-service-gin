package handler

import (
	"net/http"

	"github.com/abhishekghosh-in/web-service-gin/internal/database"
	"github.com/abhishekghosh-in/web-service-gin/internal/models"
	"github.com/gin-gonic/gin"
)

func ReturnGetAlbumsHandler(dbConn *database.DbConn) func(*gin.Context) {
	// getAlbums responds with the list of all albums as JSON.
	getAlbums := func(c *gin.Context) {
		albums, err := dbConn.GetAllAlbums()
		if err != nil {
			c.IndentedJSON(http.StatusNotFound, gin.H{"message": err.Error()})
			return
		}
		c.IndentedJSON(http.StatusOK, albums)
	}
	return getAlbums
}

func ReturnGetAlbumByIDHandler(dbConn *database.DbConn) func(*gin.Context) {
	// getAlbumByID locates the album whose ID value matches the id parameter
	// sent by the client, then returns that album as a response.
	getAlbumByID := func(c *gin.Context) {
		id := c.Param("id")
		album, err := dbConn.GetSpecificAlbum(id)
		if err != nil {
			// Album not found in database with the given ID.
			c.IndentedJSON(http.StatusNotFound, gin.H{"message": err.Error()})
			return
		}
		c.IndentedJSON(http.StatusOK, *album)
	}
	return getAlbumByID
}

func ReturnPostAlbumsHandler(dbConn *database.DbConn) func(*gin.Context) {
	// postAlbums adds an album from JSON received in the request body.
	postAlbums := func(c *gin.Context) {
		var newAlbum models.Album
		// Call BindJSON to bind the received JSON to newAlbum.
		if err := c.BindJSON(&newAlbum); err != nil {
			c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "request body not matching with album configuration."})
			return
		}

		dbConn.AppendNewAlbum(&newAlbum)
		// Adding it to the response body.
		c.IndentedJSON(http.StatusCreated, newAlbum)
	}
	return postAlbums
}

package database

import "github.com/abhishekghosh-in/web-service-gin/internal/models"

var albums = []models.Album{}

func Init() {
	albums = getSampleAlbums()
}

func GetAllAlbums() []models.Album {
	return albums
}

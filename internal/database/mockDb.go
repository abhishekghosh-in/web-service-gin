package database

import "github.com/abhishekghosh-in/web-service-gin/internal/models"

var albums = []models.Album{}

func Init() {
	albums = getSampleAlbums()
}

func GetAllAlbums() []models.Album {
	return albums
}

func AppendNewAlbum(newAlbum *models.Album) {
	albums = append(albums, *newAlbum)
}

func GetSpecificAlbum(id string) *models.Album {
	for _, a := range albums {
		if a.ID == id {
			aCopy := a
			return &aCopy
		}
	}
	return nil
}

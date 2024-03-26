package database

import (
	"context"

	"github.com/abhishekghosh-in/web-service-gin/internal/models"
	"github.com/abhishekghosh-in/web-service-gin/internal/utils"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// DbConn represents meta-data for DB connection.
type DbConn struct {
	DbClient        *mongo.Client     `json:"dbClient"`
	MusicCollection *mongo.Collection `json:"musicCollection"`
}

func Init() (DbConn, error) {
	mongoDbURI := utils.EnvVarOrFallback("MONGODB_URI", "mongodb://localhost:27017")
	dbName := utils.EnvVarOrFallback("DB_NAME", "ginpractice")
	collectionName := utils.EnvVarOrFallback("COLLECTION_NAME", "Music")
	// Creating DbConn to share.
	var dbConnection DbConn
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(mongoDbURI))
	if err != nil {
		// Empty connection and error.
		return dbConnection, err
	}
	dbConnection.DbClient = client
	dbConnection.MusicCollection = dbConnection.DbClient.Database(dbName).Collection(collectionName)
	return dbConnection, nil
}

func (db *DbConn) GetAllAlbums() ([]models.Album, error) {
	dbCursor, err := db.MusicCollection.Find(context.TODO(), bson.D{})
	if err != nil {
		return nil, err
	}
	defer dbCursor.Close(context.TODO())
	// Appending albums to the result using cursor traversal.
	albumsList := make([]models.Album, 0)
	for dbCursor.Next(context.TODO()) {
		album := models.Album{}
		// Decoding individual value into struct.
		err = dbCursor.Decode(&album)
		if err != nil {
			return nil, err
		}
		albumsList = append(albumsList, album)
	}
	return albumsList, nil
}

func (db *DbConn) GetSpecificAlbum(id string) (*models.Album, error) {
	albumQuery := db.MusicCollection.FindOne(context.TODO(), bson.D{{Key: "id", Value: id}})
	if albumQuery.Err() != nil {
		return nil, albumQuery.Err()
	}
	var album models.Album
	err := albumQuery.Decode(&album)
	if err != nil {
		return nil, err
	}
	return &album, nil
}

func (db *DbConn) AppendNewAlbum(newAlbum *models.Album) error {
	_, err := db.MusicCollection.InsertOne(context.TODO(), newAlbum)
	if err != nil {
		return err
	}
	return nil
}

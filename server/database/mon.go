package database

import (
	"context"
	"log"
	"time"

	"github.com/lareii/siker.im/utils"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
)

const databaseName = "db"

type mongoInstance struct {
	Client   *mongo.Client
	Database *mongo.Database
}

func newMongo() (*mongoInstance, error) {
	uri := utils.GetEnv("MONGO_URI")

	client, err := mongo.Connect(options.Client().ApplyURI(uri))
	if err != nil {
		return nil, err
	}

	db := client.Database(databaseName)
	if err := ensureCollections(db); err != nil {
		return nil, err
	}

	return &mongoInstance{
		Client:   client,
		Database: db,
	}, nil
}

func (m *mongoInstance) close() {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := m.Client.Disconnect(ctx); err != nil {
		log.Fatal("Error disconnecting from MongoDB: ", err)
	}
}

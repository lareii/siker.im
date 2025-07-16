package repository

import (
	"context"
	"time"

	"github.com/lareii/siker.im/internal/database"
	"github.com/lareii/siker.im/internal/models"

	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

type URLRepository struct {
	collection *mongo.Collection
}

func NewURLRepository(db *database.MongoDB) *URLRepository {
	return &URLRepository{
		collection: db.Database().Collection("urls"),
	}
}

func (r *URLRepository) Create(ctx context.Context, url *models.URL) error {
	url.CreatedAt = bson.Timestamp{T: uint32(time.Now().Unix())}
	url.IsActive = true
	url.ClickCount = 0

	result, err := r.collection.InsertOne(ctx, url)
	if err != nil {
		return err
	}

	url.ID = result.InsertedID.(bson.ObjectID)
	return nil
}

func (r *URLRepository) GetBySlug(ctx context.Context, slug string) (*models.URL, error) {
	var url models.URL
	filter := bson.M{"slug": slug}

	err := r.collection.FindOne(ctx, filter).Decode(&url)
	if err != nil {
		return nil, err
	}

	return &url, nil
}

func (r *URLRepository) GetByID(ctx context.Context, id bson.ObjectID) (*models.URL, error) {
	var url models.URL
	filter := bson.M{"_id": id}

	err := r.collection.FindOne(ctx, filter).Decode(&url)
	if err != nil {
		return nil, err
	}

	return &url, nil
}

// func (r *URLRepository) GetAll(ctx context.Context, limit, offset int64) ([]*models.URL, error) {
// 	filter := bson.M{"is_active": true}
// 	options := options.Find().
// 		SetLimit(limit).
// 		SetSkip(offset).
// 		SetSort(bson.M{"created_at": -1})

// 	cursor, err := r.collection.Find(ctx, filter, options)
// 	if err != nil {
// 		return nil, err
// 	}
// 	defer cursor.Close(ctx)

// 	var urls []*models.URL
// 	for cursor.Next(ctx) {
// 		var url models.URL
// 		if err := cursor.Decode(&url); err != nil {
// 			return nil, err
// 		}
// 		urls = append(urls, &url)
// 	}

// 	return urls, cursor.Err()
// }

func (r *URLRepository) IncrementClickCount(ctx context.Context, slug string) {
	filter := bson.M{"slug": slug, "is_active": true}
	update := bson.M{
		"$inc": bson.M{"click_count": 1},
	}

	r.collection.UpdateOne(ctx, filter, update)
}

// func (r *URLRepository) Delete(ctx context.Context, id bson.ObjectID) error {
// 	filter := bson.M{"_id": id}
// 	update := bson.M{
// 		"$set": bson.M{
// 			"is_active": false,
// 		},
// 	}

// 	_, err := r.collection.UpdateOne(ctx, filter, update)
// 	return err
// }

func (r *URLRepository) ExistsBySlug(ctx context.Context, slug string) (bool, error) {
	filter := bson.M{"slug": slug}
	count, err := r.collection.CountDocuments(ctx, filter)
	return count > 0, err
}

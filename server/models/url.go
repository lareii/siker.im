package models

import (
	"go.mongodb.org/mongo-driver/v2/bson"
)

type URL struct {
	ID        bson.ObjectID  `bson:"_id" json:"id"`
	CreatedAt bson.Timestamp `bson:"created_at" json:"created_at"`
	Original  string         `bson:"original" json:"original"`
	Shortened string         `bson:"shortened" json:"shortened"`
	Clicks    int            `bson:"clicks" json:"clicks"`
	Active    bool           `bson:"active" json:"active"`
}

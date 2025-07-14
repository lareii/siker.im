package models

import (
	"go.mongodb.org/mongo-driver/v2/bson"
)

type URL struct {
	ID         bson.ObjectID  `json:"id" bson:"_id,omitempty"`
	TargetURL  string         `json:"target_url" bson:"target_url"`
	Slug       string         `json:"slug" bson:"slug"`
	CreatedAt  bson.Timestamp `json:"created_at" bson:"created_at"`
	ClickCount int64          `json:"click_count" bson:"click_count"`
	IsActive   bool           `json:"is_active" bson:"is_active"`
}

type CreateURLRequest struct {
	TargetURL string `json:"target_url" validate:"required,url_valid,max=2048"`
	Slug      string `json:"slug,omitempty" validate:"omitempty,slug_valid,max=50"`
}

type URLResponse struct {
	ID         string         `json:"id"`
	TargetURL  string         `json:"target_url"`
	Slug       string         `json:"slug"`
	CreatedAt  bson.Timestamp `json:"created_at"`
	ClickCount int64          `json:"click_count"`
	IsActive   bool           `json:"is_active"`
}

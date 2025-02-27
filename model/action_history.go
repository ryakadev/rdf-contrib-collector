package model

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type ActionHistory struct {
	ID            primitive.ObjectID `bson:"id,omitempty"`
	Action        string             `bson:"action"`
	Repository    string             `bson:"repository"`
	PullRequestID string             `bson:"pull_request_id"`
	Point         int64              `bson:"point"`
	CreatedAt     time.Time          `bson:"created_at"`
}

package model

import "go.mongodb.org/mongo-driver/bson/primitive"

type Points struct {
	ID        primitive.ObjectID `bson:"id,omitempty"`
	ContribID primitive.ObjectID `bson:"contrib_id,omitempty"`
	Point     int64              `bson:"point"`
}

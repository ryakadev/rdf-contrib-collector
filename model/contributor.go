package model

import "go.mongodb.org/mongo-driver/bson/primitive"

type Contributor struct {
	ID   primitive.ObjectID `bson:"id,omitempty"`
	Name string             `bson:"name,omitempty"`
	Role string             `bson:"role,omitempty"`
}

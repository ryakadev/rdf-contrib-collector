package repository

import "go.mongodb.org/mongo-driver/v2/mongo"

type Repository interface {
}

type repository struct {
	db *mongo.Client
}

func New(db *mongo.Client) Repository {
	return &repository{
		db: db,
	}
}

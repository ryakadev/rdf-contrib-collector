package repository

import (
	"context"

	"github.com/ryakadev/rdf-contrib-collector/model"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// CreatePoint implements Repository.
func (r *repository) CreatePoint(ctx context.Context, payload *model.Point) error {
	panic("unimplemented")
}

// UpdatePoint implements Repository.
func (r *repository) UpdatePoint(ctx context.Context, payload *model.Contributor) error {
	panic("unimplemented")
}

// GetPoint implements Repository.
func (r *repository) GetPoint(ctx context.Context, id primitive.ObjectID) (model.Point, error) {
	panic("unimplemented")
}

// GetPoints implements Repository.
func (r *repository) GetPoints(ctx context.Context, offset int64, limit int64, filter *model.Point) ([]model.Point, error) {
	panic("unimplemented")
}

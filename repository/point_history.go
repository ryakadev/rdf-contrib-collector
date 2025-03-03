package repository

import (
	"context"

	"github.com/ryakadev/rdf-contrib-collector/model"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// CreatePointHistory implements Repository.
func (r *repository) CreatePointHistory(ctx context.Context, payload *model.PointHistory) error {
	panic("unimplemented")
}

// GetPointHistories implements Repository.
func (r *repository) GetPointHistories(ctx context.Context, offset int64, limit int64, filter *model.PointHistory) ([]model.ActionHistory, error) {
	panic("unimplemented")
}

// GetPointHistory implements Repository.
func (r *repository) GetPointHistory(ctx context.Context, id primitive.ObjectID) (model.PointHistory, error) {
	panic("unimplemented")
}

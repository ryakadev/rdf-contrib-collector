package repository

import (
	"context"

	"github.com/ryakadev/rdf-contrib-collector/model"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func (r *repository) CreateActionHistory(ctx context.Context, payload *model.ActionHistory) error {
	panic("unimplemented")
}

func (r *repository) GetActionHistories(ctx context.Context, offset int64, limit int64, filter *model.ActionHistory) ([]model.ActionHistory, error) {
	panic("unimplemented")
}

func (r *repository) GetActionHistory(ctx context.Context, id primitive.ObjectID) (model.ActionHistory, error) {
	panic("unimplemented")
}

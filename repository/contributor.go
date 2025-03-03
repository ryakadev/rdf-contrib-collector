package repository

import (
	"context"

	"github.com/ryakadev/rdf-contrib-collector/model"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// CreateContributor implements Repository.
func (r *repository) CreateContributor(ctx context.Context, payload *model.Contributor) error {
	panic("unimplemented")
}

// GetContributor implements Repository.
func (r *repository) GetContributor(ctx context.Context, id primitive.ObjectID) (model.Contributor, error) {
	panic("unimplemented")
}

// GetContributors implements Repository.
func (r *repository) GetContributors(ctx context.Context, offset int64, limit int64, filter *model.Contributor) ([]model.Contributor, error) {
	panic("unimplemented")
}

// UpdateContributor implements Repository.
func (r *repository) UpdateContributor(ctx context.Context, payload *model.Contributor) error {
	panic("unimplemented")
}

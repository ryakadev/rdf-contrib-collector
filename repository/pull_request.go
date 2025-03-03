package repository

import (
	"context"

	"github.com/ryakadev/rdf-contrib-collector/model"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// CreatePullRequest implements Repository.
func (r *repository) CreatePullRequest(ctx context.Context, payload *model.PullRequest) error {
	panic("unimplemented")
}

// GetPullRequest implements Repository.
func (r *repository) GetPullRequest(ctx context.Context, id primitive.ObjectID) (model.PullRequest, error) {
	panic("unimplemented")
}

// GetPullRequests implements Repository.
func (r *repository) GetPullRequests(ctx context.Context, offset int64, limit int64, filter *model.PullRequest) ([]model.PullRequest, error) {
	panic("unimplemented")
}

// UpdatePullRequest implements Repository.
func (r *repository) UpdatePullRequest(ctx context.Context, payload *model.Contributor) error {
	panic("unimplemented")
}

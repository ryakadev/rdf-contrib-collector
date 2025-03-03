package repository

import (
	"context"

	"github.com/ryakadev/rdf-contrib-collector/model"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// CreateGitRepo implements Repository.
func (r *repository) CreateGitRepo(ctx context.Context, payload *model.GitRepo) error {
	panic("unimplemented")
}

// GetGitRepo implements Repository.
func (r *repository) GetGitRepo(ctx context.Context, id primitive.ObjectID) (model.GitRepo, error) {
	panic("unimplemented")
}

// GetGitRepos implements Repository.
func (r *repository) GetGitRepos(ctx context.Context, offset int64, limit int64, filter *model.GitRepo) ([]model.GitRepo, error) {
	panic("unimplemented")
}

// UpdateGitRepo implements Repository.
func (r *repository) UpdateGitRepo(ctx context.Context, payload *model.Contributor) error {
	panic("unimplemented")
}

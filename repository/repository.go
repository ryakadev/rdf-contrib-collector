package repository

import (
	"context"

	"github.com/ryakadev/rdf-contrib-collector/model"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

type Repository interface {
	CreateActionHistory(ctx context.Context, payload *model.ActionHistory) error
	GetActionHistory(ctx context.Context, id primitive.ObjectID) (model.ActionHistory, error)
	GetActionHistories(ctx context.Context, offset int64, limit int64, filter *model.ActionHistory) ([]model.ActionHistory, error)

	CreateContributor(ctx context.Context, payload *model.Contributor) error
	UpdateContributor(ctx context.Context, payload *model.Contributor) error
	GetContributor(ctx context.Context, id primitive.ObjectID) (model.Contributor, error)
	GetContributors(ctx context.Context, offset int64, limit int64, filter *model.Contributor) ([]model.Contributor, error)

	CreatePoint(ctx context.Context, payload *model.Point) error
	UpdatePoint(ctx context.Context, payload *model.Contributor) error
	GetPoint(ctx context.Context, id primitive.ObjectID) (model.Point, error)
	GetPoints(ctx context.Context, offset int64, limit int64, filter *model.Point) ([]model.Point, error)

	CreatePointHistory(ctx context.Context, payload *model.PointHistory) error
	GetPointHistory(ctx context.Context, id primitive.ObjectID) (model.PointHistory, error)
	GetPointHistories(ctx context.Context, offset int64, limit int64, filter *model.PointHistory) ([]model.ActionHistory, error)

	CreatePullRequest(ctx context.Context, payload *model.PullRequest) error
	UpdatePullRequest(ctx context.Context, payload *model.Contributor) error
	GetPullRequest(ctx context.Context, id primitive.ObjectID) (model.PullRequest, error)
	GetPullRequests(ctx context.Context, offset int64, limit int64, filter *model.PullRequest) ([]model.PullRequest, error)

	CreateGitRepo(ctx context.Context, payload *model.GitRepo) error
	UpdateGitRepo(ctx context.Context, payload *model.Contributor) error
	GetGitRepo(ctx context.Context, id primitive.ObjectID) (model.GitRepo, error)
	GetGitRepos(ctx context.Context, offset int64, limit int64, filter *model.GitRepo) ([]model.GitRepo, error)
}

type repository struct {
	mongo *mongo.Client
}

func New(mongo *mongo.Client) Repository {
	return &repository{
		mongo: mongo,
	}
}

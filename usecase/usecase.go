package usecase

import (
	"github.com/ryakadev/rdf-contrib-collector/repository"
	"go.uber.org/zap"
)

type UseCase interface {
	HandleWebhook(event interface{}) error
}

type usecase struct {
	repo repository.Repository
	log  *zap.Logger
}

func New(repo repository.Repository, log *zap.Logger) UseCase {
	return &usecase{
		repo: repo,
		log:  log,
	}
}

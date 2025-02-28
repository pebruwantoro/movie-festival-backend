package viewerships

import (
	"context"
	"github.com/pebruwantoro/movie-festival-backend/internal/app/entities"
)

//go:generate mockgen -destination=mocks/repository.go -package=mocks . RepositoryInterface
type RepositoryInterface interface {
	CreateOrUpdateViewership(ctx context.Context, request entities.Viewership) (response entities.Viewership, err error)
	GetViewershipByUserUUID(ctx context.Context, uuid string) (response []entities.Viewership, err error)
}

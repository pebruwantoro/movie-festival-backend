package movies

import (
	"context"
	"github.com/pebruwantoro/movie-festival-backend/internal/app/entities"
)

//go:generate mockgen -destination=mocks/repository.go -package=mocks . RepositoryInterface
type RepositoryInterface interface {
	CreateMovie(ctx context.Context, request entities.Movie) (response entities.Movie, err error)
	UpdateMovie(ctx context.Context, request entities.Movie) (response entities.Movie, err error)
	GetMovieByUUID(ctx context.Context, uuid string) (response entities.Movie, err error)
	GetMovieByUUIDs(ctx context.Context, uuids []string) (response []entities.Movie, err error)
}

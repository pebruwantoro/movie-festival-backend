package createorupdate

import (
	"github.com/pebruwantoro/movie-festival-backend/internal/app/repositories/movies"
	"github.com/pebruwantoro/movie-festival-backend/internal/app/repositories/viewerships"
)

type Usecase struct {
	viewershipRepo viewerships.RepositoryInterface
	movieRepo      movies.RepositoryInterface
}

func NewUsecase(viewershipRepo viewerships.RepositoryInterface, movieRepo movies.RepositoryInterface) *Usecase {
	return &Usecase{
		viewershipRepo: viewershipRepo,
		movieRepo:      movieRepo,
	}
}

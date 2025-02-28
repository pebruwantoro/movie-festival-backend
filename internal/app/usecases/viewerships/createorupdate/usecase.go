package createorupdate

import (
	"github.com/pebruwantoro/movie-festival-backend/internal/app/repositories/movies"
	"github.com/pebruwantoro/movie-festival-backend/internal/app/repositories/viewerships"
)

type Usecase struct {
	viewershipRepo viewerships.Repository
	movieRepo      movies.Repository
}

func NewUsecase(viewershipRepo viewerships.Repository, movieRepo movies.Repository) *Usecase {
	return &Usecase{
		viewershipRepo: viewershipRepo,
		movieRepo:      movieRepo,
	}
}

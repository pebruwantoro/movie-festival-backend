package getmostviewedmovie

import (
	"github.com/pebruwantoro/movie-festival-backend/internal/app/repositories/movies"
)

type Usecase struct {
	movieRepo movies.RepositoryInterface
}

func NewUsecase(movieRepo movies.RepositoryInterface) *Usecase {
	return &Usecase{
		movieRepo: movieRepo,
	}
}

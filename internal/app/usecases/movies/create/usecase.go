package create

import (
	"github.com/pebruwantoro/movie-festival-backend/internal/app/repositories/movies"
)

type Usecase struct {
	movieRepo movies.Repository
}

func NewUsecase(movieRepo movies.Repository) *Usecase {
	return &Usecase{
		movieRepo: movieRepo,
	}
}

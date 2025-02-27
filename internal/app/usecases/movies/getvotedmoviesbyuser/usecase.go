package getvotedmoviesbyuser

import (
	"github.com/pebruwantoro/movie-festival-backend/internal/app/repositories/movies"
	"github.com/pebruwantoro/movie-festival-backend/internal/app/repositories/voters"
)

type Usecase struct {
	voteRepo  voters.Repository
	movieRepo movies.Repository
}

func NewUsecase(voteRepo voters.Repository, movieRepo movies.Repository) *Usecase {
	return &Usecase{
		voteRepo:  voteRepo,
		movieRepo: movieRepo,
	}
}

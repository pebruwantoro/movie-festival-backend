package getvotedmoviesbyuser

import (
	"github.com/pebruwantoro/movie-festival-backend/internal/app/repositories/movies"
	"github.com/pebruwantoro/movie-festival-backend/internal/app/repositories/voters"
)

type Usecase struct {
	voteRepo  voters.RepositoryInterface
	movieRepo movies.RepositoryInterface
}

func NewUsecase(voteRepo voters.RepositoryInterface, movieRepo movies.RepositoryInterface) *Usecase {
	return &Usecase{
		voteRepo:  voteRepo,
		movieRepo: movieRepo,
	}
}

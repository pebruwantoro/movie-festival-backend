package unvote

import "github.com/pebruwantoro/movie-festival-backend/internal/app/repositories/voters"

type Usecase struct {
	voteRepo voters.Repository
}

func NewUsecase(voteRepo voters.Repository) *Usecase {
	return &Usecase{
		voteRepo: voteRepo,
	}
}

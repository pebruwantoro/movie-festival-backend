package vote

import (
	"github.com/pebruwantoro/movie-festival-backend/internal/app/repositories/voters"
)

type Usecase struct {
	voteRepo voters.RepositoryInterface
}

func NewUsecase(voteRepo voters.RepositoryInterface) *Usecase {
	return &Usecase{
		voteRepo: voteRepo,
	}
}

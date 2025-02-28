package create

import (
	"github.com/pebruwantoro/movie-festival-backend/internal/app/repositories/users"
)

type Usecase struct {
	userRepo users.RepositoryInterface
}

func NewUsecase(userRepo users.RepositoryInterface) *Usecase {
	return &Usecase{
		userRepo: userRepo,
	}
}

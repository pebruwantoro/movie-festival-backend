package create

import (
	"github.com/pebruwantoro/movie-festival-backend/internal/app/repositories/users"
)

type Usecase struct {
	Repository users.Repository
}

func NewUsecase(userRepo users.Repository) *Usecase {
	return &Usecase{
		Repository: userRepo,
	}
}

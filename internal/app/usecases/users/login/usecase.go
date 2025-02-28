package login

import (
	"github.com/pebruwantoro/movie-festival-backend/internal/app/repositories/token"
	"github.com/pebruwantoro/movie-festival-backend/internal/app/repositories/users"
)

type Usecase struct {
	userRepo  users.RepositoryInterface
	tokenRepo token.RepositoryInterface
}

func NewUsecase(userRepo users.RepositoryInterface, tokenRepo token.RepositoryInterface) *Usecase {
	return &Usecase{
		userRepo:  userRepo,
		tokenRepo: tokenRepo,
	}
}

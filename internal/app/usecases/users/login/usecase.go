package login

import (
	"github.com/pebruwantoro/movie-festival-backend/internal/app/repositories/token"
	"github.com/pebruwantoro/movie-festival-backend/internal/app/repositories/users"
)

type Usecase struct {
	userRepo  users.Repository
	tokenRepo token.Repository
}

func NewUsecase(userRepo users.Repository, tokenRepo token.Repository) *Usecase {
	return &Usecase{
		userRepo:  userRepo,
		tokenRepo: tokenRepo,
	}
}

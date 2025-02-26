package logout

import (
	"github.com/pebruwantoro/movie-festival-backend/internal/app/repositories/token"
)

type Usecase struct {
	tokenRepo token.Repository
}

func NewUsecase(tokenRepo token.Repository) *Usecase {
	return &Usecase{
		tokenRepo: tokenRepo,
	}
}

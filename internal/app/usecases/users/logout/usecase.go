package logout

import (
	"github.com/pebruwantoro/movie-festival-backend/internal/app/repositories/token"
)

type Usecase struct {
	tokenRepo token.RepositoryInterface
}

func NewUsecase(tokenRepo token.RepositoryInterface) *Usecase {
	return &Usecase{
		tokenRepo: tokenRepo,
	}
}

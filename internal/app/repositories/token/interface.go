package token

import (
	"context"
	"github.com/pebruwantoro/movie-festival-backend/internal/app/entities"
)

//go:generate mockgen -destination=mocks/repository.go -package=mocks . RepositoryInterface
type RepositoryInterface interface {
	CreateToken(ctx context.Context, request entities.Token) (response entities.Token, err error)
	UpdateToken(ctx context.Context, request entities.Token) (response entities.Token, err error)
	GetActiveTokenByJWT(ctx context.Context, jwt string) (response entities.Token, err error)
}

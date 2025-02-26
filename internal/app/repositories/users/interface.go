package users

import (
	"context"
	"github.com/pebruwantoro/movie-festival-backend/internal/app/entities"
)

//go:generate mockgen -destination=mocks/repository.go -package=mocks . RepositoryInterface
type RepositoryInterface interface {
	CreateUser(ctx context.Context, request entities.User) (response entities.User, err error)
	GetUserByEmail(ctx context.Context, email string) (response entities.User, err error)
}

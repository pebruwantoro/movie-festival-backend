package voters

import (
	"context"
	"github.com/pebruwantoro/movie-festival-backend/internal/app/entities"
)

//go:generate mockgen -destination=mocks/repository.go -package=mocks . RepositoryInterface
type RepositoryInterface interface {
	CreateVoter(ctx context.Context, request entities.Voters) (response entities.Voters, err error)
	DeleteVoter(ctx context.Context, request entities.Voters) (response entities.Voters, err error)
	GetVotersByUserUUID(ctx context.Context, uuid string) (response []entities.Voters, err error)
}

package create

import (
	"context"
	"github.com/pebruwantoro/movie-festival-backend/internal/app/entities"
)

func (u *Usecase) Execute(ctx context.Context, request CreateUserRequest) (response CreateUserResponse, err error) {
	user := entities.User{}

	data, err := request.MapIntoUser()
	if err != nil {
		return
	}

	user, err = u.userRepo.CreateUser(ctx, data)

	response.UUID = user.UUID

	return
}

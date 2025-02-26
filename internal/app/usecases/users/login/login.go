package login

import (
	"context"
	"github.com/pebruwantoro/movie-festival-backend/internal/app/entities"
	"github.com/pebruwantoro/movie-festival-backend/internal/pkg/helper"
)

func (u *Usecase) Execute(ctx context.Context, request LoginRequest) (response LoginResponse, err error) {
	user := entities.User{}
	token := ""

	user, err = u.userRepo.GetUserByEmail(ctx, request.Email)
	if err != nil {
		return
	}

	if !helper.ValidatePassword(user.Password, request.Password) {
		err = helper.ErrorInvalidPassword
		return
	}

	token, err = helper.GenerateJWT(user)
	if err != nil {
		return
	}

	dataToken := MapIntoToken(token, user)
	_, err = u.tokenRepo.CreateToken(ctx, dataToken)
	if err != nil {
		return
	}

	response.Token = token

	return
}

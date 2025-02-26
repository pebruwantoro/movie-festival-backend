package logout

import (
	"context"
	"errors"
	"github.com/pebruwantoro/movie-festival-backend/internal/app/entities"
	"github.com/pebruwantoro/movie-festival-backend/internal/pkg/helper"
	"gorm.io/gorm"
)

func (u *Usecase) Execute(ctx context.Context, request LogoutRequest) (response LogoutResponse, err error) {
	token := entities.Token{}

	token, err = u.tokenRepo.GetActiveTokenByJWT(ctx, request.Token)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			err = helper.ErrorUserAlreadyLogout
		}
		return
	}

	token.SetInActive(token.CreatedBy)

	token, err = u.tokenRepo.UpdateToken(ctx, token)
	if err != nil {
		return
	}

	response.UserUUID = token.UserUUID

	return
}

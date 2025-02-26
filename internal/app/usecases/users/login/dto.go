package login

import (
	"github.com/google/uuid"
	"github.com/pebruwantoro/movie-festival-backend/internal/app/entities"
)

type LoginRequest struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,gte=8,lte=16"`
}

type LoginResponse struct {
	Token string `json:"token"`
}

func MapIntoToken(token string, user entities.User) entities.Token {
	data := entities.Token{
		UUID:     uuid.NewString(),
		UserUUID: user.UUID,
		Token:    token,
		IsActive: true,
	}
	data.SetCreated(user.Email)

	return data
}

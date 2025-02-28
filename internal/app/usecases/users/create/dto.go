package create

import (
	"github.com/google/uuid"
	"github.com/pebruwantoro/movie-festival-backend/internal/app/entities"
	"github.com/pebruwantoro/movie-festival-backend/internal/pkg/helper"
)

type CreateUserRequest struct {
	Name     string `json:"name" validate:"required"`
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,gte=8,lte=16,user_password"`
	Role     string `json:"role" swaggerignore:"true"`
}

type CreateUserResponse struct {
	UUID string `json:"uuid"`
}

func (r *CreateUserRequest) MapIntoUser() (entities.User, error) {
	hashPassword, err := helper.HashPassword(r.Password)
	if err != nil {
		return entities.User{}, err
	}

	user := entities.User{
		UUID:     uuid.NewString(),
		Name:     r.Name,
		Email:    r.Email,
		Password: hashPassword,
		Role:     r.Role,
	}
	user.SetCreated(r.Email)
	user.SetUpdated(r.Email)

	return user, nil
}

package users

import (
	"context"
	"github.com/pebruwantoro/movie-festival-backend/internal/app/entities"
)

func (r *Repository) CreateUser(ctx context.Context, request entities.User) (response entities.User, err error) {
	res := r.Db.WithContext(ctx).
		Omit("DeletedAt", "DeletedBy").
		Create(&request)

	response.UUID = request.UUID
	err = res.Error

	return
}

func (r *Repository) GetUserByEmail(ctx context.Context, email string) (response entities.User, err error) {
	res := r.Db.WithContext(ctx).
		Table(entities.USERS_TABLE).
		Where("email = ?", email).
		First(&response)

	err = res.Error

	return
}

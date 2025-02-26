package token

import (
	"context"
	"github.com/pebruwantoro/movie-festival-backend/internal/app/entities"
)

func (r *Repository) CreateToken(ctx context.Context, request entities.Token) (response entities.Token, err error) {
	res := r.Db.WithContext(ctx).
		Omit("UpdatedAt", "UpdatedBy").
		Create(&request)

	response.UUID = request.UUID
	err = res.Error

	return
}

func (r *Repository) UpdateToken(ctx context.Context, request entities.Token) (response entities.Token, err error) {
	res := r.Db.WithContext(ctx).
		Model(&entities.Token{}).
		Select("is_active", "updated_by").
		Where("token = ?", request.Token).
		Updates(&request)

	response.UserUUID = request.UserUUID
	err = res.Error

	return
}

func (r *Repository) GetActiveTokenByJWT(ctx context.Context, jwt string) (response entities.Token, err error) {
	res := r.Db.WithContext(ctx).
		Table(entities.TOKENS_TABLE).
		Where("token = ? AND is_active = true", jwt).
		First(&response)

	err = res.Error

	return
}

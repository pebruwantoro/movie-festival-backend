package voters

import (
	"context"
	"github.com/pebruwantoro/movie-festival-backend/internal/app/entities"
)

func (r *Repository) CreateVoter(ctx context.Context, request entities.Voters) (response entities.Voters, err error) {
	res := r.Db.WithContext(ctx).
		Omit("DeletedAt", "DeletedBy").
		Create(&request)

	response.UUID = request.UUID
	err = res.Error

	return
}

func (r *Repository) DeleteVoter(ctx context.Context, request entities.Voters) (response entities.Voters, err error) {
	res := r.Db.WithContext(ctx).
		Where("uuid = ?", request.UUID).
		Delete(&entities.Voters{})

	response.UUID = request.UUID
	err = res.Error

	return
}

func (r *Repository) GetVotersByUserUUID(ctx context.Context, uuid string) (response []entities.Voters, err error) {
	res := r.Db.WithContext(ctx).
		Where("user_uuid = ? AND deleted_at IS NULL", uuid).
		Find(&response)

	err = res.Error

	return
}

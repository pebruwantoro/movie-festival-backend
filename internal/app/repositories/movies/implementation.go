package movies

import (
	"context"
	"github.com/pebruwantoro/movie-festival-backend/internal/app/entities"
)

func (r *Repository) CreateMovie(ctx context.Context, request entities.Movie) (response entities.Movie, err error) {
	res := r.Db.WithContext(ctx).
		Omit("DeletedAt", "DeletedBy").
		Create(&request)

	response.UUID = request.UUID
	err = res.Error

	return
}

func (r *Repository) UpdateMovie(ctx context.Context, request entities.Movie) (response entities.Movie, err error) {
	res := r.Db.WithContext(ctx).
		Model(&entities.Movie{}).
		Omit("UUID", "CreatedAt", "CreatedBY", "DeletedAt", "DeletedBy").
		Where("uuid = ? AND deleted_at IS NULL", request.UUID).
		Updates(&request)

	response = request
	err = res.Error

	return
}

func (r *Repository) GetMovieByUUID(ctx context.Context, uuid string) (response entities.Movie, err error) {
	res := r.Db.WithContext(ctx).
		Table(entities.MOVIES_TABLE).
		Where("uuid = ? AND deleted_at IS NULL", uuid).
		First(&response)

	err = res.Error

	return
}

func (r *Repository) GetMovieByUUIDs(ctx context.Context, uuids []string) (response []entities.Movie, err error) {
	res := r.Db.WithContext(ctx).
		Table(entities.MOVIES_TABLE).
		Where("uuid IN (?) AND deleted_at IS NULL", uuids).
		Find(&response)

	err = res.Error

	return
}

package viewerships

import (
	"context"
	"github.com/pebruwantoro/movie-festival-backend/internal/app/entities"
	"gorm.io/gorm/clause"
)

func (r *Repository) CreateOrUpdateViewership(ctx context.Context, request entities.Viewership) (response entities.Viewership, err error) {
	res := r.Db.WithContext(ctx).Clauses(clause.OnConflict{
		Columns: []clause.Column{
			{Name: "movie_uuid"},
			{Name: "user_uuid"},
		},
		DoUpdates: clause.AssignmentColumns([]string{"watching_duration"}),
	}).Create(&request)

	response.MovieUUID = request.MovieUUID
	response.UserUUID = request.UserUUID

	err = res.Error

	return
}

func (r *Repository) GetViewershipByUserUUID(ctx context.Context, uuid string) (response []entities.Viewership, err error) {
	res := r.Db.WithContext(ctx).
		Table(entities.VIEWERSHIPS_TABLE).
		Where("user_uuid = ?", uuid).
		Find(&response)

	err = res.Error

	return
}

package movies

import (
	"context"
	"fmt"
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

func (r *Repository) GetMoviesByFilter(ctx context.Context, filter Filter) (response []entities.Movie, err error) {
	res := r.Db.WithContext(ctx).
		Table(entities.MOVIES_TABLE)

	if filter.Title != "" {
		res = res.Where("title LIKE ?", "%"+filter.Title+"%")
	}

	if filter.Description != "" {
		res = res.Where("description LIKE ?", "%"+filter.Description+"%")
	}

	if len(filter.Artists) > 0 {
		queryUUIDs := ""
		for i, artist := range filter.Artists {
			queryUUIDs += fmt.Sprintf("'%s'::UUID", artist)
			if i < len(filter.Artists)-1 {
				queryUUIDs += ","
			}
		}

		res = res.Where(fmt.Sprintf("artists && ARRAY[%s]", queryUUIDs))
	}

	if len(filter.Genres) > 0 {
		queryUUIDs := ""
		for i, artist := range filter.Genres {
			queryUUIDs += fmt.Sprintf("'%s'::UUID", artist)
			if i < len(filter.Genres)-1 {
				queryUUIDs += ","
			}
		}

		res = res.Where(fmt.Sprintf("genres && ARRAY[%s]", queryUUIDs))
	}

	res = res.Where("deleted_at IS NULL").
		Limit(filter.Pagination.PerPage).
		Offset(filter.Pagination.PerPage * (filter.Pagination.Page - 1)).
		Find(&response)

	err = res.Error

	return
}

func (r *Repository) CountTotalMoviesByFilter(ctx context.Context, filter Filter) (total int64, err error) {
	res := r.Db.WithContext(ctx).
		Table(entities.MOVIES_TABLE)

	if filter.Title != "" {
		res = res.Where("title LIKE ?", "%"+filter.Title+"%")
	}

	if filter.Description != "" {
		res = res.Where("description LIKE ?", "%"+filter.Description+"%")
	}

	if len(filter.Artists) > 0 {
		queryUUIDs := ""
		for i, artist := range filter.Artists {
			queryUUIDs += fmt.Sprintf("'%s'::UUID", artist)
			if i < len(filter.Artists)-1 {
				queryUUIDs += ","
			}
		}

		res = res.Where(fmt.Sprintf("artists && ARRAY[%s]", queryUUIDs))
	}

	if len(filter.Genres) > 0 {
		queryUUIDs := ""
		for i, artist := range filter.Genres {
			queryUUIDs += fmt.Sprintf("'%s'::UUID", artist)
			if i < len(filter.Genres)-1 {
				queryUUIDs += ","
			}
		}

		res = res.Where(fmt.Sprintf("genres && ARRAY[%s]", queryUUIDs))
	}

	res = res.Where("deleted_at IS NULL").
		Count(&total)

	err = res.Error

	return
}

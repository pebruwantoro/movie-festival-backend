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

func (r *Repository) GetMostViewedMovie(ctx context.Context) (response entities.MovieWithViewership, err error) {
	res := r.Db.WithContext(ctx).
		Table(entities.VIEWERSHIPS_TABLE + " AS v").
		Select("m.*, COUNT(v.movie_uuid) AS total_viewed").
		Joins("JOIN " + entities.MOVIES_TABLE + " AS m ON m.uuid = v.movie_uuid").
		Where("m.deleted_at IS NULL").
		Group("m.uuid").
		Order("total_viewed DESC").
		Limit(1).
		Scan(&response)

	err = res.Error

	return
}

func (r *Repository) GetMostViewedMovieGenre(ctx context.Context) (response entities.MovieGenreWithViewership, err error) {
	query := `
		SELECT g.uuid, g.name, COUNT(v.movie_uuid) AS total_viewed
		FROM viewerships v
		JOIN movies m ON m.uuid = v.movie_uuid
		JOIN LATERAL unnest(m.genres) AS genre_uuid ON TRUE
		JOIN genres g ON g.uuid = genre_uuid
		WHERE m.deleted_at IS NULL
		GROUP BY g.uuid, g.name
		ORDER BY total_viewed DESC
		LIMIT 1;
	`

	res := r.Db.WithContext(ctx).Raw(query).Scan(&response)
	err = res.Error

	return
}

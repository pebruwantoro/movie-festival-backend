package create

import (
	"github.com/google/uuid"
	"github.com/pebruwantoro/movie-festival-backend/internal/app/entities"
)

type CreateMovieRequest struct {
	Title       string   `json:"title" validate:"required"`
	Description string   `json:"description" validate:"required"`
	Duration    int      `json:"duration" validate:"required"`
	Artists     []string `json:"artists" validate:"required"`
	Genres      []string `json:"genres" validate:"required"`
	Url         string   `json:"url" validate:"required"`
	CreatedBy   string   `json:"created_by" swaggerignore:"true"`
}

type CreateMovieResponse struct {
	UUID string `json:"uuid"`
}

func (r *CreateMovieRequest) MapIntoMovie() entities.Movie {
	movie := entities.Movie{
		UUID:        uuid.NewString(),
		Title:       r.Title,
		Description: r.Description,
		Duration:    r.Duration,
		Artists:     r.Artists,
		Genres:      r.Genres,
		Url:         r.Url,
	}
	movie.SetCreated(r.CreatedBy)
	movie.SetUpdated(r.CreatedBy)

	return movie
}

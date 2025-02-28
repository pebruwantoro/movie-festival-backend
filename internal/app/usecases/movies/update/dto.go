package update

import (
	"github.com/lib/pq"
	"github.com/pebruwantoro/movie-festival-backend/internal/app/entities"
)

type UpdateMovieRequest struct {
	UUID        string   `json:"uuid"`
	Title       string   `json:"title" validate:"required"`
	Description string   `json:"description" validate:"required"`
	Duration    int      `json:"duration" validate:"required"`
	Artists     []string `json:"artists" validate:"required"`
	Genres      []string `json:"genres" validate:"required"`
	Url         string   `json:"url" validate:"required"`
	UpdatedBy   string   `json:"created_by" swaggerignore:"true"`
}

type UpdateMovieResponse struct {
	UUID        string         `json:"uuid"`
	Title       string         `json:"title"`
	Description string         `json:"description"`
	Duration    int            `json:"duration"`
	Artists     pq.StringArray `json:"artists"`
	Genres      pq.StringArray `json:"genres"`
	Url         string         `json:"url"`
}

func (r *UpdateMovieRequest) MapIntoMovie() entities.Movie {
	movie := entities.Movie{
		UUID:        r.UUID,
		Title:       r.Title,
		Description: r.Description,
		Duration:    r.Duration,
		Artists:     r.Artists,
		Genres:      r.Genres,
		Url:         r.Url,
	}
	movie.SetUpdated(r.UpdatedBy)

	return movie
}

func (r *UpdateMovieResponse) MapIntoResponse(movie entities.Movie) {
	r.UUID = movie.UUID
	r.Title = movie.Title
	r.Description = movie.Description
	r.Duration = movie.Duration
	r.Artists = movie.Artists
	r.Genres = movie.Genres
	r.Url = movie.Url
}

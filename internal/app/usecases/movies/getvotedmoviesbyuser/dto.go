package getvotedmoviesbyuser

import (
	"github.com/lib/pq"
	"github.com/pebruwantoro/movie-festival-backend/internal/app/entities"
)

type GetVotedMovieByUserRequest struct {
	UserUUID string `json:"user_uuid" validate:"required"`
}

type GetVotedMovieByUserResponse struct {
	Data []MovieResponse `json:"data"`
}

type MovieResponse struct {
	UUID        string         `json:"uuid"`
	Title       string         `json:"title"`
	Description string         `json:"description"`
	Duration    int            `json:"duration"`
	Artists     pq.StringArray `json:"artists"`
	Genres      pq.StringArray `json:"genres"`
	Url         string         `json:"url"`
}

func (r *GetVotedMovieByUserResponse) MapIntoResponse(movies []entities.Movie) {
	for _, movie := range movies {
		r.Data = append(r.Data, MovieResponse{
			UUID:        movie.UUID,
			Title:       movie.Title,
			Description: movie.Description,
			Duration:    movie.Duration,
			Artists:     movie.Artists,
			Genres:      movie.Genres,
			Url:         movie.Url,
		})
	}
}

package getmostviewedmoviegenre

import (
	"github.com/pebruwantoro/movie-festival-backend/internal/app/entities"
)

type GetMostViewedMovieGenreRequest struct{}

type GetMostViewedMovieGenreResponse struct {
	UUID   string `json:"uuid"`
	Name   string `json:"name"`
	Viewed int    `json:"viewed"`
}

func (r *GetMostViewedMovieGenreResponse) MapIntoResponse(movieGenre entities.MovieGenreWithViewership) {
	r.UUID = movieGenre.UUID
	r.Name = movieGenre.Name
	r.Viewed = movieGenre.TotalViewed
}

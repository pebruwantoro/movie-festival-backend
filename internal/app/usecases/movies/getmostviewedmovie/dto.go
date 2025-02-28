package getmostviewedmovie

import (
	"github.com/lib/pq"
	"github.com/pebruwantoro/movie-festival-backend/internal/app/entities"
)

type GetMostViewedMovieRequest struct{}

type GetMostViewedMovieRequestResponse struct {
	UUID        string         `json:"uuid"`
	Title       string         `json:"title"`
	Description string         `json:"description"`
	Duration    int            `json:"duration"`
	Artists     pq.StringArray `json:"artists"`
	Genres      pq.StringArray `json:"genres"`
	Url         string         `json:"url"`
	Viewed      int            `json:"viewed"`
}

func (r *GetMostViewedMovieRequestResponse) MapIntoResponse(movie entities.MovieWithViewership) {
	r.UUID = movie.UUID
	r.Title = movie.Title
	r.Description = movie.Description
	r.Duration = movie.Duration
	r.Artists = movie.Artists
	r.Genres = movie.Genres
	r.Url = movie.Url
	r.Viewed = movie.TotalViewed
}

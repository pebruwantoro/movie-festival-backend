package getmostviewedmovie

import (
	"context"
	"github.com/pebruwantoro/movie-festival-backend/internal/app/entities"
)

func (u *Usecase) Execute(ctx context.Context) (response GetMostViewedMovieRequestResponse, err error) {
	movie := entities.MovieWithViewership{}

	movie, err = u.movieRepo.GetMostViewedMovie(ctx)
	if err != nil {
		return
	}

	response.MapIntoResponse(movie)

	return
}

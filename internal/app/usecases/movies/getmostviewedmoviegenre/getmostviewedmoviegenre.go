package getmostviewedmoviegenre

import (
	"context"
	"github.com/pebruwantoro/movie-festival-backend/internal/app/entities"
)

func (u *Usecase) Execute(ctx context.Context) (response GetMostViewedMovieGenreResponse, err error) {
	movie := entities.MovieGenreWithViewership{}

	movie, err = u.movieRepo.GetMostViewedMovieGenre(ctx)
	if err != nil {
		return
	}

	response.MapIntoResponse(movie)

	return
}

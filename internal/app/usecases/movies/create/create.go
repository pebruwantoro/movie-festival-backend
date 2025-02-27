package create

import (
	"context"
	"github.com/pebruwantoro/movie-festival-backend/internal/app/entities"
)

func (u *Usecase) Execute(ctx context.Context, request CreateMovieRequest) (response CreateMovieResponse, err error) {
	movie := entities.Movie{}

	data := request.MapIntoMovie()

	movie, err = u.movieRepo.CreateMovie(ctx, data)
	if err != nil {
		return
	}
	response.UUID = movie.UUID

	return
}

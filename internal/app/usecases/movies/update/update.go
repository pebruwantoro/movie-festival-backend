package update

import (
	"context"
	"github.com/pebruwantoro/movie-festival-backend/internal/app/entities"
)

func (u *Usecase) Execute(ctx context.Context, request UpdateMovieRequest) (response UpdateMovieResponse, err error) {
	movie := entities.Movie{}

	data := request.MapIntoMovie()

	movie, err = u.movieRepo.UpdateMovie(ctx, data)
	if err != nil {
		return
	}
	response.MapIntoResponse(movie)

	return
}

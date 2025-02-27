package create

import (
	"context"
	"github.com/pebruwantoro/movie-festival-backend/internal/app/entities"
)

func (u *Usecase) Execute(ctx context.Context, request CreateMovieRequest) (response CreateMovieResponse, err error) {
	moview := entities.Movie{}

	data := request.MapIntoMovie()

	moview, err = u.movieRepo.CreateMovie(ctx, data)
	if err != nil {
		return
	}
	response.UUID = moview.UUID

	return
}

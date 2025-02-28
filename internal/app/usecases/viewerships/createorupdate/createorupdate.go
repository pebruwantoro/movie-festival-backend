package createorupdate

import (
	"context"
	"github.com/pebruwantoro/movie-festival-backend/internal/app/entities"
)

func (u *Usecase) Execute(ctx context.Context, request CreateOrUpdateViewershipRequest) (response CreateOrUpdateViewershipResponse, err error) {
	movie := entities.Movie{}
	viewership := entities.Viewership{}
	data := request.MapIntoViewership()

	movie, err = u.movieRepo.GetMovieByUUID(ctx, data.MovieUUID)
	if err != nil {
		return
	}

	if err = request.ValidateRequest(movie.Duration); err != nil {
		return
	}

	viewership, err = u.viewershipRepo.CreateOrUpdateViewership(ctx, data)
	if err != nil {
		return
	}

	response.MapIntoResponse(viewership)

	return
}

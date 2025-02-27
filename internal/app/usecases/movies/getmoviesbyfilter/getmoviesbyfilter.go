package getmoviesbyfilter

import (
	"context"
	"github.com/pebruwantoro/movie-festival-backend/internal/app/entities"
	"github.com/pebruwantoro/movie-festival-backend/internal/pkg/helper"
	"math"
)

func (u *Usecase) Execute(ctx context.Context, request GetMovieByFilterRequest) (response GetMovieByFilterResponse, err error) {
	movies := []entities.Movie{}
	total := int64(0)
	filter := request.MapIntoFilter()

	movies, err = u.movieRepo.GetMoviesByFilter(ctx, filter)
	if err != nil {
		return
	}

	total, err = u.movieRepo.CountTotalMoviesByFilter(ctx, filter)
	if err != nil {
		return
	}

	response.MapIntoResponse(movies, helper.PageInfo{
		CurrentPage: request.Pagination.Page,
		PerPage:     request.Pagination.PerPage,
		TotalData:   int(total),
		TotalPage:   int(math.Ceil(float64(total) / float64(request.Pagination.PerPage))),
	})

	return
}

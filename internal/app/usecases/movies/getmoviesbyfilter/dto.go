package getmoviesbyfilter

import (
	"github.com/lib/pq"
	"github.com/pebruwantoro/movie-festival-backend/internal/app/entities"
	"github.com/pebruwantoro/movie-festival-backend/internal/app/repositories/movies"
	"github.com/pebruwantoro/movie-festival-backend/internal/pkg/helper"
)

type GetMovieByFilterRequest struct {
	Title       string   `json:"title"`
	Description string   `json:"description"`
	Artists     []string `json:"artists"`
	Genres      []string `json:"genres"`
	Pagination  helper.Pagination
}

type GetMovieByFilterResponse struct {
	List       []MovieResponse `json:"list"`
	Pagination helper.PageInfo `json:"pagination"`
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

func (r *GetMovieByFilterRequest) MapIntoFilter() movies.Filter {
	return movies.Filter{
		Title:       r.Title,
		Description: r.Description,
		Artists:     r.Artists,
		Genres:      r.Genres,
		Pagination: helper.Pagination{
			Page:    r.Pagination.Page,
			PerPage: r.Pagination.PerPage,
		},
	}
}

func (r *GetMovieByFilterResponse) MapIntoResponse(movies []entities.Movie, pagination helper.PageInfo) {
	for _, movie := range movies {
		r.List = append(r.List, MovieResponse{
			UUID:        movie.UUID,
			Title:       movie.Title,
			Description: movie.Description,
			Duration:    movie.Duration,
			Artists:     movie.Artists,
			Genres:      movie.Genres,
			Url:         movie.Url,
		})
	}

	r.Pagination = pagination
}

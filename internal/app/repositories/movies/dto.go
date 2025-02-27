package movies

import "github.com/pebruwantoro/movie-festival-backend/internal/pkg/helper"

type Filter struct {
	Title       string
	Description string
	Artists     []string
	Genres      []string
	Pagination  helper.Pagination
}

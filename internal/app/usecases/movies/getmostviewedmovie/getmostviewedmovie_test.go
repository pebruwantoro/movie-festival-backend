package getmostviewedmovie

import (
	"context"
	"errors"
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/go-playground/assert/v2"
	"github.com/pebruwantoro/movie-festival-backend/internal/app/entities"
	movieMock "github.com/pebruwantoro/movie-festival-backend/internal/app/repositories/movies/mocks"
	"go.uber.org/mock/gomock"
	"testing"
)

type testMock struct {
	ctrl *gomock.Controller
	db   sqlmock.Sqlmock
	mm   *movieMock.MockRepositoryInterface
}

type testCase struct {
	name         string
	mock         func(mock testMock)
	ExpectedResp GetMostViewedMovieRequestResponse
	ExpectedErr  error
}

func TestExecute(t *testing.T) {
	exampleRes := entities.MovieWithViewership{
		Movie: entities.Movie{
			UUID:        "uuid",
			Title:       "example",
			Description: "example",
			Duration:    1000,
			Artists:     []string{"example"},
			Genres:      []string{"example"},
			Url:         "https://www.google.com",
		},
		TotalViewed: 1000,
	}

	cases := []testCase{
		{
			name: "success",
			mock: func(mock testMock) {
				mock.mm.EXPECT().GetMostViewedMovie(gomock.Any()).Return(exampleRes, nil)
			},
			ExpectedResp: GetMostViewedMovieRequestResponse{
				UUID:        "uuid",
				Title:       "example",
				Description: "example",
				Duration:    1000,
				Artists:     []string{"example"},
				Genres:      []string{"example"},
				Url:         "https://www.google.com",
				Viewed:      1000,
			},
			ExpectedErr: nil,
		},
		{
			name: "error",
			mock: func(mock testMock) {
				mock.mm.EXPECT().GetMostViewedMovie(gomock.Any()).Return(entities.MovieWithViewership{}, errors.New("got an error"))
			},
			ExpectedResp: GetMostViewedMovieRequestResponse{},
			ExpectedErr:  errors.New("got an error"),
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			ctx := context.Background()
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			_, mockDB, _ := sqlmock.New()

			tm := testMock{
				ctrl: ctrl,
				db:   mockDB,
				mm:   movieMock.NewMockRepositoryInterface(ctrl),
			}

			tc.mock(tm)

			usecase := NewUsecase(tm.mm)

			res, err := usecase.Execute(ctx)

			assert.Equal(t, tc.ExpectedResp, res)
			assert.Equal(t, tc.ExpectedErr, err)
		})
	}
}

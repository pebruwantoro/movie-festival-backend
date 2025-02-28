package getmostviewedmoviegenre

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
	ExpectedResp GetMostViewedMovieGenreResponse
	ExpectedErr  error
}

func TestExecute(t *testing.T) {
	exampleRes := entities.MovieGenreWithViewership{
		Genre: entities.Genre{
			UUID: "uuid",
			Name: "example",
		},
		TotalViewed: 1000,
	}

	cases := []testCase{
		{
			name: "success",
			mock: func(mock testMock) {
				mock.mm.EXPECT().GetMostViewedMovieGenre(gomock.Any()).Return(exampleRes, nil)
			},
			ExpectedResp: GetMostViewedMovieGenreResponse{
				UUID:   "uuid",
				Name:   "example",
				Viewed: 1000,
			},
			ExpectedErr: nil,
		},
		{
			name: "error",
			mock: func(mock testMock) {
				mock.mm.EXPECT().GetMostViewedMovieGenre(gomock.Any()).Return(entities.MovieGenreWithViewership{}, errors.New("got an error"))
			},
			ExpectedResp: GetMostViewedMovieGenreResponse{},
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

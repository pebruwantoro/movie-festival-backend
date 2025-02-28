package update

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
	name        string
	mock        func(mock testMock)
	Request     UpdateMovieRequest
	ExpectedErr error
}

func TestExecute(t *testing.T) {
	req := UpdateMovieRequest{
		UUID:        "uuid",
		Title:       "example",
		Description: "example",
		Duration:    1000,
		Artists:     []string{"example"},
		Genres:      []string{"example"},
		Url:         "https://www.google.com",
		UpdatedBy:   "doni@gmail.com",
	}

	cases := []testCase{
		{
			name: "success",
			mock: func(mock testMock) {
				mock.mm.EXPECT().UpdateMovie(gomock.Any(), gomock.Any()).Return(entities.Movie{}, nil)
			},
			Request:     req,
			ExpectedErr: nil,
		},
		{
			name: "error",
			mock: func(mock testMock) {
				mock.mm.EXPECT().UpdateMovie(gomock.Any(), gomock.Any()).Return(entities.Movie{}, errors.New("got an error"))
			},
			Request:     req,
			ExpectedErr: errors.New("got an error"),
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

			_, err := usecase.Execute(ctx, tc.Request)

			assert.Equal(t, tc.ExpectedErr, err)
		})
	}
}

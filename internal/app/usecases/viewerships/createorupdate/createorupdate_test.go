package createorupdate

import (
	"context"
	"errors"
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/go-playground/assert/v2"
	"github.com/pebruwantoro/movie-festival-backend/internal/app/entities"
	movieMock "github.com/pebruwantoro/movie-festival-backend/internal/app/repositories/movies/mocks"
	vieweshipsMock "github.com/pebruwantoro/movie-festival-backend/internal/app/repositories/viewerships/mocks"
	"github.com/pebruwantoro/movie-festival-backend/internal/pkg/helper"
	"go.uber.org/mock/gomock"
	"testing"
)

type testMock struct {
	ctrl *gomock.Controller
	db   sqlmock.Sqlmock
	mm   *movieMock.MockRepositoryInterface
	vm   *vieweshipsMock.MockRepositoryInterface
}

type testCase struct {
	name         string
	mock         func(mock testMock)
	request      CreateOrUpdateViewershipRequest
	expectedResp CreateOrUpdateViewershipResponse
	expectedErr  error
}

func TestExecute(t *testing.T) {
	req := CreateOrUpdateViewershipRequest{
		MovieUUID:        "movie-uuid",
		UserUUID:         "user-uuid",
		WatchingDuration: 1000,
		CreatedBy:        "doni@gmail.com",
	}
	data := req.MapIntoViewership()
	resp := CreateOrUpdateViewershipResponse{}
	resp.MapIntoResponse(data)

	cases := []testCase{
		{
			name: "success",
			mock: func(mock testMock) {
				mock.mm.EXPECT().GetMovieByUUID(gomock.Any(), data.MovieUUID).Return(entities.Movie{
					UUID:     "movie-uuid",
					Duration: 10000,
				}, nil)

				mock.vm.EXPECT().CreateOrUpdateViewership(gomock.Any(), gomock.Any()).Return(data, nil)
			},
			request:      req,
			expectedResp: resp,
			expectedErr:  nil,
		},
		{
			name: "error-get-movie",
			mock: func(mock testMock) {
				mock.mm.EXPECT().GetMovieByUUID(gomock.Any(), data.MovieUUID).Return(entities.Movie{}, errors.New("got an error"))
			},
			request:      req,
			expectedResp: CreateOrUpdateViewershipResponse{},
			expectedErr:  errors.New("got an error"),
		},
		{
			name: "error-validation",
			mock: func(mock testMock) {
				mock.mm.EXPECT().GetMovieByUUID(gomock.Any(), data.MovieUUID).Return(entities.Movie{
					UUID:     "movie-uuid",
					Duration: 200,
				}, nil)
			},
			request:      req,
			expectedResp: CreateOrUpdateViewershipResponse{},
			expectedErr:  helper.ErrorMovieDurationViewershipInvalid,
		},
		{
			name: "error-create-or-update-viewership",
			mock: func(mock testMock) {
				mock.mm.EXPECT().GetMovieByUUID(gomock.Any(), data.MovieUUID).Return(entities.Movie{
					UUID:     "movie-uuid",
					Duration: 10000,
				}, nil)

				mock.vm.EXPECT().CreateOrUpdateViewership(gomock.Any(), gomock.Any()).Return(entities.Viewership{}, errors.New("got an error"))
			},
			request:      req,
			expectedResp: CreateOrUpdateViewershipResponse{},
			expectedErr:  errors.New("got an error"),
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
				vm:   vieweshipsMock.NewMockRepositoryInterface(ctrl),
			}

			tc.mock(tm)

			usecase := NewUsecase(tm.vm, tm.mm)

			res, err := usecase.Execute(ctx, tc.request)

			assert.Equal(t, tc.expectedResp, res)
			assert.Equal(t, tc.expectedErr, err)
		})
	}
}

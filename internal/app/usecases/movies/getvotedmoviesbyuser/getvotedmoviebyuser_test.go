package getvotedmoviesbyuser

import (
	"context"
	"errors"
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/go-playground/assert/v2"
	"github.com/pebruwantoro/movie-festival-backend/internal/app/entities"
	movieMock "github.com/pebruwantoro/movie-festival-backend/internal/app/repositories/movies/mocks"
	votersMock "github.com/pebruwantoro/movie-festival-backend/internal/app/repositories/voters/mocks"
	"go.uber.org/mock/gomock"
	"testing"
)

type testMock struct {
	ctrl *gomock.Controller
	db   sqlmock.Sqlmock
	mm   *movieMock.MockRepositoryInterface
	vm   *votersMock.MockRepositoryInterface
}

type testCase struct {
	name         string
	mock         func(mock testMock)
	Request      GetVotedMovieByUserRequest
	ExpectedResp GetVotedMovieByUserResponse
	ExpectedErr  error
}

func TestExecute(t *testing.T) {
	voters := []entities.Voters{
		{
			UUID:      "uuid1",
			UserUUID:  "user-uuid1",
			MovieUUID: "movie-uuid1",
		},
		{
			UUID:      "uuid2",
			UserUUID:  "user-uuid2",
			MovieUUID: "movie-uuid2",
		},
		{
			UUID:      "uuid2",
			UserUUID:  "user-uuid1",
			MovieUUID: "movie-uuid2",
		},
		{
			UUID:      "uuid3",
			UserUUID:  "user-uuid1",
			MovieUUID: "movie-uuid3",
		},
	}

	movies := []entities.Movie{
		{
			UUID: "movie-uuid1",
		},
		{
			UUID: "movie-uuid2",
		},
		{
			UUID: "movie-uuid2",
		},
	}

	cases := []testCase{
		{
			name: "success",
			mock: func(mock testMock) {
				mock.vm.EXPECT().GetVotersByUserUUID(gomock.Any(), gomock.Any()).Return(voters, nil)
				mock.mm.EXPECT().GetMovieByUUIDs(gomock.Any(), []string{"movie-uuid1", "movie-uuid2", "movie-uuid3"}).Return(movies, nil)
			},
			Request: GetVotedMovieByUserRequest{UserUUID: "uuid"},
			ExpectedResp: GetVotedMovieByUserResponse{
				Data: []MovieResponse{
					{
						UUID: "movie-uuid1",
					},
					{
						UUID: "movie-uuid2",
					},
					{
						UUID: "movie-uuid2",
					},
				},
			},
			ExpectedErr: nil,
		},
		{
			name: "error-get-voters",
			mock: func(mock testMock) {
				mock.vm.EXPECT().GetVotersByUserUUID(gomock.Any(), gomock.Any()).Return(nil, errors.New("got an error"))
			},
			Request:      GetVotedMovieByUserRequest{UserUUID: "uuid"},
			ExpectedResp: GetVotedMovieByUserResponse{},
			ExpectedErr:  errors.New("got an error"),
		},
		{
			name: "error-get-movie",
			mock: func(mock testMock) {
				mock.vm.EXPECT().GetVotersByUserUUID(gomock.Any(), gomock.Any()).Return(voters, nil)
				mock.mm.EXPECT().GetMovieByUUIDs(gomock.Any(), []string{"movie-uuid1", "movie-uuid2", "movie-uuid3"}).Return(nil, errors.New("got an error"))
			},
			Request:      GetVotedMovieByUserRequest{UserUUID: "uuid"},
			ExpectedResp: GetVotedMovieByUserResponse{},
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
				vm:   votersMock.NewMockRepositoryInterface(ctrl),
			}

			tc.mock(tm)

			usecase := NewUsecase(tm.vm, tm.mm)

			res, err := usecase.Execute(ctx, tc.Request)

			assert.Equal(t, tc.ExpectedResp, res)
			assert.Equal(t, tc.ExpectedErr, err)
		})
	}
}

package unvote

import (
	"context"
	"errors"
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/go-playground/assert/v2"
	"github.com/pebruwantoro/movie-festival-backend/internal/app/entities"
	votersMock "github.com/pebruwantoro/movie-festival-backend/internal/app/repositories/voters/mocks"
	"go.uber.org/mock/gomock"
	"testing"
)

type testMock struct {
	ctrl *gomock.Controller
	db   sqlmock.Sqlmock
	vm   *votersMock.MockRepositoryInterface
}

type testCase struct {
	name        string
	mock        func(mock testMock)
	Request     UnVoteMovieRequest
	ExpectedErr error
}

func TestExecute(t *testing.T) {
	req := UnVoteMovieRequest{
		UserUUID:  "uuid-user",
		UUID:      "uuid",
		DeletedBy: "doni@gmail.com",
	}

	cases := []testCase{
		{
			name: "success",
			mock: func(mock testMock) {
				mock.vm.EXPECT().DeleteVoter(gomock.Any(), gomock.Any()).Return(entities.Voters{}, nil)
			},
			Request:     req,
			ExpectedErr: nil,
		},
		{
			name: "error",
			mock: func(mock testMock) {
				mock.vm.EXPECT().DeleteVoter(gomock.Any(), gomock.Any()).Return(entities.Voters{}, errors.New("got an error"))
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
				vm:   votersMock.NewMockRepositoryInterface(ctrl),
			}

			tc.mock(tm)

			usecase := NewUsecase(tm.vm)

			_, err := usecase.Execute(ctx, tc.Request)

			assert.Equal(t, tc.ExpectedErr, err)
		})
	}
}

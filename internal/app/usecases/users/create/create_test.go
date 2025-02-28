package create

import (
	"context"
	"errors"
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/go-playground/assert/v2"
	"github.com/pebruwantoro/movie-festival-backend/constants"
	"github.com/pebruwantoro/movie-festival-backend/internal/app/entities"
	userMock "github.com/pebruwantoro/movie-festival-backend/internal/app/repositories/users/mocks"
	"go.uber.org/mock/gomock"
	"testing"
)

type testMock struct {
	ctrl *gomock.Controller
	db   sqlmock.Sqlmock
	um   *userMock.MockRepositoryInterface
}

type testCase struct {
	name        string
	mock        func(mock testMock)
	Request     CreateUserRequest
	ExpectedErr error
}

func TestExecute(t *testing.T) {
	req := CreateUserRequest{
		Name:     "Doni",
		Email:    "doni@gmail.com",
		Password: "1234*Doni",
		Role:     constants.ROLE_USER,
	}

	cases := []testCase{
		{
			name: "success",
			mock: func(mock testMock) {
				mock.um.EXPECT().CreateUser(gomock.Any(), gomock.Any()).Return(entities.User{}, nil)
			},
			Request:     req,
			ExpectedErr: nil,
		},
		{
			name: "error",
			mock: func(mock testMock) {
				mock.um.EXPECT().CreateUser(gomock.Any(), gomock.Any()).Return(entities.User{}, errors.New("got an error"))
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
				um:   userMock.NewMockRepositoryInterface(ctrl),
			}

			tc.mock(tm)

			usecase := NewUsecase(tm.um)

			_, err := usecase.Execute(ctx, tc.Request)

			assert.Equal(t, tc.ExpectedErr, err)
		})
	}
}

package login

import (
	"context"
	"errors"
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/go-playground/assert/v2"
	"github.com/pebruwantoro/movie-festival-backend/constants"
	"github.com/pebruwantoro/movie-festival-backend/internal/app/entities"
	tokenMock "github.com/pebruwantoro/movie-festival-backend/internal/app/repositories/token/mocks"
	userMock "github.com/pebruwantoro/movie-festival-backend/internal/app/repositories/users/mocks"
	"github.com/pebruwantoro/movie-festival-backend/internal/pkg/helper"
	"go.uber.org/mock/gomock"
	"testing"
)

type testMock struct {
	ctrl *gomock.Controller
	db   sqlmock.Sqlmock
	um   *userMock.MockRepositoryInterface
	tkm  *tokenMock.MockRepositoryInterface
}

type testCase struct {
	name        string
	mock        func(mock testMock)
	Request     LoginRequest
	ExpectedRes LoginResponse
	ExpectedErr error
}

func TestExecute(t *testing.T) {
	req := LoginRequest{
		Email:    "doni@gmail.com",
		Password: "1234*Doni",
	}
	hashPassword, _ := helper.HashPassword(req.Password)
	user := entities.User{
		UUID:     "uuid",
		Name:     "Doni",
		Email:    "doni@gmail.com",
		Password: hashPassword,
		Role:     constants.ROLE_USER,
	}

	token, _ := helper.GenerateJWT(user)

	cases := []testCase{
		{
			name: "success",
			mock: func(mock testMock) {
				mock.um.EXPECT().GetUserByEmail(gomock.Any(), "doni@gmail.com").Return(user, nil)
				mock.tkm.EXPECT().CreateToken(gomock.Any(), gomock.Any()).Return(entities.Token{}, nil)
			},
			Request:     req,
			ExpectedRes: LoginResponse{Token: token},
			ExpectedErr: nil,
		},
		{
			name: "error-get-user",
			mock: func(mock testMock) {
				mock.um.EXPECT().GetUserByEmail(gomock.Any(), "doni@gmail.com").Return(entities.User{}, errors.New("got an error"))
			},
			Request:     req,
			ExpectedErr: errors.New("got an error"),
		},

		{
			name: "error-create-token",
			mock: func(mock testMock) {
				mock.um.EXPECT().GetUserByEmail(gomock.Any(), "doni@gmail.com").Return(user, nil)
				mock.tkm.EXPECT().CreateToken(gomock.Any(), gomock.Any()).Return(entities.Token{}, errors.New("got an error"))
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
				tkm:  tokenMock.NewMockRepositoryInterface(ctrl),
			}

			tc.mock(tm)

			usecase := NewUsecase(tm.um, tm.tkm)

			resp, err := usecase.Execute(ctx, tc.Request)

			assert.Equal(t, tc.ExpectedRes, resp)
			assert.Equal(t, tc.ExpectedErr, err)
		})
	}
}

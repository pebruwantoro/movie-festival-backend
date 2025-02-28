package logout

import (
	"context"
	"errors"
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/go-playground/assert/v2"
	"github.com/pebruwantoro/movie-festival-backend/constants"
	"github.com/pebruwantoro/movie-festival-backend/internal/app/entities"
	tokenMock "github.com/pebruwantoro/movie-festival-backend/internal/app/repositories/token/mocks"
	"github.com/pebruwantoro/movie-festival-backend/internal/pkg/helper"
	"go.uber.org/mock/gomock"
	"gorm.io/gorm"
	"testing"
)

type testMock struct {
	ctrl *gomock.Controller
	db   sqlmock.Sqlmock
	tkm  *tokenMock.MockRepositoryInterface
}

type testCase struct {
	name        string
	mock        func(mock testMock)
	Request     LogoutRequest
	ExpectedErr error
}

func TestExecute(t *testing.T) {
	hashPassword, _ := helper.HashPassword("1234Doni*")
	user := entities.User{
		UUID:     "uuid",
		Name:     "Doni",
		Email:    "doni@gmail.com",
		Password: hashPassword,
		Role:     constants.ROLE_USER,
	}

	token, _ := helper.GenerateJWT(user)
	req := LogoutRequest{
		Token: token,
	}

	cases := []testCase{
		{
			name: "success",
			mock: func(mock testMock) {
				mock.tkm.EXPECT().GetActiveTokenByJWT(gomock.Any(), req.Token).Return(entities.Token{
					Token:    token,
					IsActive: true,
				}, nil)

				mock.tkm.EXPECT().UpdateToken(gomock.Any(), gomock.Any()).Return(entities.Token{}, nil)
			},
			Request:     req,
			ExpectedErr: nil,
		},
		{
			name: "error-user-already-logout",
			mock: func(mock testMock) {
				mock.tkm.EXPECT().GetActiveTokenByJWT(gomock.Any(), req.Token).Return(entities.Token{}, gorm.ErrRecordNotFound)
			},
			Request:     req,
			ExpectedErr: helper.ErrorUserAlreadyLogout,
		},

		{
			name: "error-get-active-token",
			mock: func(mock testMock) {
				mock.tkm.EXPECT().GetActiveTokenByJWT(gomock.Any(), req.Token).Return(entities.Token{}, errors.New("got an error"))
			},
			Request:     req,
			ExpectedErr: errors.New("got an error"),
		},
		{
			name: "error-set-inactive-token",
			mock: func(mock testMock) {
				mock.tkm.EXPECT().GetActiveTokenByJWT(gomock.Any(), req.Token).Return(entities.Token{
					Token:    token,
					IsActive: true,
				}, nil)

				mock.tkm.EXPECT().UpdateToken(gomock.Any(), gomock.Any()).Return(entities.Token{}, errors.New("got an error"))
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
				tkm:  tokenMock.NewMockRepositoryInterface(ctrl),
			}

			tc.mock(tm)

			usecase := NewUsecase(tm.tkm)

			_, err := usecase.Execute(ctx, tc.Request)

			assert.Equal(t, tc.ExpectedErr, err)
		})
	}
}

package usecases

import (
	"github.com/pebruwantoro/movie-festival-backend/internal/app/repositories/token"
	"github.com/pebruwantoro/movie-festival-backend/internal/app/repositories/users"
	createUser "github.com/pebruwantoro/movie-festival-backend/internal/app/usecases/users/create"
	loginUser "github.com/pebruwantoro/movie-festival-backend/internal/app/usecases/users/login"
	"github.com/pebruwantoro/movie-festival-backend/internal/app/usecases/users/logout"
	"gorm.io/gorm"
)

type Container struct {
	CreateUserUseacse createUser.Usecase
	LoginUserUsecase  loginUser.Usecase
	LogoutUsecase     logout.Usecase
}

func NewContainer(db *gorm.DB) *Container {
	userRepo := *users.NewRepository(db)
	tokenRepo := *token.NewRepository(db)

	return &Container{
		CreateUserUseacse: *createUser.NewUsecase(userRepo),
		LoginUserUsecase:  *loginUser.NewUsecase(userRepo, tokenRepo),
		LogoutUsecase:     *logout.NewUsecase(tokenRepo),
	}
}

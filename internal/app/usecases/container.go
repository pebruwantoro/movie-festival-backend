package usecases

import (
	"github.com/pebruwantoro/movie-festival-backend/internal/app/repositories/movies"
	"github.com/pebruwantoro/movie-festival-backend/internal/app/repositories/token"
	"github.com/pebruwantoro/movie-festival-backend/internal/app/repositories/users"
	createMovie "github.com/pebruwantoro/movie-festival-backend/internal/app/usecases/movies/create"
	updateMovie "github.com/pebruwantoro/movie-festival-backend/internal/app/usecases/movies/update"
	createUser "github.com/pebruwantoro/movie-festival-backend/internal/app/usecases/users/create"
	loginUser "github.com/pebruwantoro/movie-festival-backend/internal/app/usecases/users/login"
	"github.com/pebruwantoro/movie-festival-backend/internal/app/usecases/users/logout"
	"gorm.io/gorm"
)

type Container struct {
	CreateUserUseacse createUser.Usecase
	LoginUserUsecase  loginUser.Usecase
	LogoutUsecase     logout.Usecase
	CreateMovie       createMovie.Usecase
	UpdateMovie       updateMovie.Usecase
}

func NewContainer(db *gorm.DB) *Container {
	userRepo := *users.NewRepository(db)
	tokenRepo := *token.NewRepository(db)
	movieRepo := *movies.NewRepository(db)

	return &Container{
		CreateUserUseacse: *createUser.NewUsecase(userRepo),
		LoginUserUsecase:  *loginUser.NewUsecase(userRepo, tokenRepo),
		LogoutUsecase:     *logout.NewUsecase(tokenRepo),
		CreateMovie:       *createMovie.NewUsecase(movieRepo),
		UpdateMovie:       *updateMovie.NewUsecase(movieRepo),
	}
}

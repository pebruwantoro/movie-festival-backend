package usecases

import (
	"github.com/pebruwantoro/movie-festival-backend/internal/app/repositories/movies"
	"github.com/pebruwantoro/movie-festival-backend/internal/app/repositories/token"
	"github.com/pebruwantoro/movie-festival-backend/internal/app/repositories/users"
	"github.com/pebruwantoro/movie-festival-backend/internal/app/repositories/viewerships"
	"github.com/pebruwantoro/movie-festival-backend/internal/app/repositories/voters"
	createMovie "github.com/pebruwantoro/movie-festival-backend/internal/app/usecases/movies/create"
	getMostViewedMovie "github.com/pebruwantoro/movie-festival-backend/internal/app/usecases/movies/getmostviewedmovie"
	getMostViewedMovieGenre "github.com/pebruwantoro/movie-festival-backend/internal/app/usecases/movies/getmostviewedmoviegenre"
	getMoviesByFilter "github.com/pebruwantoro/movie-festival-backend/internal/app/usecases/movies/getmoviesbyfilter"
	getVotedMoviesByUser "github.com/pebruwantoro/movie-festival-backend/internal/app/usecases/movies/getvotedmoviesbyuser"
	unvoteMovie "github.com/pebruwantoro/movie-festival-backend/internal/app/usecases/movies/unvote"
	updateMovie "github.com/pebruwantoro/movie-festival-backend/internal/app/usecases/movies/update"
	voteMovie "github.com/pebruwantoro/movie-festival-backend/internal/app/usecases/movies/vote"
	createUser "github.com/pebruwantoro/movie-festival-backend/internal/app/usecases/users/create"
	loginUser "github.com/pebruwantoro/movie-festival-backend/internal/app/usecases/users/login"
	"github.com/pebruwantoro/movie-festival-backend/internal/app/usecases/users/logout"
	createOrUpdateViewership "github.com/pebruwantoro/movie-festival-backend/internal/app/usecases/viewerships/createorupdate"
	"gorm.io/gorm"
)

type Container struct {
	CreateUserUseacse        createUser.Usecase
	LoginUserUsecase         loginUser.Usecase
	LogoutUsecase            logout.Usecase
	CreateMovie              createMovie.Usecase
	UpdateMovie              updateMovie.Usecase
	VoteMovie                voteMovie.Usecase
	UnVoteMovie              unvoteMovie.Usecase
	GetVotedMoviesByUser     getVotedMoviesByUser.Usecase
	GetMoviesByFilter        getMoviesByFilter.Usecase
	CreateOrUpdateViewership createOrUpdateViewership.Usecase
	GetMostViewedMovie       getMostViewedMovie.Usecase
	GetMostViewedMovieGenre  getMostViewedMovieGenre.Usecase
}

func NewContainer(db *gorm.DB) *Container {
	userRepo := users.NewRepository(db)
	tokenRepo := token.NewRepository(db)
	movieRepo := movies.NewRepository(db)
	voteRepo := voters.NewRepository(db)
	viewershipRepo := viewerships.NewRepository(db)

	return &Container{
		CreateUserUseacse:        *createUser.NewUsecase(userRepo),
		LoginUserUsecase:         *loginUser.NewUsecase(userRepo, tokenRepo),
		LogoutUsecase:            *logout.NewUsecase(tokenRepo),
		CreateMovie:              *createMovie.NewUsecase(movieRepo),
		UpdateMovie:              *updateMovie.NewUsecase(movieRepo),
		VoteMovie:                *voteMovie.NewUsecase(voteRepo),
		UnVoteMovie:              *unvoteMovie.NewUsecase(voteRepo),
		GetVotedMoviesByUser:     *getVotedMoviesByUser.NewUsecase(voteRepo, movieRepo),
		GetMoviesByFilter:        *getMoviesByFilter.NewUsecase(movieRepo),
		CreateOrUpdateViewership: *createOrUpdateViewership.NewUsecase(viewershipRepo, movieRepo),
		GetMostViewedMovie:       *getMostViewedMovie.NewUsecase(movieRepo),
		GetMostViewedMovieGenre:  *getMostViewedMovieGenre.NewUsecase(movieRepo),
	}
}

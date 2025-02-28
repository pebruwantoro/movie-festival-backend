package handlers

import (
	"context"
	"github.com/labstack/echo/v4"
	_ "github.com/pebruwantoro/movie-festival-backend/docs"
	echoSwagger "github.com/swaggo/echo-swagger"
)

type Router struct {
	ctx    context.Context
	echo   *echo.Echo
	server *Server
}

func NewRouter(ctx context.Context, echo *echo.Echo, server *Server) *Router {
	return &Router{
		ctx:    ctx,
		echo:   echo,
		server: server,
	}
}

func (r *Router) RegisterRouter() {

	r.echo.GET("/docs/swagger/*", echoSwagger.WrapHandler)

	users := r.echo.Group("users")
	users.POST("/sign-up", r.server.UserSignUpHandler)
	users.POST("/sign-up/admin", r.server.AdminSignUpHandler)
	users.POST("/login", r.server.LoginUserHandler)
	users.POST("/logout", r.server.LogoutUserHandler, AuthenticationMiddleware())

	movies := r.echo.Group("movies")
	movies.POST("", r.server.CreateMovieHandler, AuthenticationMiddleware(), AuthorizationAdminMiddleware())
	movies.POST("/upload", r.server.UploadMovieHandler, AuthenticationMiddleware(), AuthorizationAdminMiddleware())
	movies.PUT("/:uuid", r.server.UpdateMovieHandler, AuthenticationMiddleware(), AuthorizationAdminMiddleware())
	movies.POST("/vote", r.server.VoteMovieHandler, AuthenticationMiddleware(), AuthorizationUserMiddleware())
	movies.DELETE("/vote/:uuid", r.server.UnVoteMovieHandler, AuthenticationMiddleware(), AuthorizationUserMiddleware())
	movies.GET("/votes/list", r.server.GetVotesListHandler, AuthenticationMiddleware(), AuthorizationUserMiddleware())
	movies.GET("/list", r.server.SearchMoviesByFilterHandler, AuthenticationMiddleware(), AuthorizationUserMiddleware())
	movies.POST("/track", r.server.TrackMovieViewershipHandler, AuthenticationMiddleware(), AuthorizationUserMiddleware())
	movies.GET("/most-viewed", r.server.GetMostViewedMovie, AuthenticationMiddleware(), AuthorizationAdminMiddleware())
	movies.GET("/genres/most-viewed", r.server.GetMostViewedGenre, AuthenticationMiddleware(), AuthorizationAdminMiddleware())
}

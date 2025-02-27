package handlers

import (
	"context"
	"github.com/labstack/echo/v4"
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

	users := r.echo.Group("users")
	r.echo.Group("users")
	users.POST("/sign-up", r.server.UserSignUpHandler)
	users.POST("/sign-up/admin", r.server.AdminSignUpHandler)
	users.POST("/login", r.server.LoginUserHandler)
	users.POST("/logout", r.server.LogoutUserHandler, AuthenticationMiddleware())

	movies := r.echo.Group("movies")
	movies.POST("/", r.server.CreateMovieHandler, AuthenticationMiddleware(), AuthorizationAdminMiddleware())
	movies.POST("/upload", r.server.UploadMovieHandler, AuthenticationMiddleware(), AuthorizationAdminMiddleware())
	movies.PUT("/:uuid", r.server.UpdateMovieHandler, AuthenticationMiddleware(), AuthorizationAdminMiddleware())
}

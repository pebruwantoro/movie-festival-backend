package handlers

import (
	"github.com/pebruwantoro/movie-festival-backend/internal/app/usecases"
)

type Server struct {
	Usecase usecases.Container
}

func NewServer(opts usecases.Container) *Server {
	return &Server{
		Usecase: opts,
	}
}

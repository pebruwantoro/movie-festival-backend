package cmd

import (
	"context"
	"fmt"
	"github.com/labstack/echo/v4"
	"github.com/pebruwantoro/movie-festival-backend/config"
	"github.com/pebruwantoro/movie-festival-backend/internal/app/handlers"
	"github.com/pebruwantoro/movie-festival-backend/internal/app/usecases"
	"github.com/pebruwantoro/movie-festival-backend/internal/pkg/helper"
	"time"
)

func init() {
	//err := godotenv.Load()
	//if err != nil {
	//	log.Fatalf(err.Error())
	//}

	config.Load()
}

func Execute() {
	app := echo.New()
	app.Server.WriteTimeout = time.Duration(config.Get().ServerWriteTimeout) * time.Second
	app.Server.ReadTimeout = time.Duration(config.Get().ServerReadTimeout) * time.Second

	server := newServer()

	helper.InitValidator()
	handlers.NewRouter(context.Background(), app, server).RegisterRouter()
	handlers.SetupMiddleware(app)

	app.Logger.Fatal(app.Start(fmt.Sprintf(":%s", config.Get().AppPort)))
}

func newServer() *handlers.Server {
	database := config.NewDatabase("postgres")
	container := usecases.NewContainer(database)

	return handlers.NewServer(*container)
}

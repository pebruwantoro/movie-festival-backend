package cmd

import (
	"context"
	"fmt"
	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"github.com/pebruwantoro/movie-festival-backend/config"
	"github.com/pebruwantoro/movie-festival-backend/internal/app/handlers"
	"github.com/pebruwantoro/movie-festival-backend/internal/app/usecases"
	"github.com/pebruwantoro/movie-festival-backend/internal/pkg/helper"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf(err.Error())
	}

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

	go func() {
		if err := app.Start(fmt.Sprintf(":%s", config.Get().AppPort)); err != nil {
			app.Logger.Info("shutting down the server")
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	if err := app.Shutdown(ctx); err != nil {
		app.Logger.Fatal(err)
	}
}

func newServer() *handlers.Server {
	database := config.NewDatabase("postgres")
	container := usecases.NewContainer(database)

	return handlers.NewServer(*container)
}

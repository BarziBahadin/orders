package application

import (
	"context"
	"fmt"
	"net/http"

	"github.com/redis/go-redis/v9"
)

type App struct {
	router http.Handler
	rdb    *redis.Client
}

func New() *App {
	app := &App{
		rdb: redis.NewClient(&redis.Options{}),
	}

	app.loadRoutes()
	return app
}

func (a *App) Start(ctx context.Context) error {
	server := &http.Server{
		Addr:    ":3004",
		Handler: a.router,
	}

	err := a.rdb.Ping(ctx).Err()
	if err != nil {
		fmt.Println("failed to connect to redis:  %w", err)
	}
	fmt.Println("starting sever")
	err = server.ListenAndServe()
	if err != nil {
		fmt.Println("failed to listen to server: %w", err)
	}
	return nil
}

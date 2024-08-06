package main

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"os"
	"os/signal"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/joho/godotenv"
	"github.com/leobritto95/semana-tech-go-react-server.git/internal/api"
	"github.com/leobritto95/semana-tech-go-react-server.git/internal/store/pgstore"
)

func main() {
	if err := godotenv.Load(); err != nil {
		panic(err)
	}

	ctx := context.Background()

	poll, err:= pgxpool.New(ctx, fmt.Sprintf(
		"user=%s password=%s host=%s port=%s dbname=%s",
		os.Getenv("WSRS_DATABASE_USER"),
		os.Getenv("WSRS_DATABASE_PASSWORD"),
		os.Getenv("WSRS_DATABASE_HOST"),
		os.Getenv("WSRS_DATABASE_PORT"),
		os.Getenv("WSRS_DATABASE_NAME"),
		))
		if err!= nil {
      panic(err)
    }

		defer poll.Close()

		if err := poll.Ping(ctx); err != nil {
			panic(err)
		}

		handler := api.NewHandler(pgstore.New(poll))

		go func() {
				if err := http.ListenAndServe(":8080", handler); err != nil {
					if !errors.Is(err, http.ErrServerClosed) {
						panic(err)
				}
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)
	<-quit
}
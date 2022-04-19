package main

import (
	"net/http"
	"os"
	"time"

	"github.com/IceWreck/BookStack2Site/config"
	"github.com/rs/zerolog"
)

func main() {
	app := &config.Application{
		Logger: zerolog.New(
			zerolog.ConsoleWriter{
				Out:        os.Stdout,
				TimeFormat: time.RFC822,
			},
		).With().Timestamp().Logger(),
	}

	zerolog.SetGlobalLevel(zerolog.InfoLevel)

	app.Config = config.LoadConfig(app)

	if app.Config.BookStackEndpoint == "" || app.Config.BookStackAPITokenID == "" || app.Config.BookStackAPITokenSecret == "" {
		app.Logger.Fatal().Msg("BookStackEndpoint, BookStackAPITokenID, BookStackAPITokenSecret cannot be empty")
	}

	app.Client = &http.Client{
		Timeout: 120 * time.Second,
	}

}

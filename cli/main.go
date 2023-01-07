package main

import (
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/IceWreck/BookStack2Site/downloader"

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

	app.Config = config.LoadConfig(app)
	app.Logger.Info().Str("config", fmt.Sprint(app.Config)).Msg("")

	if app.Config.VerboseLogs {
		zerolog.SetGlobalLevel(zerolog.DebugLevel)
	} else {
		zerolog.SetGlobalLevel(zerolog.InfoLevel)
	}

	if app.Config.BookStackEndpoint == "" || app.Config.BookStackAPITokenID == "" || app.Config.BookStackAPITokenSecret == "" {
		app.Logger.Fatal().Msg("BookStackEndpoint, BookStackAPITokenID, BookStackAPITokenSecret cannot be empty")
	}

	app.Client = &http.Client{
		Timeout: 120 * time.Second,
	}

	app.Logger.Info().Msg("Trying to establish wiki structure. This might take a while.")
	downloader.Download(app)
}

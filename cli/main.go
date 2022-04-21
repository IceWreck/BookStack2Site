package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/IceWreck/BookStack2Site/bookstackclient"

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

	zerolog.SetGlobalLevel(zerolog.DebugLevel)

	app.Config = config.LoadConfig(app)
	app.Logger.Info().Str("config", fmt.Sprint(app.Config)).Msg("")

	if app.Config.BookStackEndpoint == "" || app.Config.BookStackAPITokenID == "" || app.Config.BookStackAPITokenSecret == "" {
		app.Logger.Fatal().Msg("BookStackEndpoint, BookStackAPITokenID, BookStackAPITokenSecret cannot be empty")
	}

	app.Client = &http.Client{
		Timeout: 120 * time.Second,
	}

	//fmt.Println(bookstackclient.FetchBooks(app))
	//fmt.Println(bookstackclient.FetchChapters(app, 1))
	//fmt.Println(bookstackclient.FetchPages(app, 10, 0))

	w, _ := bookstackclient.FetchWiki(app)
	jsonWiki, _ := json.Marshal(w)
	fmt.Println(string(jsonWiki))
}

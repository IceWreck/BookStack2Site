package config

import (
	"net/http"

	"github.com/rs/zerolog"
)

// Application struct to hold the dependencies for our application.
type Application struct {
	Config Config
	Logger zerolog.Logger
	Client *http.Client
}

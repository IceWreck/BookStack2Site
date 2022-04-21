package config

import (
	"flag"
)

const Version = "1.0.0"

// Config struct to hold all the configuration settings for our application.
type Config struct {
	BookStackEndpoint       string // Ex:  https://wiki.example.com
	BookStackAPITokenID     string // token ID
	BookStackAPITokenSecret string // token secret
	GenerateHTML            bool   // generate static site HTML
	MdBookLocation          string // path to the mdBook binary
	Concurrency             int    // number of concurrent goroutines
}

func LoadConfig(app *Application) Config {
	var settings Config

	// Command line flags and their default values

	// Required
	flag.StringVar(&settings.BookStackEndpoint, "bookstack-url", "", "BookStack Endpoint")
	flag.StringVar(&settings.BookStackAPITokenID, "token-id", "", "BookStack API Token ID")
	flag.StringVar(&settings.BookStackAPITokenSecret, "token-secret", "", "BookStack API Token Secret")

	// Optional
	flag.IntVar(&settings.Concurrency, "concurrency", 10, "Number of concurrent page downloads")
	flag.BoolVar(&settings.GenerateHTML, "generate-html", true, "Generate Static Site HTML or just markdown")
	flag.StringVar(&settings.MdBookLocation, "mdbook-location", "mdbook", "Custom path of mdbook binary")

	flag.Parse()
	return settings
}

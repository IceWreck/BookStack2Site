package config

import "flag"

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
	settings.BookStackEndpoint = *flag.String("bookstack-url", "", "BookStack Endpoint")
	settings.BookStackAPITokenID = *flag.String("token-id", "", "BookStack API Token ID")
	settings.BookStackAPITokenSecret = *flag.String("token-secret", "", "BookStack API Token Secret")

	// Optional
	settings.Concurrency = *flag.Int("concurrency", 10, "Number of concurrent page downloads")
	settings.GenerateHTML = *flag.Bool("generate-html", true, "Generate Static Site HTML or just markdown")
	settings.MdBookLocation = *flag.String("mdbook-location", "mdbook", "Custom path of mdbook binary")

	flag.Parse()

	return settings
}

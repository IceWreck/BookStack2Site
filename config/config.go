package config

import (
	"flag"
	"path"
)

const Version = "0.1.1"

// Config struct to hold all the configuration settings for our application.
type Config struct {
	BookStackEndpoint       string // Ex:  https://wiki.example.com
	BookStackAPITokenID     string // token ID
	BookStackAPITokenSecret string // token secret
	GenerateHTML            bool   // generate static site HTML
	MdBookLocation          string // path to the mdBook binary
	Concurrency             int    // number of concurrent goroutines
	DownloadLocation        string // path of downloaded markdown
	VerboseLogs             bool   // print detailed logs
	DownloadImages          bool   // download static images from the wiki as well
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
	// flag.BoolVar(&settings.GenerateHTML, "generate-html", true, "Generate Static Site HTML or just markdown")
	// flag.StringVar(&settings.MdBookLocation, "mdbook-location", "mdbook", "Custom path of mdbook binary")
	flag.StringVar(&settings.DownloadLocation, "download-location", "./book", "Path of downloaded markdown")
	flag.BoolVar(&settings.VerboseLogs, "verbose", false, "Print detailed logs")
	// flag.BoolVar(&settings.DownloadImages, "download-images", true, "Download static images from the wiki")

	flag.Parse()

	// Clean up file paths.
	settings.MdBookLocation = path.Clean(settings.MdBookLocation)
	settings.DownloadLocation = path.Clean(settings.DownloadLocation)

	return settings
}

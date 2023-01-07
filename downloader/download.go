package downloader

import (
	"fmt"
	"os"
	"path"
	"path/filepath"
	"sync"

	"github.com/IceWreck/BookStack2Site/bookstackclient"
	"github.com/IceWreck/BookStack2Site/config"
)

func Download(app *config.Application) {

	w, _ := bookstackclient.FetchWiki(app)
	sem := make(chan struct{}, app.Config.Concurrency)

	var wg sync.WaitGroup

	for _, book := range w.Books {
		for _, chapter := range book.Chapters {
			for _, page := range chapter.Pages {
				wg.Add(1)
				sem <- struct{}{}
				page.FilePath = fmt.Sprint("/", book.Slug, "/", chapter.Slug, "/", page.Slug)
				go func(p bookstackclient.WikiPage) {
					defer wg.Done()
					downloadPage(app, p)
					// release semaphore
					<-sem
				}(page)

			}
		}
		for _, indiePage := range book.IndiePages {

			wg.Add(1)
			sem <- struct{}{}
			indiePage.FilePath = fmt.Sprint("/", book.Slug, "/", indiePage.Slug)
			go func(p bookstackclient.WikiPage) {
				defer wg.Done()
				downloadPage(app, p)
				// release semaphore
				<-sem
			}(indiePage)
		}
	}

	wg.Wait()

	// Create the Summary.md file

	summaryContents := "# Summary\n"
	for _, book := range w.Books {
		summaryContents += "\n# " + book.Name + "\n\n"
		for _, chapter := range book.Chapters {
			summaryContents += fmt.Sprint("- [", chapter.Name, "]()\n")

			for _, page := range chapter.Pages {
				summaryContents += fmt.Sprint("    - [", page.Name, "](", book.Slug, "/", chapter.Slug, "/", page.Slug, ".md)\n")

			}
		}
		for _, indiePage := range book.IndiePages {
			summaryContents += fmt.Sprint("- [", indiePage.Name, "](", book.Slug, "/", indiePage.Slug, ".md)\n")
		}
	}

	file, err := createFile(path.Clean(fmt.Sprint(app.Config.DownloadLocation, "/SUMMARY.md")))
	if err != nil {
		app.Logger.Error().Err(err).Msg("Error creating SUMMARY.md file")
	}

	_, err = file.Write([]byte(summaryContents))
	if err != nil {
		app.Logger.Error().Err(err).Str("page", "SUMMARY.md").Msg("Error writing to file")
	} else {
		app.Logger.Info().Msg("Written SUMMARY.md")
	}

}

func downloadPage(app *config.Application, page bookstackclient.WikiPage) {
	app.Logger.Info().Str("page", page.Name).Msg("Downloading Page")
	markdownBytes, err := bookstackclient.FetchPageMarkdown(app, page.PageID)
	if err != nil {
		app.Logger.Error().Err(err).Str("page", page.Name).Msg("Error downloading page")
		return
	}
	fileLocation := path.Clean(fmt.Sprint(app.Config.DownloadLocation, "/", page.FilePath, ".md"))
	file, err := createFile(fileLocation)
	if err != nil {
		app.Logger.Error().Str("fileLocation", fileLocation).Err(err).Str("page", page.Name).Msg("Error creating file")
	}
	defer func() {
		if err = file.Close(); err != nil {
			app.Logger.Error().Err(err).Str("page", page.Name).Msg("Error closing file")
		}
	}()
	_, err = file.Write(markdownBytes)
	if err != nil {
		app.Logger.Error().Err(err).Str("page", page.Name).Msg("Error writing to file")
	}

}

// createFile creates nested directories if needed and then calls os.Create
func createFile(p string) (*os.File, error) {
	if err := os.MkdirAll(filepath.Dir(p), 0770); err != nil {
		return nil, err
	}
	return os.Create(p)
}

package downloader

import (
	"fmt"
	"os"
	"path"
	"path/filepath"
	"sort"
	"sync"
	"time"

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

	// Create the contents of the SUMMARY.md file

	summaryContents := "# Summary\n"
	summaryContents += "[Introduction](README.md)\n\n"

	for _, book := range w.Books {
		summaryContents += "\n# " + book.Name + "\n\n"
		priorityQueue := make([]interface{}, 0)
		for _, chapter := range book.Chapters {
			priorityQueue = append(priorityQueue, chapter)
		}
		for _, indiePage := range book.IndiePages {
			priorityQueue = append(priorityQueue, indiePage)
		}
		sort.Slice(priorityQueue, func(i, j int) bool {
			var pi, pj int
			switch priorityQueue[i].(type) {
			case bookstackclient.WikiChapter:
				pi = priorityQueue[i].(bookstackclient.WikiChapter).Priority
			case bookstackclient.WikiPage:
				pi = priorityQueue[i].(bookstackclient.WikiPage).Priority
			}
			switch priorityQueue[j].(type) {
			case bookstackclient.WikiChapter:
				pj = priorityQueue[j].(bookstackclient.WikiChapter).Priority
			case bookstackclient.WikiPage:
				pj = priorityQueue[j].(bookstackclient.WikiPage).Priority
			}
			return pi < pj
		})
		for _, item := range priorityQueue {
			switch item := item.(type) {
			case bookstackclient.WikiChapter:
				summaryContents += fmt.Sprint("- [", item.Name, "]()\n")
				for _, page := range item.Pages {
					summaryContents += fmt.Sprint("    - [", page.Name, "](", book.Slug, "/", item.Slug, "/", page.Slug, ".md)\n")
				}
			case bookstackclient.WikiPage:
				summaryContents += fmt.Sprint("- [", item.Name, "](", book.Slug, "/", item.Slug, ".md)\n")
			}
		}
	}
	summaryContents += fmt.Sprintf("\n---\n\nGenerated on %s.\n\n", time.Now().Format(time.RFC850))

	createFileWithContents(app, summaryContents, fmt.Sprint(app.Config.DownloadLocation, "/SUMMARY.md"))
	createFileWithContents(app, summaryContents, fmt.Sprint(app.Config.DownloadLocation, "/README.md"))
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

func createFileWithContents(app *config.Application, contents string, filepath string) {
	file, err := createFile(path.Clean(filepath))
	if err != nil {
		app.Logger.Error().Err(err).Str("page", filepath).Msg("Error creating file")
	}

	_, err = file.Write([]byte(contents))
	if err != nil {
		app.Logger.Error().Err(err).Str("page", filepath).Msg("Error writing to file")
	} else {
		app.Logger.Info().Str("page", filepath).Msg("Written to file")
	}
}

// createFile creates nested directories if needed and then calls os.Create
func createFile(p string) (*os.File, error) {
	if err := os.MkdirAll(filepath.Dir(p), 0770); err != nil {
		return nil, err
	}
	return os.Create(p)
}

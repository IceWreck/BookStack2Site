package downloader

import (
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

			go func(p bookstackclient.WikiPage) {
				defer wg.Done()
				downloadPage(app, p)
				// release semaphore
				<-sem
			}(indiePage)
		}
	}

	wg.Wait()
}

func downloadPage(app *config.Application, page bookstackclient.WikiPage) {
	app.Logger.Info().Str("page", page.Name).Msg("Downloading Page")
}

package bookstackclient

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/IceWreck/BookStack2Site/config"
)

// FetchBooks returns all the books in the wiki sorted by name.
func FetchBooks(app *config.Application) (Books, error) {

	resStruct := Books{}
	count := 10 // limit
	offset := 0

	// keep doing subsequent requests until you get the entire queryset
	for {
		req, err := http.NewRequest("GET", fmt.Sprint(app.Config.BookStackEndpoint, "/api/books"), nil)
		if err != nil {
			return resStruct, err
		}

		q := req.URL.Query()
		q.Set("offset", fmt.Sprint(offset))
		q.Set("count", fmt.Sprint(count))
		q.Set("sort", "+name")

		req.URL.RawQuery = q.Encode()

		//app.Logger.Debug().Str("req", fmt.Sprint(req)).Msg("")
		res, err := authenticatedDo(app, req)
		if err != nil {
			app.Logger.Debug().Err(err).Msg("performing request")
			return resStruct, err
		}
		body, err := io.ReadAll(res.Body)
		if err != nil {
			app.Logger.Debug().Err(err).Msg("reading body")
			return resStruct, err
		}
		tmpStruct := Books{}
		err = json.Unmarshal(body, &tmpStruct)
		if err != nil {
			app.Logger.Debug().Err(err).Msg("unmarshaling JSON")
			app.Logger.Debug().Str("body", string(body)).Msg("")
			return resStruct, err
		}
		app.Logger.Debug().Str("tmpStruct", fmt.Sprint(tmpStruct)).Msg("")

		resStruct.Data = append(resStruct.Data, tmpStruct.Data...)

		// stop further requests if number of results < limit
		if len(tmpStruct.Data) == 0 {
			break
		}
		offset += count
	}

	return resStruct, nil
}

// FetchChapters returns all the chapters of a book sorted by priority.
func FetchChapters(app *config.Application, bookID int) (Chapters, error) {

	resStruct := Chapters{}
	count := 10 // limit
	offset := 0

	// keep doing subsequent requests until you get the entire queryset
	for {
		req, err := http.NewRequest("GET", fmt.Sprint(app.Config.BookStackEndpoint, "/api/chapters"), nil)
		if err != nil {
			return resStruct, err
		}

		q := req.URL.Query()
		q.Set("offset", fmt.Sprint(offset))
		q.Set("count", fmt.Sprint(count))
		q.Set("count", fmt.Sprint(count))
		q.Set("sort", "+priority")
		q.Set("filter[book_id]", fmt.Sprint(bookID))

		req.URL.RawQuery = q.Encode()

		//app.Logger.Debug().Str("req", fmt.Sprint(req)).Msg("")
		res, err := authenticatedDo(app, req)
		if err != nil {
			app.Logger.Debug().Err(err).Msg("performing request")
			return resStruct, err
		}
		body, err := io.ReadAll(res.Body)
		if err != nil {
			app.Logger.Debug().Err(err).Msg("reading body")
			return resStruct, err
		}
		tmpStruct := Chapters{}
		err = json.Unmarshal(body, &tmpStruct)
		if err != nil {
			app.Logger.Debug().Err(err).Msg("unmarshaling JSON")
			app.Logger.Debug().Str("body", string(body)).Msg("")
			return resStruct, err
		}
		app.Logger.Debug().Str("tmpStruct", fmt.Sprint(tmpStruct)).Msg("")

		resStruct.Data = append(resStruct.Data, tmpStruct.Data...)

		// stop further requests if number of results < limit
		if len(tmpStruct.Data) == 0 {
			break
		}
		offset += count
	}

	return resStruct, nil
}

// FetchPages returns pages sorted by priority.
// If chapterID = 0 then it returns independent (non chapter) pages of the book.
// If chapterID != 0 then it returns pages of the chapter.
func FetchPages(app *config.Application, bookID int, chapterID int) (Chapters, error) {

	resStruct := Chapters{}
	count := 10 // limit
	offset := 0

	// keep doing subsequent requests until you get the entire queryset
	for {
		req, err := http.NewRequest("GET", fmt.Sprint(app.Config.BookStackEndpoint, "/api/pages"), nil)
		if err != nil {
			return resStruct, err
		}

		q := req.URL.Query()
		q.Set("offset", fmt.Sprint(offset))
		q.Set("count", fmt.Sprint(count))
		q.Set("count", fmt.Sprint(count))
		q.Set("sort", "+priority")
		q.Set("filter[book_id]", fmt.Sprint(bookID))
		q.Set("filter[chapter_id]", fmt.Sprint(chapterID))

		req.URL.RawQuery = q.Encode()

		//app.Logger.Debug().Str("req", fmt.Sprint(req)).Msg("")
		res, err := authenticatedDo(app, req)
		if err != nil {
			app.Logger.Debug().Err(err).Msg("performing request")
			return resStruct, err
		}
		body, err := io.ReadAll(res.Body)
		if err != nil {
			app.Logger.Debug().Err(err).Msg("reading body")
			return resStruct, err
		}
		tmpStruct := Chapters{}
		err = json.Unmarshal(body, &tmpStruct)
		if err != nil {
			app.Logger.Debug().Err(err).Msg("unmarshaling JSON")
			app.Logger.Debug().Str("body", string(body)).Msg("")
			return resStruct, err
		}
		app.Logger.Debug().Str("tmpStruct", fmt.Sprint(tmpStruct)).Msg("")

		resStruct.Data = append(resStruct.Data, tmpStruct.Data...)

		// stop further requests if number of results < limit
		if len(tmpStruct.Data) == 0 {
			break
		}
		offset += count
	}

	return resStruct, nil
}

func FetchWiki(app *config.Application) (Wiki, error) {
	w := Wiki{
		Name:    "",
		Books:   []WikiBook{},
		Shelves: nil,
	}
	books, err := FetchBooks(app)
	if err != nil {
		return w, err
	}

	for _, book := range books.Data {

		// create a temporary book and fill initial data
		tmpBook := WikiBook{
			BookID:     book.ID,
			Name:       book.Name,
			Slug:       book.Slug,
			Chapters:   []WikiChapter{},
			IndiePages: []WikiPage{},
		}

		// fetch chapter for that book
		chapters, err := FetchChapters(app, book.ID)
		if err != nil {
			return w, err
		}
		for _, chapter := range chapters.Data {
			tmpChapter := WikiChapter{
				ChapterID: chapter.ID,
				Name:      chapter.Name,
				Slug:      chapter.Slug,
				Priority:  chapter.Priority,
				Pages:     []WikiPage{},
			}

			// fetch pages for that chapter
			pages, err := FetchPages(app, book.ID, chapter.ID)
			if err != nil {
				return w, err
			}
			for _, page := range pages.Data {

				// add temporary page to the temporary chapter
				tmpChapter.Pages = append(tmpChapter.Pages, WikiPage{
					PageID:   page.ID,
					Name:     page.Name,
					Slug:     page.Slug,
					Priority: page.Priority,
				})
			}

			// add temporary chapter to the temporary book
			tmpBook.Chapters = append(tmpBook.Chapters, tmpChapter)
		}

		// fetch independent pages for that book
		pages, err := FetchPages(app, book.ID, 0)
		if err != nil {
			return w, err
		}
		for _, page := range pages.Data {
			// add temporary independent page to the temporary book
			tmpBook.IndiePages = append(tmpBook.IndiePages, WikiPage{
				PageID:   page.ID,
				Name:     page.Name,
				Slug:     page.Slug,
				Priority: page.Priority,
			})
		}

		// add temporary book to the wiki
		w.Books = append(w.Books, tmpBook)
	}

	// TODO: fill up shelves

	return w, nil
}

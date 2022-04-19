package bookstackclient

import (
	"fmt"
	"net/http"

	"github.com/IceWreck/BookStack2Site/config"
)

func authenticatedDo(app *config.Application, req *http.Request) (*http.Response, error) {
	req.Header.Add("User-Agent", "BookStack2Site")
	req.Header.Add("Authorization",
		fmt.Sprintf("Token %s:%s", app.Config.BookStackAPITokenID, app.Config.BookStackAPITokenSecret))
	res, err := app.Client.Do(req)
	if err != nil {
		return nil, err
	}
	return res, nil
}

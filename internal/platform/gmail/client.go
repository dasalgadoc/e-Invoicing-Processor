package gmail

import (
	"context"
	"golang.org/x/oauth2"
	"net/http"
)

func getClient(config *oauth2.Config) *http.Client {
	tok, err := tokenFromFile()
	if err != nil {
		tok = getTokenFromWeb(config)
		saveToken(tok)
	}

	return config.Client(context.Background(), tok)
}

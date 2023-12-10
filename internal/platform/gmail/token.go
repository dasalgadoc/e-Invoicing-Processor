package gmail

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/dasalgadoc/e-Invoicing-Processor/kit/domain/errors"
	"golang.org/x/oauth2"
	"log"
	"os"
)

const (
	tokenGoFile = "gmail/token.go"

	tokenFileName = "env/token.json"
	stateToken    = "state-token"
)

func tokenFromFile() (*oauth2.Token, error) {
	f, err := os.Open(tokenFileName)
	if err != nil {
		return nil, errors.NewProjectError(tokenGoFile, errors.OSError, err.Error())
	}
	defer func(f *os.File) {
		err := f.Close()
		if err != nil {
			panic(errors.NewProjectError(tokenGoFile, errors.OSError, err.Error()))
		}
	}(f)

	tok := &oauth2.Token{}
	err = json.NewDecoder(f).Decode(tok)
	if err != nil {
		return nil, errors.NewProjectError(tokenGoFile, errors.ExternalLibraryError, err.Error())
	}

	return tok, nil
}

func getTokenFromWeb(config *oauth2.Config) *oauth2.Token {
	authURL := config.AuthCodeURL(stateToken, oauth2.AccessTypeOffline)
	fmt.Printf("Go to the following link in your browser then type the "+
		"authorization code: \n%v\n", authURL)

	var authCode string
	if _, err := fmt.Scan(&authCode); err != nil {
		log.Fatalf("Unable to read authorization authCode: %v", err)
	}

	tok, err := config.Exchange(context.Background(), authCode)
	if err != nil {
		log.Fatalf("Error exchange authCode for access token: %v", err)
	}

	return tok
}

func saveToken(token *oauth2.Token) {
	log.Printf("Saving credential token file to: %s\n", tokenFileName)
	f, err := os.OpenFile(tokenFileName, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0600)
	if err != nil {
		log.Fatalf("Unable to create oauth token: %v", err)
	}
	defer func(f *os.File) {
		err := f.Close()
		if err != nil {
			panic(errors.NewProjectError(tokenGoFile, errors.OSError, err.Error()))
		}
	}(f)

	err = json.NewEncoder(f).Encode(token)
	if err != nil {
		log.Fatalf("Unable to encode oauth token: %v", err)
	}
}

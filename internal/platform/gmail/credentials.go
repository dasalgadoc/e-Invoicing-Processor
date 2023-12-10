package gmail

import (
	"fmt"
	"github.com/dasalgadoc/e-Invoicing-Processor/kit/domain/errors"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
	"google.golang.org/api/gmail/v1"
	"os"
)

const (
	credentialsGoFile = "gmail/credentials.go"
	credentialsFile   = "env/credentials.json"
)

func setConfigFromCredentials() (*oauth2.Config, error) {
	b, err := os.ReadFile(credentialsFile)
	if err != nil {
		msg := fmt.Sprintf("Error reading credentials file: %s", err)
		return nil, errors.NewProjectError(credentialsGoFile, errors.OSError, msg)
	}

	config, err := google.ConfigFromJSON(b, gmail.GmailReadonlyScope)
	if err != nil {
		msg := fmt.Sprintf("Error parsing credentials file: %s", err)
		return nil, errors.NewProjectError(credentialsGoFile, errors.ExternalLibraryError, msg)
	}

	return config, nil
}

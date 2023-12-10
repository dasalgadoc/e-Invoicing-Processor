package gmail

import (
	"fmt"
	"github.com/dasalgadoc/e-Invoicing-Processor/kit/domain/errors"
	"google.golang.org/api/gmail/v1"
)

type ScrapSource struct {
	service *gmail.Service
}

func NewGmailScrapSource() (*ScrapSource, *errors.ProjectError) {
	srv, err := buildService()
	if err != nil {
		return nil, errors.NewProjectError("gmail_scrap_source.go", errors.BuilderError, err.Error())
	}
	return &ScrapSource{
		service: srv,
	}, nil
}

func (s *ScrapSource) ListMessages() *errors.ProjectError {
	messages, err := s.service.Users.Messages.List("me").Do()
	if err != nil {
		return errors.NewProjectError("gmail_scrap_source.go", errors.ServiceError, err.Error())
	}

	for _, message := range messages.Messages {
		fmt.Println("Subject:", message.Snippet)
	}

	return nil
}

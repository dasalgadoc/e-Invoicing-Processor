package gmail

import (
	"fmt"
	"github.com/dasalgadoc/e-Invoicing-Processor/kit/domain/criteria"
	"github.com/dasalgadoc/e-Invoicing-Processor/kit/domain/errors"
	"google.golang.org/api/gmail/v1"
)

const (
	sourceFile = "gmail_scrap_source.go"

	userId          = "me"
	attachmentQuery = "has:attachment"

	subject = "Subject"
)

type ScrapSource struct {
	service *gmail.Service
}

func NewGmailScrapSource() (*ScrapSource, *errors.ProjectError) {
	srv, err := buildService()
	if err != nil {
		return nil, errors.NewProjectError(sourceFile, errors.BuilderError, err.Error())
	}
	return &ScrapSource{
		service: srv,
	}, nil
}

func (s *ScrapSource) ListMessagesWithAttachments(criteria criteria.Criteria) *errors.ProjectError {
	resp, err := s.service.Users.Messages.List(userId).Q(attachmentQuery).Do()
	if err != nil {
		return errors.NewProjectError(sourceFile, errors.ServiceError, err.Error())
	}

	for i, message := range resp.Messages {
		msgContent, err := s.service.Users.Messages.Get(userId, message.Id).Do()
		if err != nil {
			return errors.NewProjectError(sourceFile, errors.ServiceError, err.Error())
		}

		for _, header := range msgContent.Payload.Headers {
			if header.Name == subject {
				fmt.Println(i, header.Value)
			}
		}
	}

	return nil
}

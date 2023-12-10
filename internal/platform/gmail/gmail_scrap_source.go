package gmail

import (
	"github.com/dasalgadoc/e-Invoicing-Processor/internal/scraping"
	"github.com/dasalgadoc/e-Invoicing-Processor/kit/domain/criteria"
	"github.com/dasalgadoc/e-Invoicing-Processor/kit/domain/errors"
	"github.com/schollz/progressbar/v3"
	"google.golang.org/api/gmail/v1"
)

const (
	sourceFile = "gmail_scrap_source.go"

	userId                 = "me"
	attachmentCompressFile = "has:attachment filename:zip"

	to      = "Delivered-To"
	from    = "From"
	subject = "Subject"
	date    = "Date"
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

func (s *ScrapSource) GetInvoicingMessages(criteria criteria.Criteria) ([]scraping.Message, *errors.ProjectError) {
	resp, err := s.service.Users.Messages.List(userId).Q(attachmentCompressFile).Do()
	if err != nil {
		return nil, errors.NewProjectError(sourceFile, errors.ServiceError, err.Error())
	}

	progressBar := progressbar.Default(int64(len(resp.Messages)), "Gathering messages")
	var messages []scraping.Message
	for _, message := range resp.Messages {
		msgContent, err := s.service.Users.Messages.Get(userId, message.Id).Do()
		if err != nil {
			return nil, errors.NewProjectError(sourceFile, errors.PartialError, err.Error())
		}

		var messagesParts = make(map[string]string)
		for _, header := range msgContent.Payload.Headers {
			messagesParts[header.Name] = header.Value
		}

		msg := scraping.NewMessage(message.Id, messagesParts[to], messagesParts[from], messagesParts[subject], messagesParts[date])
		if msg != nil {
			messages = append(messages, *msg)
		}
		progressBar.Add(1)
	}

	return messages, nil
}

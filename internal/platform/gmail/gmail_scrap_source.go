package gmail

import (
	"encoding/base64"
	"fmt"
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
	startDateQuery         = "after: 2023/01/01"
	endDateQuery           = "before: 2023/12/31"

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
	query := fmt.Sprintf("%s %s %s", attachmentCompressFile, startDateQuery, endDateQuery)
	resp, err := s.service.Users.Messages.List(userId).Q(query).Do()
	if err != nil {
		return nil, errors.NewProjectError(sourceFile, errors.ServiceError, err.Error())
	}

	return s.getMessagesAndAttachment(resp)
}

func (s *ScrapSource) getMessagesAndAttachment(resp *gmail.ListMessagesResponse) ([]scraping.Message, *errors.ProjectError) {
	progressBar := progressbar.Default(int64(len(resp.Messages)), "Gathering messages")

	var messages []scraping.Message
	for _, message := range resp.Messages {
		messageContent, err := s.service.Users.Messages.Get(userId, message.Id).Do()
		if err != nil {
			return nil, errors.NewProjectError(sourceFile, errors.PartialError, err.Error())
		}

		messagesParts := s.getMessageParts(messageContent)

		attachmentName, attachment, attachmentErr := s.getAttachmentForMessage(message.Id, messageContent)
		if attachmentErr != nil {
			return nil, errors.NewProjectError(sourceFile, errors.PartialError, err.Error())
		}

		msg := scraping.NewMessage(message.Id, messagesParts[to], messagesParts[from], messagesParts[subject], messagesParts[date], attachmentName, attachment)
		if msg != nil {
			messages = append(messages, *msg)
		}

		progressBar.Add(1)
	}

	return messages, nil
}

func (s *ScrapSource) getMessageParts(messageContent *gmail.Message) map[string]string {
	var messagesParts = make(map[string]string)
	for _, header := range messageContent.Payload.Headers {
		messagesParts[header.Name] = header.Value
	}
	return messagesParts
}

func (s *ScrapSource) getAttachmentForMessage(msgId string, msgContent *gmail.Message) (string, []byte, *errors.ProjectError) {
	var attachment []byte
	var attachmentName string
	for _, part := range msgContent.Payload.Parts {
		if part.Filename != "" {
			attachmentName = part.Filename

			data, err := s.service.Users.Messages.Attachments.Get(userId, msgId, part.Body.AttachmentId).Do()
			if err != nil {
				return "", nil, errors.NewProjectError(sourceFile, errors.PartialError, err.Error())
			}

			attachment, err = base64.URLEncoding.DecodeString(data.Data)
			if err != nil {
				return "", nil, errors.NewProjectError(sourceFile, errors.PartialError, err.Error())
			}
		}
	}
	return attachmentName, attachment, nil
}

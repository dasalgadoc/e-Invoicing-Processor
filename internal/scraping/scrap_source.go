package scraping

import "github.com/dasalgadoc/e-Invoicing-Processor/kit/domain/errors"

type ScrapSource interface {
	ListMessagesWithAttachments() *errors.ProjectError
}

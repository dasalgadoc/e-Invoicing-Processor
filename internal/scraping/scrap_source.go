package scraping

import (
	"github.com/dasalgadoc/e-Invoicing-Processor/kit/domain/criteria"
	"github.com/dasalgadoc/e-Invoicing-Processor/kit/domain/errors"
)

type ScrapSource interface {
	GetInvoicingMessages(criteria criteria.Criteria) ([]Message, *errors.ProjectError)
}

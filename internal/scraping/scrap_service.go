package scraping

import (
	"github.com/dasalgadoc/e-Invoicing-Processor/kit/domain/criteria"
	"github.com/dasalgadoc/e-Invoicing-Processor/kit/domain/errors"
)

type ScrapService struct {
	scrapSource ScrapSource
}

func NewScrapSource(source ScrapSource) *ScrapService {
	return &ScrapService{
		scrapSource: source,
	}
}

func (s *ScrapService) Invoke() *errors.ProjectError {
	c := criteria.NewCriteria(0, 0, "", "", nil)
	_, err := s.scrapSource.GetInvoicingMessages(*c)
	return err
}

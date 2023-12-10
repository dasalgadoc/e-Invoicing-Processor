package scraping

import "github.com/dasalgadoc/e-Invoicing-Processor/kit/domain/errors"

type ScrapService struct {
	scrapSource ScrapSource
}

func NewScrapSource(source ScrapSource) *ScrapService {
	return &ScrapService{
		scrapSource: source,
	}
}

func (s *ScrapService) Invoke() *errors.ProjectError {
	return s.scrapSource.ListMessagesWithAttachments()
}

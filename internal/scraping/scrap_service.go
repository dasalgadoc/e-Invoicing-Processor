package scraping

type ScrapService struct {
	scrapSource ScrapSource
}

func NewScrapSource(source ScrapSource) *ScrapService {
	return &ScrapService{
		scrapSource: source,
	}
}

func (s *ScrapService) Invoke() {
	s.scrapSource.ListMessages()
}

package scraping

type ServiceResponse struct {
	TotalMessages     int
	DestinationFolder string
	Messages          []Message
}

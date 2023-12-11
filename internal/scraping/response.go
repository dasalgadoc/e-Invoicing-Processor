package scraping

type ServiceResponse struct {
	TotalMessages     int
	DestinationFolder string
	Messages          []Message
}

func (s *ServiceResponse) FromByMessage(index int) string {
	if !s.isIndexValid(index) {
		return ""
	}
	return s.Messages[index].from
}

func (s *ServiceResponse) SubjectByMessage(index int) string {
	if !s.isIndexValid(index) {
		return ""
	}
	return s.Messages[index].subject
}

func (s *ServiceResponse) DateByMessage(index int) string {
	if !s.isIndexValid(index) {
		return ""
	}
	return s.Messages[index].date
}

func (s *ServiceResponse) AttachmentNameByMessage(index int) string {
	if !s.isIndexValid(index) {
		return ""
	}
	return s.Messages[index].attachmentName
}

func (s *ServiceResponse) isIndexValid(index int) bool {
	return index > 0 && index < len(s.Messages)
}

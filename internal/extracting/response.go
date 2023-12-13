package extracting

type ServiceResponse struct {
	OriginalFolder  string
	ExtractedFolder string
	Attachments     []Attachment
}

type Attachment struct {
	Receiver       string
	Sender         string
	Subject        string
	Date           string
	Name           string
	ExtractedFiles []string
}

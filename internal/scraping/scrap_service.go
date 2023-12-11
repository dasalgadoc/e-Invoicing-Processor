package scraping

import (
	"fmt"
	"github.com/dasalgadoc/e-Invoicing-Processor/kit/domain/criteria"
	"github.com/dasalgadoc/e-Invoicing-Processor/kit/domain/errors"
	"github.com/schollz/progressbar/v3"
	"os"
	"time"
)

const (
	serviceName  = "scraping/scrap_service.go"
	basePath     = "temp"
	folderFormat = "2006-01-02_15_04_05"
)

type ServiceResponse struct {
	TotalMessages     int
	DestinationFolder string
	Messages          []Message
}

type ScrapService struct {
	scrapSource ScrapSource
}

func NewScrapSource(source ScrapSource) *ScrapService {
	return &ScrapService{
		scrapSource: source,
	}
}

func (s *ScrapService) Invoke() (*ServiceResponse, *errors.ProjectError) {
	c := criteria.NewCriteria(0, 0, "", "", nil)
	messages, err := s.scrapSource.GetInvoicingMessages(*c)
	if err != nil {
		return nil, errors.NewProjectError(serviceName, errors.ServiceError, err.Error())
	}
	fmt.Println("Message gathered: ", len(messages))

	destination, attachErr := s.saveAttachments(messages)
	if attachErr != nil {
		return nil, attachErr
	}

	resp := &ServiceResponse{
		TotalMessages:     len(messages),
		DestinationFolder: destination,
		Messages:          messages,
	}
	return resp, nil
}

func (s *ScrapService) saveAttachments(messages []Message) (string, *errors.ProjectError) {
	basePath := folderPath()
	err := os.Mkdir(basePath, os.ModePerm)
	if err != nil {
		return basePath, errors.NewProjectError(serviceName, errors.OSError, err.Error())
	}

	fmt.Println("Saving attachments in: ", basePath)

	progressBar := progressbar.Default(int64(len(messages)), "Saving Attachments")
	for _, m := range messages {
		path := fmt.Sprintf("%s/%s", basePath, m.attachmentName)

		file, err := os.Create(path)
		if err != nil {
			return basePath, errors.NewProjectError(serviceName, errors.OSError, err.Error())
		}
		defer file.Close()

		_, err = file.Write(m.attachment)
		if err != nil {
			return basePath, errors.NewProjectError(serviceName, errors.OSError, err.Error())
		}

		progressBar.Add(1)
	}
	return basePath, nil
}

func folderPath() string {
	now := time.Now()

	return fmt.Sprintf("%s/%s", basePath, now.Format(folderFormat))
}

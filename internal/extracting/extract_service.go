package extracting

import (
	"fmt"
	"github.com/dasalgadoc/e-Invoicing-Processor/internal/scraping"
	"github.com/dasalgadoc/e-Invoicing-Processor/kit/domain/errors"
	"github.com/schollz/progressbar/v3"
)

const (
	serviceName       = "extracting/extract_service.go"
	destinationSuffix = "_extracted/"
)

type ExtractService struct {
}

func NewExtractService() *ExtractService {
	return &ExtractService{}
}

func (e *ExtractService) Invoke(data scraping.ServiceResponse) *errors.ProjectError {
	originFolder := data.DestinationFolder
	destinationFolder := fmt.Sprintf("%s%s", originFolder, destinationSuffix)

	progressBar := progressbar.Default(int64(data.TotalMessages), "Extracting attachments")
	for _, m := range data.Messages {
		err := unzipFileToDestination(data.DestinationFolder, m.AttachmentName(), destinationFolder)
		if err != nil {
			return errors.NewProjectError(serviceName, errors.ServiceError, err.Error())
		}

		progressBar.Add(1)
	}

	return nil
}

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

func (e *ExtractService) Invoke(data scraping.ServiceResponse) (*ServiceResponse, *errors.ProjectError) {
	originFolder := data.DestinationFolder
	destinationFolder := fmt.Sprintf("%s%s", originFolder, destinationSuffix)

	resp := &ServiceResponse{
		OriginalFolder:  originFolder,
		ExtractedFolder: destinationFolder,
		Attachments:     make([]Attachment, 0),
	}

	fmt.Printf("Extracting files in: %s\n", destinationFolder)
	progressBar := progressbar.Default(int64(data.TotalMessages), "Extracting attachments")
	for _, m := range data.Messages {
		files, err := unzipFileToDestination(data.DestinationFolder, m.AttachmentName(), destinationFolder)
		if err != nil {
			return resp, errors.NewProjectError(serviceName, errors.ServiceError, err.Error())
		}
		resp.Attachments = append(resp.Attachments, Attachment{
			Receiver:       m.To(),
			Sender:         m.From(),
			Subject:        m.Subject(),
			Date:           m.Date(),
			Name:           m.AttachmentName(),
			ExtractedFiles: files,
		})

		progressBar.Add(1)
	}

	return resp, nil
}

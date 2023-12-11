package extracting

import "github.com/dasalgadoc/e-Invoicing-Processor/kit/domain/errors"

type ExtractService struct {
}

func NewExtractService() *ExtractService {
	return &ExtractService{}
}

func (e *ExtractService) Invoke() *errors.ProjectError {
	return nil
}

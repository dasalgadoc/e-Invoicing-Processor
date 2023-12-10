package gmail

import (
	"context"
	"github.com/dasalgadoc/e-Invoicing-Processor/kit/domain/errors"
	"google.golang.org/api/gmail/v1"
	"google.golang.org/api/option"
)

func buildService() (*gmail.Service, *errors.ProjectError) {
	ctx := context.Background()

	config, err := setConfigFromCredentials()
	if err != nil {
		return nil, errors.NewProjectError("gmail/service.go", errors.ServiceError, err.Error())
	}

	client := getClient(config)

	service, err := gmail.NewService(ctx, option.WithHTTPClient(client))
	if err != nil {
		return nil, errors.NewProjectError("gmail/service.go", errors.ServiceError, err.Error())
	}

	return service, nil
}

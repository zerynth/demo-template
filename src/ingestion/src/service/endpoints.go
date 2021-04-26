package service

import (
	"context"
	"ingestion/models"

	"github.com/go-kit/kit/endpoint"
)

// Endpoints struct containing endpoints definition
type Endpoints struct {
	InsertDataEndpoint      endpoint.Endpoint
	InsertConditionEndpoint endpoint.Endpoint
}

// MakeServerEndpoints returns service endpoints
func MakeServerEndpoints(s IIngestionService) Endpoints {
	return Endpoints{
		InsertDataEndpoint:      MakeInsertDataEndpoint(s),
		InsertConditionEndpoint: MakeInsertConditionEndpoint(s),
	}
}

func MakeInsertDataEndpoint(service IIngestionService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req := request.(models.InsertDataRequest)
		e := service.InsertData(ctx, &req)
		if e != nil {
			return models.Response{Error: e}, e
		}
		return models.Response{Error: nil}, nil
	}
}

func MakeInsertConditionEndpoint(service IIngestionService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req := request.(models.InsertConditionRequest)
		e := service.InsertCondition(ctx, &req)
		if e != nil {
			return models.Response{Error: e}, e
		}
		return models.Response{Error: nil}, nil
	}
}

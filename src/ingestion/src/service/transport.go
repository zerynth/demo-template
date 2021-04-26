package service

import (
	"context"
	"encoding/json"
	"ingestion/models"
	"net/http"

	"github.com/go-kit/kit/transport"

	"github.com/go-kit/kit/log"
	httpTransport "github.com/go-kit/kit/transport/http"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

func MakeHTTPHandler(s IIngestionService, logger log.Logger) http.Handler {
	r := mux.NewRouter()

	r.Use(handlers.ProxyHeaders)

	var e = MakeServerEndpoints(s)

	options := []httpTransport.ServerOption{
		httpTransport.ServerErrorHandler(transport.NewLogErrorHandler(logger)),
		httpTransport.ServerErrorEncoder(encodeError),
	}

	// POST /zdm/data							insert a data stream
	r.Methods("POST").Path("/zdm/data").Handler(postInsertDataHandler(&e, options))
	// POST /zdm/condition							insert a condition stream
	r.Methods("POST").Path("/zdm/condition").Handler(postInsertConditionHandler(&e, options))
	return r
}

func postInsertDataHandler(e *Endpoints, options []httpTransport.ServerOption) *httpTransport.Server {
	return httpTransport.NewServer(
		e.InsertDataEndpoint,
		decodeInsertDataRequest,
		encodeResponse,
		options...,
	)
}

func postInsertConditionHandler(e *Endpoints, options []httpTransport.ServerOption) *httpTransport.Server {
	return httpTransport.NewServer(
		e.InsertConditionEndpoint,
		decodeInsertConditionRequest,
		encodeResponse,
		options...,
	)
}

func decodeInsertDataRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var req models.InsertDataRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		return nil, err
	}
	return req, nil
}

func decodeInsertConditionRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var req models.InsertConditionRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		return nil, err
	}
	return req, nil
}

type errorer interface {
	error() error
}

func encodeResponse(ctx context.Context, w http.ResponseWriter, response interface{}) error {
	e, ok := response.(errorer)
	if ok && e.error() != nil {
		// Not a Go kit transport error, but a business-logic error.
		// Provide those as HTTP errors.
		encodeError(ctx, e.error(), w)
		return nil
	}
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	return json.NewEncoder(w).Encode(response)
}

func encodeError(_ context.Context, err error, w http.ResponseWriter) {
	if err == nil {
		panic("encodeError with nil error")
	}
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(http.StatusBadRequest)
	_ = json.NewEncoder(w).Encode(err)
}

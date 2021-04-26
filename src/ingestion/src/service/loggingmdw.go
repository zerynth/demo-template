package service

import (
	"context"
	"fmt"
	"ingestion/models"
	"time"

	"github.com/go-kit/kit/log"
)

type Middleware func(service IIngestionService) IIngestionService

type loggingMiddleware struct {
	next   IIngestionService
	logger log.Logger
}

func LoggingMiddleware(logger log.Logger) Middleware {
	return func(next IIngestionService) IIngestionService {
		return &loggingMiddleware{
			next:   next,
			logger: logger,
		}
	}
}

func (mw loggingMiddleware) InsertData(ctx context.Context, req *models.InsertDataRequest) (e error) {
	defer func(begin time.Time) {
		mw.logger.Log("method", "InsertData", "len", len(req.Result), "req", fmt.Sprintf("%+v", req), "duration", time.Since(begin), "error", e)
	}(time.Now())
	return mw.next.InsertData(ctx, req)
}

func (mw loggingMiddleware) InsertCondition(ctx context.Context, req *models.InsertConditionRequest) (e error) {
	defer func(begin time.Time) {
		mw.logger.Log("method", "InsertCondition", "len", len(req.IncomingCondition), "req", fmt.Sprintf("%+v", req), "duration", time.Since(begin), "error", e)
	}(time.Now())
	return mw.next.InsertCondition(ctx, req)
}

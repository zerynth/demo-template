package service

import (
	"context"

	"ingestion/database"
	"ingestion/models"
)

type IIngestionService interface {
	InsertData(ctx context.Context, data *models.InsertDataRequest) error
	InsertCondition(ctx context.Context, data *models.InsertConditionRequest) error
}

type IngestionService struct {
	DbRepository database.IRepository
}

func NewIngManagerService(db database.IRepository) IIngestionService {
	return &IngestionService{
		DbRepository: db,
	}
}

func (ts *IngestionService) InsertData(_ context.Context, req *models.InsertDataRequest) error {
	for _, d := range req.Result {
		d := &models.Data{
			Id:              d.Id,
			TimestampDevice: d.TimestampDevice,
			WorkspaceID:     d.WorkspaceID,
			FleetID:         d.FleetID,
			Tag:             d.Tag,
			DeviceID:        d.DeviceID,
			DeviceName:      d.DeviceName,
			Payload:         d.Payload,
		}
		err := ts.DbRepository.InsertData(d, req.BatchId)
		if err != nil {
			return err
		}
	}
	return nil
}

func (ts *IngestionService) InsertCondition(_ context.Context, req *models.InsertConditionRequest) error {
	for _, d := range req.IncomingCondition {
		d := &models.Condition{
			ConditionID: d.ConditionID,
			Tag:         d.Tag,
			WorkspaceID: d.WorkspaceID,
			FleetID:     d.FleetID,
			DeviceName:  d.DeviceName,
			DeviceID:    d.DeviceID,
			Payload:     d.Payload,
			PayloadF:    d.PayloadF,
			Start:       d.Start,
			Finish:      d.Finish,
			Duration:    d.Duration,
		}
		err := ts.DbRepository.InsertCondition(d)
		if err != nil {
			return err
		}
	}
	return nil
}

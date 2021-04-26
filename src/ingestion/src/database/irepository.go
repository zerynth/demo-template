package database

import (
	"ingestion/models"
)

type IRepository interface {
	InsertData(row *models.Data, batchId string) error
	InsertCondition(row *models.Condition) error
}

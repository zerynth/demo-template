package database

import (
	"database/sql"
	"errors"
	"fmt"
	"time"

	"ingestion/models"
)

type TSConfig struct {
	Host         string `default:"tsdb"`
	Port         string `default:"5432"`
	User         string `default:"postgres"`
	Password     string `default:"password"`
	Name         string `default:"dbdemo"`
	RetryConnect int    `default:"20"`
	RetryDelay   int    `default:"5"`
}

// ModelRepository defines the service's database
type ModelRepository struct {
	db *sql.DB
}

// ConnectToDB allows to connect to a postgres db using database/sql package
func ConnectToDB(ts TSConfig) (*sql.DB, error) {

	var err error
	var database *sql.DB
	for i := 0; i < ts.RetryConnect; i++ {
		psqlInfo := fmt.Sprintf("host=%s port=%s user=%s password=%s  dbname=%s sslmode=disable", ts.Host, ts.Port, ts.User, ts.Password, ts.Name)
		fmt.Println("DB connection: " + psqlInfo)
		database, err = sql.Open("postgres", psqlInfo)
		if err != nil {
			fmt.Println("Waiting for db: ", i, err.Error())
			time.Sleep(time.Duration(ts.RetryDelay) * time.Second)
		} else {
			if err = database.Ping(); err != nil {
				fmt.Println("Cannot ping db: ", err.Error())
				//return nil, err
				time.Sleep(time.Duration(ts.RetryDelay) * time.Second)
			} else {
				fmt.Println("Connected to db succesfully")
				return database, nil
			}
		}

	}
	return nil, errors.New("Can't connect to database")

}

// NewModelRepository creates a new database struct
func NewModelRepository(database *sql.DB) *ModelRepository {
	return &ModelRepository{
		db: database,
	}
}

// InsertData insert periodic from devices into workspace tables
func (ts *ModelRepository) InsertData(data *models.Data, batchId string) error {
	q := fmt.Sprintf(`INSERT INTO data (batch_id, id, timestamp_device, timestamp_in, workspace_id, fleet_id, device_id, device_name, tag, payload) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10)
                             ON CONFLICT (batch_id, id, timestamp_device) DO NOTHING;`)
	rows, err := ts.db.Query(q,
		batchId,
		data.Id,
		data.TimestampDevice,
		data.TimestampIn,
		data.WorkspaceID,
		data.FleetID,
		data.DeviceID,
		data.DeviceName,
		data.Tag,
		data.Payload)
	if rows != nil {
		defer rows.Close()
	}
	return err
}

// InsertCondition insert condition from devices into workspace tables
func (ts *ModelRepository) InsertCondition(data *models.Condition) error {
	q := fmt.Sprintf(`INSERT INTO condition (condition_id, workspace_id, device_id, tag, payload, payloadf, start, finish, duration, fleet_id, device_name) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11)
					ON CONFLICT (workspace_id, device_id, condition_id, start) DO UPDATE SET payloadf = $6, finish = $8, duration = $9;`)
	rows, err := ts.db.Query(q,
		data.ConditionID,
		data.WorkspaceID,
		data.DeviceID,
		data.Tag,
		data.Payload,
		data.PayloadF,
		data.Start,
		data.Finish,
		data.Duration,
		data.FleetID,
		data.DeviceName)
	if rows != nil {
		defer rows.Close()
	}
	return err
}

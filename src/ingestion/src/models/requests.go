package models

import (
	"encoding/json"
	"time"
)

type Data struct {
	Id              string          `pg:"id" json:"id"`
	TimestampDevice time.Time       `pg:"timestamp_device" json:"timestamp_device"`
	TimestampIn     time.Time       `pg:"ts_in" json:"timestamp_in"`
	Tag             string          `pg:"tag" json:"tag"`
	WorkspaceID     string          `pg:"workspace_id" json:"workspace_id"`
	DeviceID        string          `pg:"device_id" json:"device_id"`
	DeviceName      string          `pg:"device_name" json:"device_name"`
	FleetID         string          `pg:"fleet_id" json:"fleet_id"`
	Payload         json.RawMessage `pg:"payload" json:"payload"`
}

type Condition struct {
	ConditionID string          `pg:"condition_id" json:"uuid"`
	Tag         string          `pg:"tag" json:"tag"`
	WorkspaceID string          `pg:"workspace_id" json:"workspace_id"`
	DeviceID    string          `pg:"device_id" json:"device_id"`
	FleetID     string          `pg:"fleet_id" json:"fleet_id"`
	DeviceName  string          `pg:"device_name" json:"device_name"`
	Payload     json.RawMessage `pg:"payload" json:"payload"`
	PayloadF    json.RawMessage `pg:"payloadf" json:"payloadf"`
	Start       *time.Time      `pg:"start" json:"start"`
	Finish      *time.Time      `pg:"finish" json:"finish"`
	Duration    float64         `pg:"duration" json:"duration"`
}

type InsertDataRequest struct {
	BatchId string `json:"batch_id"`
	Result  []Data `json:"result"`
}

type InsertConditionRequest struct {
	IncomingCondition []Condition
}

func (n *InsertConditionRequest) UnmarshalJSON(p []byte) error {
	var tmp []Condition
	if err := json.Unmarshal(p, &tmp); err != nil {
		return err
	}
	n.IncomingCondition = tmp
	return nil
}

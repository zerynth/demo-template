
CREATE TABLE IF NOT EXISTS data (
    batch_id varchar(16) NOT NULL,
    id varchar(16) NOT NULL,
    timestamp_device timestamptz NOT NULL,
    timestamp_in timestamptz NOT NULL,
    tag varchar(32) NOT NULL,
    workspace_id varchar(16) NOT NULL,
    fleet_id varchar(16) NOT NULL,
    device_id varchar(16) NOT NULL,
    device_name varchar(32) NOT NULL,
    payload jsonb,
    created_at timestamptz NOT NULL DEFAULT CURRENT_TIMESTAMP,
    PRIMARY KEY (batch_id, id, timestamp_device)
);

CREATE INDEX timestamp_idx  on data(timestamp_device);
SELECT create_hypertable('data', 'timestamp_device');

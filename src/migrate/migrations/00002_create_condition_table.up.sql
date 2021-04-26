
CREATE TABLE IF NOT EXISTS condition (
    condition_id varchar(20) NOT NULL,
    tag varchar(64) NOT NULL,
    workspace_id varchar(16) NOT NULL,
    device_id varchar(16) NOT NULL,
    fleet_id varchar(16),
    device_name varchar(32),
    payload jsonb,
    payloadf jsonb,
    start timestamptz NOT NULL,
    finish timestamptz,
    duration float,
    created_at timestamptz NOT NULL DEFAULT CURRENT_TIMESTAMP,
    PRIMARY KEY (workspace_id, device_id, condition_id, start)
);

CREATE INDEX workspace_tag_start_idx ON condition (workspace_id, tag, start DESC);
SELECT create_hypertable('condition', 'start');
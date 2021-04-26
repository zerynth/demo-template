#!/bin/sh

echo "Running migrations"

/migrate -path /db/ -database postgres://postgres:password@tsdb:5432/postgres?sslmode=disable up

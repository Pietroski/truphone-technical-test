#!/usr/bin/env bash

sqlc generate -f internal/adaptors/datastore/postgresql/sqlc/auth/auth.sqlc.yaml
sqlc generate -f internal/adaptors/datastore/postgresql/sqlc/manager/devices/devices.sqlc.yaml

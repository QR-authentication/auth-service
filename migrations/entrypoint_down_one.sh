#!/bin/bash

goose -dir ./migrations postgres "user=$AUTH_SERVICE_POSTGRES_USER password=$AUTH_SERVICE_POSTGRES_PASSWORD dbname=$AUTH_SERVICE_POSTGRES_DB host=$AUTH_SERVICE_POSTGRES_HOST port=$AUTH_SERVICE_POSTGRES_PORT sslmode=disable" down

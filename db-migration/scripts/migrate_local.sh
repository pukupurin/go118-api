#!/bin/sh
cd `dirname $0`

URL="postgres://${DB_USERNAME}:${DB_PASSWORD}@${DB_HOST}:${DB_PORT}/${DB_NAME}?search_path=public&sslmode=disable"

migrate -source file://../migrations -database ${URL} "$@"

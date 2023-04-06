#!/bin/sh
cd `dirname $0`

URL="postgres://${DB_USERNAME}:${DB_PASSWORD}@${DB_HOST}:${DB_PORT}/${DB_NAME}?search_path=public&sslmode=disable"

atlas migrate apply --dir "file://../migrations" --url ${URL}

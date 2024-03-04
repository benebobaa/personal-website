#!/bin/sh

set -e

echo "run app"


/app/migrate -path /app/migrations -database "$DB_DSN" -verbose up


echo "start app"

exec "$@"
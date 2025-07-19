#!/bin/bash


# Docker PostgreSQL Setup Script
# Usage: ./setup_postgres.sh [db_name] [username] [password] (defaults: mydb, postgres, postgres)

set -e

# Set variables with defaults
DB_NAME=${1:-db}
DB_USER=${2:-postgres}
DB_PASSWORD=${3:-postgres}
PROJECT_DIR=$(pwd)
DATA_DIR="$PROJECT_DIR/database"
CONTAINER_NAME="postgres_${DB_NAME}"
PORT=5432

# Create data directory if it doesn't exist
mkdir -p "$DATA_DIR"

# Check if container already exists
if [ "$(docker ps -aq -f name=$CONTAINER_NAME)" ]; then
    echo "Container $CONTAINER_NAME already exists."
    read -p "Do you want to stop and remove it? (y/n) " -n 1 -r
    echo
    if [[ $REPLY =~ ^[Yy]$ ]]; then
        docker stop "$CONTAINER_NAME"
        docker rm "$CONTAINER_NAME"
    else
        echo "Aborting."
        exit 1
    fi
fi

echo "Starting PostgreSQL container..."
docker run --name "$CONTAINER_NAME" \
    -e POSTGRES_DB="$DB_NAME" \
    -e POSTGRES_USER="$DB_USER" \
    -e POSTGRES_PASSWORD="$DB_PASSWORD" \
    -p "$PORT":5432 \
    -v "$DATA_DIR":/var/lib/postgresql/data \
    -d postgres
# sleep 10
echo "The container started successfully"

echo "Initializing tables"
docker exec -i $CONTAINER_NAME psql -U $DB_USER -d $DB_NAME < init_tables.sql
echo "Done"

echo ""
echo "PostgreSQL container setup complete!"
echo "Connection details:"
echo "  Host: localhost"
echo "  Port: $PORT"
echo "  Database: $DB_NAME"
echo "  Username: $DB_USER"
echo "  Password: $DB_PASSWORD"
echo ""
echo "Data is stored in: $DATA_DIR"
echo "To connect using psql:"
echo "  docker exec -it $CONTAINER_NAME psql -U $DB_USER -d $DB_NAME"
echo "Or from host machine:"
echo "  PGPASSWORD=$DB_PASSWORD psql -h localhost -p $PORT -U $DB_USER -d $DB_NAME"

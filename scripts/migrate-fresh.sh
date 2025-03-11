#!/bin/bash

# Load environment variables from .env
set -a
source .env
set +a

# Function to check if a command was successful
check_error() {
    if [ $? -ne 0 ]; then
        echo "❌ Error: $1"
        exit 1
    fi
}

echo "🐸 Clearing database..."

# Format the database URL correctly for Turso
DB_URL="${TURSO_DATABASE_URL}?authToken=${TURSO_AUTH_TOKEN}"

# Run down migrations for seeds
echo "📝 Running seed migrations down..."
goose -dir seeds turso "${DB_URL}" down
check_error "Failed to run seed migrations down"

# Run down migrations for schema
echo "📝 Running schema migrations down..."
goose -dir migrations turso "${DB_URL}" down
check_error "Failed to run schema migrations down"

echo "✅ Database cleared successfully" 
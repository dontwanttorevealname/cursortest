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

echo "🐸 Setting up database..."

# Format the database URL correctly for Turso
DB_URL="${TURSO_DATABASE_URL}?authToken=${TURSO_AUTH_TOKEN}"

# Run up migrations for schema
echo "📝 Running schema migrations..."
goose -dir migrations turso "${DB_URL}" up
check_error "Failed to run schema migrations"

echo "🌱 Running seeds..."

# Run up migrations for seeds
echo "📝 Running seed migrations..."
goose -dir seeds turso "${DB_URL}" up
check_error "Failed to run seed migrations"

echo "✅ Database migrated and seeded successfully" 
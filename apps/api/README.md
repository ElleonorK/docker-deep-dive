# API Server

A Node.js REST API that connects to PostgreSQL database.

## What it does

- REST API on port 3000 (configurable via PORT env var)
- Connects to PostgreSQL database
- CRUD operations for items
- Health check and database connection check endpoints

## Endpoints

- `GET /health` - Health check
- `GET /db-check` - Test database connection
- `GET /items` - List all items
- `POST /items` - Create new item (send JSON: `{"name": "item name"}`)

## Environment Variables

Required:
- `DB_HOST` - Database hostname (default: localhost)
- `DB_PORT` - Database port (default: 5432)
- `DB_NAME` - Database name (default: postgres)
- `DB_USER` - Database user (default: postgres)
- `DB_PASSWORD` - Database password (default: mysecretpassword)

Optional:
- `PORT` - API port (default: 3000)

## Running locally (without Docker)

```bash
# Start postgres first
npm install
npm start
```

## Testing

```bash
# Health check
curl http://localhost:3000/health

# Database check
curl http://localhost:3000/db-check

# List items
curl http://localhost:3000/items

# Create item
curl -X POST http://localhost:3000/items \
  -H "Content-Type: application/json" \
  -d '{"name": "test item"}'
```

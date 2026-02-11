# Simple Web App

A minimal Node.js web application for Docker training exercises.

## What it does

- Exposes a REST API on port 8080 (configurable via PORT env var)
- Returns JSON with app info at `/`
- Health check endpoint at `/health`
- Shows version from APP_VERSION environment variable

## Running locally (without Docker)

```bash
npm install
npm start
```

Then visit http://localhost:8080

## Environment Variables

- `PORT` - Port to listen on (default: 8080)
- `APP_VERSION` - Version string to display (default: 'unknown')

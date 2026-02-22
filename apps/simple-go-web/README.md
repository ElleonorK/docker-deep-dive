# Simple Go Web App

A minimal Go web application for Docker training exercises.

## What it does

- Exposes a REST API on port 8080 (configurable via PORT env var)
- Returns JSON with app info at `/`
- Health check endpoint at `/health`
- Shows version from APP_VERSION environment variable
- Displays a message read from `message.txt`

## Running locally (without Docker)

### Option 1

Run using Go:
```bash
go run main.go
```

### Option 2

Build first:
```bash
go build -o simple-web
```

Then run the binary:
```bash
./simple-web
```

> [!NOTE]
> You need `message.txt` file for the app to run properly.

Visit http://localhost:8080

### Testing

For testing run:
```bash
./simple-web test
```

This will print the content of `message.txt` to console instead of starting a web server.

## Environment Variables

- `PORT` - Port to listen on (default: 8080)
- `APP_VERSION` - Version string to display (default: 'unknown')

## Why Go for Docker?

Go is a compiled language that produces standalone executables. Unlike interpreted languages, Go programs don't need the language runtime installed to run - just the compiled binary.

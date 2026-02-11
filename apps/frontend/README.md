# Frontend

A simple static HTML frontend for the Docker training application.

## What it does

- Single-page application with vanilla JavaScript
- Connects to the API backend
- Tests API health and database connectivity
- CRUD operations for items
- No build process required

## Running

Serve the static files with any web server. For example:

```bash
# Using Python
python3 -m http.server 8080

# Using Node.js http-server
npx http-server -p 8080

# Using nginx in Docker
docker run -v $(pwd):/usr/share/nginx/html:ro -p 8080:80 nginx
```

Then visit http://localhost:8080

## Configuration

The frontend expects the API to be available at `http://localhost:3000`. This is configured in the JavaScript code and can be modified if needed.

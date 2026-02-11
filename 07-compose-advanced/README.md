# Exercise 07: Docker Compose Advanced

Build a production-ready multi-container application with scaling, network isolation, and proper resource management.

## Objectives

Master advanced Docker Compose patterns:
- Service replication and scaling
- Multi-network architecture
- Different volume types for different use cases
- Resource limits and constraints
- Health checks and restart policies
- Production-ready configurations

## The Challenge

Build a complete application stack with:
- **Frontend** (simple-web) - user-facing web interface
- **Backend API** (api) - business logic and data access
- **Database** (postgres) - data storage
- **Cache** (redis) - session and data caching

## Tasks

### Task 1: Basic Stack

Create a docker-compose.yml that runs all four services. All services should start and be healthy.

Verify: `docker compose ps` shows all services running.

### Task 2: Scale the Frontend

The frontend needs to handle high traffic. Run 3 instances of the frontend service.

Verify: `docker compose ps` shows 3 frontend containers running.

### Task 3: Scale the Backend

Run 2 instances of the backend API service.

Verify: `docker compose ps` shows 2 API containers running.

### Task 4: Don't Scale the Database

Try to scale the database to 2 instances. It should fail or warn you.

Configure your compose file so the database can only ever run as a single instance.

### Task 5: Network Isolation - Frontend Network

Create a network architecture where:
- Frontend can reach Backend API
- Frontend CANNOT reach Database
- Frontend CANNOT reach Cache

Verify: `docker compose exec frontend ping database` should fail.

### Task 6: Network Isolation - Backend Network

Continue the network architecture:
- Backend API can reach Database
- Backend API can reach Cache
- Database and Cache are isolated from Frontend

Verify: `docker compose exec api ping database` should work.

### Task 7: Volume Types - Persistent Data

The database needs persistent storage that survives `docker compose down`.

Use a named volume for database data.

Verify: Create data, run `docker compose down`, run `docker compose up`, data persists.

### Task 8: Volume Types - Configuration

You need to provide a configuration file to the API that:
- Lives on your host machine
- Can be edited while containers are running
- Changes are immediately visible to the API

Use a bind mount for this.

Verify: Edit the config file on your host, check it inside the container - changes are instant.

### Task 9: Volume Types - Temporary Storage

The cache service needs fast temporary storage that:
- Doesn't need to persist
- Lives in memory
- Is fast for read/write operations

Use tmpfs for the cache data directory.

Verify: Create data in the tmpfs mount, restart the container, data is gone.

### Task 10: Resource Limits

Set resource limits:
- Frontend: max 256MB memory, 0.5 CPU
- Backend API: max 512MB memory, 1.0 CPU
- Database: max 1GB memory, 2.0 CPU
- Cache: max 128MB memory, 0.25 CPU

Verify: `docker stats` shows containers respecting these limits under load.

### Task 11: Restart Policies

Configure restart policies:
- Frontend: always restart
- Backend API: restart unless stopped manually
- Database: restart on failure only
- Cache: restart unless stopped manually

Test by stopping containers manually and via failures.

### Task 12: Health Checks

Add health checks to all services:
- Frontend: HTTP check on /health
- Backend API: HTTP check on /health
- Database: postgres connection check
- Cache: redis ping

Verify: `docker compose ps` shows health status for all services.

### Task 13: Startup Order

Services should start in the correct order and wait for dependencies to be healthy:
1. Database and Cache start first
2. Backend API waits for Database and Cache to be healthy
3. Frontend waits for Backend API to be healthy

Verify: Check logs - services start in correct order and wait for dependencies.

### Task 14: Environment-Specific Configs

Create three compose files:
- `docker-compose.yml` - base configuration
- `docker-compose.dev.yml` - development overrides (debug ports, verbose logging)
- `docker-compose.prod.yml` - production overrides (resource limits, no debug ports)

Run with: `docker compose -f docker-compose.yml -f docker-compose.prod.yml up`

### Task 15: Secrets Management

Database credentials should not be in plain text in the compose file or .env file.

Use Docker secrets to manage sensitive data.

Verify: Credentials are not visible in `docker compose config` output.

## Resources

- [Compose file reference](https://docs.docker.com/compose/compose-file/)
- [Networking in Compose](https://docs.docker.com/compose/networking/)
- [Volumes in Compose](https://docs.docker.com/compose/compose-file/07-volumes/)
- [Resource constraints](https://docs.docker.com/compose/compose-file/deploy/#resources)
- [Health checks](https://docs.docker.com/compose/compose-file/05-services/#healthcheck)
- [Secrets](https://docs.docker.com/compose/use-secrets/)

## Verification Checklist

You've completed this exercise when:
- ✓ All four services running in a stack
- ✓ Frontend scaled to 3 instances
- ✓ Backend scaled to 2 instances
- ✓ Database cannot be scaled (single instance only)
- ✓ Network isolation: frontend cannot reach database/cache
- ✓ Network isolation: backend can reach database/cache
- ✓ Named volume for persistent database data
- ✓ Bind mount for editable configuration
- ✓ tmpfs for temporary cache storage
- ✓ Resource limits configured and enforced
- ✓ Restart policies working correctly
- ✓ Health checks showing status
- ✓ Services start in correct order with health dependencies
- ✓ Environment-specific configurations working
- ✓ Secrets properly managed

This is a production-ready Docker Compose setup. Congratulations!

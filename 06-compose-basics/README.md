# Exercise 06: Docker Compose Basics

Learn to define and run multi-container applications using Docker Compose.

## Objectives

Master Docker Compose fundamentals:
- Write docker-compose.yml files
- Define services, networks, and volumes
- Use environment variables
- Manage multi-container applications
- Understand service dependencies

## Background

In exercise 03, you manually created networks, ran multiple containers, and connected them. That works, but it's tedious and error-prone.

Docker Compose lets you define all of that in a YAML file and manage it with simple commands.

## Tasks

### Task 1: Convert to Compose

Take your setup from exercise 03 (API + database) and convert it to a docker-compose.yml file.

Requirements:
- Both services should start with `docker compose up`
- The API should be accessible at http://localhost:3000
- The database should NOT be accessible from the host
- Everything should stop with `docker compose down`

### Task 2: Named Resources

Run `docker compose up` and check what Docker created:
- What are the container names?
- What's the network name?
- What's the volume name (if you added one)?

Notice a pattern? Docker Compose adds prefixes. Why?

### Task 3: Environment Variables

The API needs database credentials. Instead of hardcoding them in docker-compose.yml, use environment variables.

Create a `.env` file with:
```
DB_PASSWORD=supersecret
DB_NAME=myapp
```

Reference these in your docker-compose.yml.

Verify: Change the password in .env, restart, and confirm the API uses the new password.

### Task 4: Depends On

Right now, both services start simultaneously. But the API needs the database to be ready first.

Add a dependency so the API waits for the database.

Verify: Check the logs - database should start before API.

### Task 5: Persistent Data

Stop and remove everything with `docker compose down`.

Your database data is gone. Fix the compose file so data persists.

Verify: Create data, run `docker compose down`, run `docker compose up`, data is still there.

### Task 6: Multiple Environments

Create two compose files:
- `docker-compose.yml` - base configuration
- `docker-compose.override.yml` - development overrides (like port mappings for debugging)

Run `docker compose up` - it should automatically merge both files.

### Task 7: Scaling Services

Try to run 3 instances of the API:
```bash
docker compose up --scale api=3
```

What happens? Fix any issues so you can actually run multiple API instances.

Hint: Port conflicts are a problem when scaling.

### Task 8: Health Checks

Add a health check to the database service. The API should wait for the database to be healthy, not just started.

Verify: Check `docker compose ps` - it should show health status.

## Resources

- [Docker Compose overview](https://docs.docker.com/compose/)
- [Compose file reference](https://docs.docker.com/compose/compose-file/)
- [Environment variables in Compose](https://docs.docker.com/compose/environment-variables/)
- [Networking in Compose](https://docs.docker.com/compose/networking/)

## Verification Checklist

You've completed this exercise when:
- ✓ Multi-container app defined in docker-compose.yml
- ✓ Can start/stop entire stack with single commands
- ✓ Using environment variables from .env file
- ✓ Services start in correct order
- ✓ Data persists across compose down/up
- ✓ Can use override files for different environments
- ✓ Can scale services without port conflicts
- ✓ Health checks working properly

You're now ready for production-grade compose configurations!

## Exercise 07: Docker Compose Advanced

Build production-ready Docker Compose applications with health checks, resource management, and advanced networking.

## Objectives

By the end of this exercise, you'll be comfortable with:
* Implementing health checks and startup dependencies
* Configuring restart policies for resilience
* Setting resource limits and reservations
* Using different volume types for different use cases
* Building multi-network architectures
* Creating environment-specific configurations
* Implementing load balancing and service discovery

## Prerequisites

You should have completed exercise 06. You'll build on that knowledge to create production-grade configurations.

## The Mission

Take a basic compose stack and make it production-ready with proper health checks, resource management, resilience, and scalability.

## Tasks

### Task 1: Health Checks - Database

In exercise 06, you used `depends_on` but the API still showed connection errors at startup because postgres wasn't ready yet.

Add a health check to your database service. Postgres includes a `pg_isready` command that checks if the database is accepting connections.

Configure the health check with:
* A test command that runs `pg_isready`
* An interval (how often to check)
* A timeout (how long to wait for response)
* Retries (how many failures before marking unhealthy)
* A start period (grace period during startup)

Run `docker compose up` and watch `docker compose ps` - you should see the health status change from "starting" to "healthy".

**Think:** What values make sense for interval and timeout? Too frequent wastes resources, too infrequent delays startup.

### Task 2: Health Checks - Redis

Add a health check to your Redis service. Redis includes a `redis-cli ping` command that returns `PONG` when healthy.

The tricky part: `redis-cli ping` returns output, but health checks need an exit code. Figure out how to make the command return exit code 0 when healthy.

Verify: `docker compose ps` should show Redis health status.

**Think:** Why do health checks need exit codes instead of output? What exit code means healthy vs unhealthy?

### Task 3: Health Checks - API

Your API needs a health check too. Most APIs expose a `/health` endpoint that returns HTTP 200 when healthy.

But there's a problem - the API container might not have `curl` installed. You have a few options:
* Install curl in your API image
* Use `wget` if available
* Use a different health check method

Add a health check to your API service that verifies it's responding to HTTP requests.

**Think:** Should the health check just verify the API is running, or should it also check database connectivity? What's the tradeoff?

### Task 4: Startup Dependencies with Health Checks

Now that you have health checks, fix the startup order properly.

The API should wait for the database to be healthy (not just started) before starting.

Modify your `depends_on` to use the `condition: service_healthy` option.

Run `docker compose up` and watch the logs carefully. The API should start only after the database is healthy, with zero connection errors.

**Think:** What happens if the database health check never passes? Does the API ever start?

### Task 5: Complex Startup Dependencies

You now have API, database, and Redis. The API needs both database and Redis to be healthy before starting.

Configure dependencies so:
1. Database and Redis start first (they don't depend on anything)
2. API waits for BOTH database and Redis to be healthy
3. No connection errors in any logs

Verify: Run `docker compose up` and watch the startup sequence. Services should start in the correct order.

**Think:** What if you add a frontend that depends on the API? How deep can the dependency chain go?

### Task 6: Restart Policies - Understanding the Options

Services crash in production. Configure restart policies for resilience.

Try each policy and understand what it does:
* `no` - never restart
* `always` - always restart, even if stopped manually
* `on-failure` - only restart if the container exits with an error
* `unless-stopped` - always restart unless explicitly stopped

Set your database to `unless-stopped` and test it:
1. Start your stack
2. Kill the database: `docker compose kill database`
3. Watch it restart automatically
4. Stop the database: `docker compose stop database`
5. Restart the entire stack: `docker compose restart`
6. Check if the database started - it shouldn't have

**Think:** What's the difference between killing and stopping? When would you use each restart policy?

### Task 7: Restart Policies - On Failure with Limits

Set your API to restart `on-failure` but with a maximum of 3 attempts.

Test it by making the API crash repeatedly:
1. Start your stack
2. Kill the API: `docker compose kill api`
3. Watch it restart
4. Kill it again immediately
5. Keep killing it rapidly

After 3 restarts, it should give up. Check `docker compose ps` - the API should be in an exited state.

**Think:** Why limit restart attempts? What happens in production if a service is misconfigured and crashes immediately on startup?

### Task 8: Resource Limits - Memory

Prevent services from consuming all system memory. Set memory limits:
* Database: 512MB
* Redis: 256MB
* API: 256MB

Start your stack and monitor resource usage:
```bash
docker stats
```

Now try to make the API consume more than 256MB of memory. What happens when it hits the limit?

**Think:** What happens when a container hits its memory limit? Does it slow down or get killed?

### Task 9: Resource Limits - CPU

Set CPU limits:
* Database: 1.0 CPU (100% of one core)
* Redis: 0.5 CPU (50% of one core)
* API: 0.5 CPU (50% of one core)

Monitor with `docker stats` and observe the CPU percentages.

Try to make the API do CPU-intensive work. Can it exceed 50% of one core?

**Think:** What's the difference between CPU limits and memory limits? What happens when you hit each?

### Task 10: Resource Reservations

Limits are maximums, but you can also set reservations (guaranteed minimums).

Set both limits and reservations:
* Database: 512MB limit, 256MB reserved
* Redis: 256MB limit, 128MB reserved
* API: 256MB limit, 128MB reserved

**Think:** When would reservations matter? What happens if you try to reserve more resources than your system has?

### Task 11: Volume Types - Named Volumes

You're already using named volumes for persistent data. But let's understand them better.

Create a named volume explicitly in your compose file (in the top-level `volumes:` section) instead of letting Compose create it automatically.

Give it a custom name without the project prefix.

Verify: `docker volume ls` should show your custom-named volume.

**Think:** When would you want to control the volume name? What if multiple projects need to share data?

### Task 12: Volume Types - Bind Mounts

You need to provide a configuration file that can be edited while containers are running.

Create a `redis.conf` file:
```
maxmemory 256mb
maxmemory-policy allkeys-lru
save 60 1000
```

Mount this file into your Redis container using a bind mount (not a volume).

Verify: Edit `redis.conf` on your host, then check inside the container - changes should appear instantly without restarting.

**Think:** What's the difference between a bind mount and a volume? When would you use each?

### Task 13: Volume Types - Tmpfs

Redis needs fast temporary storage for some operations. Tmpfs mounts live in memory and are very fast but don't persist.

Add a tmpfs mount to your Redis container at `/tmp/redis-tmp`.

Verify: 
1. Exec into Redis and create a file: `docker compose exec redis sh -c "echo test > /tmp/redis-tmp/file.txt"`
2. Check it exists: `docker compose exec redis cat /tmp/redis-tmp/file.txt`
3. Restart just the Redis container: `docker compose restart redis`
4. Check again - the file should be gone

**Think:** When would you use tmpfs? What's the tradeoff between speed and persistence?

### Task 14: Multi-Network Architecture

Build a realistic network architecture with proper isolation:

Services:
* Frontend (nginx serving static files)
* API (your existing API)
* Database (postgres)
* Cache (redis)

Networks:
* `frontend-network` - frontend and API
* `backend-network` - API, database, and cache

Requirements:
* Frontend can reach API only
* API can reach database and cache
* Database and cache cannot reach each other
* Frontend cannot reach database or cache

Verify each isolation rule by trying to ping between services.

**Think:** Why isolate database and cache from each other? What security principle is this?

### Task 15: Load Balancing with Nginx

Scale your API to 3 instances, but users should only hit one endpoint.

Add an nginx load balancer service that:
* Sits in front of the API
* Distributes traffic across all API instances using round-robin
* Is the only service exposed to the host

Create an `nginx.conf` file:
```nginx
upstream api_backend {
    server api:3000;
}

server {
    listen 80;
    location / {
        proxy_pass http://api_backend;
        proxy_set_header Host $host;
        proxy_set_header X-Real-IP $remote_addr;
    }
}
```

But this only points to one API instance. Research how nginx discovers all instances when you scale the API service.

Verify: Hit the load balancer multiple times and check logs - different API containers should handle requests.

**Think:** How does nginx discover all API instances? What happens when you scale up or down?

### Task 16: Environment-Specific Configurations

Create three compose files:
* `docker-compose.yml` - base configuration (all service definitions)
* `docker-compose.dev.yml` - development overrides
* `docker-compose.prod.yml` - production overrides

Development should:
* Expose all ports for debugging (database, redis, etc.)
* Disable resource limits
* Use `restart: "no"` for faster iteration
* Mount source code for live reload
* Enable verbose logging

Production should:
* Only expose the load balancer port
* Enable all resource limits
* Use `restart: unless-stopped`
* No source code mounts
* Minimal logging

Test both:
```bash
docker compose -f docker-compose.yml -f docker-compose.dev.yml up
docker compose -f docker-compose.yml -f docker-compose.prod.yml up
```

**Think:** How would you handle secrets differently in each environment? What about SSL certificates?

### Task 17: Init Containers Pattern

Your database needs to run initialization scripts on first startup.

Create an `init.sql` file:
```sql
CREATE TABLE IF NOT EXISTS users (
    id SERIAL PRIMARY KEY,
    username VARCHAR(255) UNIQUE NOT NULL,
    email VARCHAR(255) UNIQUE NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE INDEX idx_users_username ON users(username);
CREATE INDEX idx_users_email ON users(email);

INSERT INTO users (username, email) VALUES 
    ('admin', 'admin@example.com'),
    ('user1', 'user1@example.com');
```

Mount this into the postgres container so it runs automatically on first startup.

Verify: Start fresh (remove volumes), bring up the stack, then check if the table and data exist.

Now modify `init.sql` to add another table and restart (without removing volumes). Does the new table appear?

**Think:** When do init scripts run? How would you handle schema migrations in production?

### Task 18: Graceful Shutdown

Configure your services to shut down gracefully:
* Set appropriate `stop_grace_period` for each service
* Database needs time to flush data (30 seconds)
* API needs time to finish requests (10 seconds)
* Redis needs time to save data (10 seconds)

Test: Start your stack, create some data, then run `docker compose down`. Watch the logs - services should shut down cleanly without errors.

**Think:** What happens if a service doesn't stop within the grace period? How long is too long?

### Task 19: Logging Configuration

Configure logging for production:
* Set log driver to `json-file`
* Limit log size to 10MB per container
* Keep only the last 3 log files

This prevents logs from filling up disk space.

Verify: Check the logging configuration with `docker inspect <container-name>`.

**Think:** What happens when logs exceed the size limit? Where do logs go in production?

### Task 20: Health-Based Load Balancing

Your nginx load balancer should only send traffic to healthy API instances.

Add health checks to your load balancer configuration so it:
* Checks each API instance's `/health` endpoint
* Removes unhealthy instances from the pool
* Adds them back when they become healthy

Test: Kill one API instance and verify the load balancer stops sending traffic to it.

**Think:** How does this improve reliability? What happens if all instances become unhealthy?

## Resources

* [Compose file reference](https://docs.docker.com/compose/compose-file/)
* [Health checks](https://docs.docker.com/engine/reference/builder/#healthcheck)
* [Resource constraints](https://docs.docker.com/compose/compose-file/deploy/)
* [Networking in Compose](https://docs.docker.com/compose/networking/)
* [Volumes in Compose](https://docs.docker.com/compose/compose-file/07-volumes/)
* [Logging configuration](https://docs.docker.com/config/containers/logging/configure/)

## Verification Checklist

You've completed this exercise when:
* [ ] All services have working health checks
* [ ] Services start in correct order based on health
* [ ] Restart policies configured and tested
* [ ] Resource limits (CPU and memory) set and enforced
* [ ] Resource reservations configured
* [ ] Using named volumes, bind mounts, and tmpfs appropriately
* [ ] Multi-network architecture with proper isolation
* [ ] Load balancer distributing traffic across scaled services
* [ ] Separate dev and prod configurations working
* [ ] Database initialization scripts running on first startup
* [ ] Graceful shutdown configured
* [ ] Logging configured to prevent disk space issues
* [ ] Health-based load balancing working

## Cleanup

Stop and remove all resources:
```bash
docker compose down -v
```

Remove custom networks:
```bash
docker network prune
```

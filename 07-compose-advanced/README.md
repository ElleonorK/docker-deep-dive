# Exercise 07: Docker Compose Advanced

Master production-ready Docker Compose configurations with health checks, resource management, and operational best practices.

## Objectives

By the end of this exercise, you'll be comfortable with:
* Implementing health checks and smart startup dependencies
* Configuring restart policies for resilience
* Managing resource limits and reservations
* Controlling logging to prevent disk space issues
* Using profiles for optional services
* Implementing graceful shutdown
* Using extension fields to keep configs DRY
* Understanding volume options and drivers

## Prerequisites

You should have completed exercise 06 and have a working compose stack with api, database, and simple-go-web services.

## The Mission

Transform a basic compose stack into a production-ready configuration with proper health checks, resource management, and operational resilience.

## Tasks

### Task 1: Health Checks

In exercise 06, you used `depends_on` but the API still showed connection errors at startup because postgres wasn't ready yet.

The API has two endpoints: `/health` returns a simple status check, and `/db-check` actually tests the database connection. These represent two different concepts in container orchestration.

First, let's make the problem more visible. Modify your database service to start slowly by adding a startup delay:

```yaml
command: sh -c "sleep 60 && docker-entrypoint.sh postgres"
```

Start your stack. Watch the API logs - you'll see connection errors for the first minute.

**Your challenge:** Add health checks to both services. Check the postgres documentation and API README to understand what commands or endpoints are available for testing the health of your containers.

Configure the API to wait for the database to be healthy before starting using `depends_on` with conditions.

Start your stack. Does the API still show connection errors? If yes, why? The database is healthy, so what's missing?

Try using the API's `/db-check` endpoint for its health check instead. What changes?

**Success criteria:** Run `docker compose up` and watch the logs. The API should start only after the database is healthy, with zero connection errors. Check `docker compose ps -w` to see health status transition from "starting" to "healthy".

**Think:** What's the difference between application running vs an application being ready to serve traffic? Docker Compose doesn't differentiate between the two - which one is more important in your case?

### Task 2: When Containers Crash

Containers don't always run forever. Sometimes they crash because they run out of resources, sometimes because of bugs, sometimes because of external factors.

**Your challenge:** Configure aggressive resource limits that will cause problems:
* Database: 128MB memory limit, 64MB reserved
* API: 16B memory limit, 8MB reserved

Start your stack and watch what happens. Monitor with `docker stats` in another terminal.

Visit the API endpoints a few times. Create some database tables. What happens to the containers?

Check the container status with `docker compose ps`. What do you see? Are they running, or something else?

Now look at the logs for whichever container crashed. What does the last message say?

**Success criteria:** You can make at least one container crash by hitting its resource limit. You understand what "OOMKilled" means.

**Think:** What happens to your application when a container suddenly dies? What about the data that was in memory? What if this was handling a user's payment?

### Task 3: Graceful Shutdown and Automatic Recovery

In Task 2, you saw containers crash suddenly. That's not ideal. Let's learn how to handle shutdowns gracefully and recover automatically.

First, remove those aggressive resource limits from Task 2. Set reasonable limits instead:
* Database: 512MB memory limit, 256MB reserved
* API: 256MB memory limit, 128MB reserved

**Part A: Understanding Shutdown**

Start your stack. Now run `docker compose down` and watch the logs carefully.

What do you see? Do the containers stop immediately or do they take time? Do you see any "killed" messages?

Try stopping just the database:
```bash
docker compose stop database
```

How long does it take? Now check the API logs - what happened to it when the database disappeared?

**Your challenge:** Research what signal Docker sends when stopping a container. Find out how much time containers get before they're force-killed. Can you give the database more time to shut down cleanly?

**Part B: Automatic Restart**

Start your stack again. Now simulate a crash by killing the database:
```bash
docker compose kill database
```

What happens? Check `docker compose ps`. Is the database running?

**Your challenge:** Configure restart policies so containers recover automatically from crashes. Research what options exist and configure:
* Database: Always restart when it crashes, but stay stopped if you explicitly stop it
* API: Restart on failures, but give up after 3 attempts

Test your configuration:
* Kill the database with `docker compose kill database` - does it come back?
* Stop the database with `docker compose stop database` - does it come back?
* Kill the API repeatedly - does it eventually give up?

**Part C: The Cascade Effect**

With your restart policies configured, kill the database again. Watch what happens to the API.

The API can't connect to the database anymore. What does it do? Does it crash? Does it keep trying?

Now wait for the database to restart. Does the API recover automatically, or is it stuck in a bad state?

**Your challenge:** Make the API resilient to database restarts. It should automatically reconnect when the database comes back. You might need to modify the restart policy or add health checks.

**Success criteria:**
* Database takes at least 30 seconds to stop gracefully (you configured this)
* Database automatically restarts after being killed
* Database stays stopped when you explicitly stop it
* API gives up after 3 failed restart attempts
* Both services recover when the database restarts

**Think:** What's the difference between a crash and an explicit stop? Why would you want different behavior for each? What happens to in-flight requests when a container is killed vs stopped gracefully?

### Task 4: Logging Configuration

You want to run different services in different scenarios without maintaining separate compose files.

**Your challenge:** Add a Redis cache service, but make it optional using profiles:
* Default profile: runs database and API only
* `cache` profile: also runs Redis
* `full` profile: runs everything including a monitoring tool (use `prom/prometheus` image)

Test each profile:
```bash
docker compose up                    # default
docker compose --profile cache up   # with cache
docker compose --profile full up    # everything
```

**Success criteria:** You can start different combinations of services without editing the compose file.

**Think:** When would you use profiles vs separate compose files? What are real-world scenarios for this?

### Task 7: Extension Fields for DRY Configs

You're repeating the same configuration across multiple services. Extension fields let you define common config once and reuse it.

**Your challenge:** Create extension fields for:
* Common logging configuration
* Common health check settings
* Common resource limits

Then reference these in your services using anchors and aliases.

**Success criteria:** Change the logging config in one place and it applies to all services.

**Think:** What's the tradeoff between DRY configs and readability? When does this become more confusing than helpful?

### Task 8: Volume Options and Drivers

Volumes have options beyond just mounting paths. Explore what's possible.

**Your challenge:** Configure your database volume with specific options:
* Set the volume to use the `local` driver explicitly
* Add a label to the volume for documentation
* Make the volume external (created outside compose)

Create the external volume first:
```bash
docker volume create --label project=training postgres-data
```

Then reference it in your compose file.

**Success criteria:** `docker volume inspect postgres-data` shows your custom configuration. The volume persists even after `docker compose down -v`.

**Think:** When would you use external volumes? What other volume drivers exist? What happens to external volumes when you run `docker compose down -v`?

## Resources

* [Compose file reference](https://docs.docker.com/compose/compose-file/)
* [Health checks](https://docs.docker.com/engine/reference/builder/#healthcheck)
* [Depends on with conditions](https://docs.docker.com/compose/compose-file/05-services/#depends_on)
* [Restart policies](https://docs.docker.com/compose/compose-file/05-services/#restart)
* [Resource constraints](https://docs.docker.com/compose/compose-file/deploy/)
* [Logging configuration](https://docs.docker.com/config/containers/logging/configure/)
* [Profiles](https://docs.docker.com/compose/profiles/)
* [Extension fields](https://docs.docker.com/compose/compose-file/11-extension/)
* [Volume configuration](https://docs.docker.com/compose/compose-file/07-volumes/)

## Verification Checklist

You've completed this exercise when:
* [ ] Services have health checks and start in correct order
* [ ] Restart policies configured and tested
* [ ] Resource limits prevent services from consuming all system resources
* [ ] Logging configured to prevent disk space issues
* [ ] Profiles let you run different service combinations
* [ ] Extension fields reduce configuration repetition
* [ ] Volume options and external volumes working
* [ ] Graceful shutdown configured
* [ ] Can explain when to use each feature in production

## Cleanup

Stop and remove all resources:
```bash
docker compose down -v
```

Remove external volumes if created:
```bash
docker volume rm postgres-data
```

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

Let's see what happens when a container runs out of memory.

**Your challenge:** Open two terminals (or split your terminal). In the first terminal, run:

```bash
docker stats
```

Leave this running so you can watch memory usage in real-time.

In the second terminal, start a memory-hungry container using the file in the exercise directory:

```bash
docker compose -f memory-hog.yml up
```

Watch both terminals. What do you see happening to the memory usage? What happens to the logs? How long does it take before something goes wrong?

After it's done, run:

```bash
docker compose -f memory-hog.yml ps -a
```

What's the exit code? What does it mean?

**Think:** 
* What happens to your application when a container suddenly dies? What about the data that was in memory? What if this was handling a user's payment?
* Compare the numbers you see in the logs vs what `docker stats` showed. Do they match? If not, why might that be?

### Task 3: Graceful Shutdown and Automatic Recovery

What you just witnessed in Task 2 is called **OOMKilled** - the Linux kernel's Out-Of-Memory killer terminated your container when it exceeded its memory limit.

**Part A: Stopping the Bleeding**

**Your challenge:** Modify `memory-hog.yml` to prevent the OOMKilled crash in two different ways:

1. The obvious way - adjust the resources
2. The risky way - research if Docker has an option to tell the kernel to leave your container alone, no matter how much memory it consumes

Try both. What happens with each approach?

**Think:** Would that second option ever make sense in production? What's the worst that could happen?

**Part B: How Containers Actually Die**

OOMKilled is just one way containers stop. Let's explore others.

Run a simple container that sleeps for 30 seconds:

```bash
docker run --name sleeper alpine sh -c "echo 'Running...'; sleep 30; echo 'Done'"
```

Let it finish. Check `docker ps -a` - what's the exit code?

Now run one that fails naturally:

```bash
docker run --name failer alpine cat /nonexistent-file
```

What's the exit code this time? Why is it different?

Start another sleeper, but this time press Ctrl+C to interrupt it:

```bash
docker run --name interrupted alpine sleep 30
# Press Ctrl+C (you might need to press it 2-3 times)
```

Check the exit code. What is it?

Now start a sleeper in detached mode and stop it with Docker:

```bash
docker run -d --name sleeper2 alpine sh -c "echo 'Running...'; sleep 30; echo 'Done'"
docker stop sleeper2
```

Watch how long it takes. Check `docker ps -a` for the exit code. Check the logs with `docker logs sleeper2` - did it finish or get interrupted?

**Think:** Have you seen that exit code before? What's the connection between `docker stop` and what happened in Task 2?

**Your challenge:** The stop took a specific amount of time before forcing the kill. Can you make it wait longer? Look for docker run options that control how containers are stopped.

**Part C: Automatic Recovery**

Back to the memory-hog. When it crashed in Task 2, it just died and stayed dead. In production, you'd want critical services to recover automatically from crashes.

**Your challenge:** Modify `memory-hog.yml` so the container automatically restarts when it crashes.

Start it and watch. What happens after it crashes? Does it come back? How many times? Check `docker compose ps` while it's running - what information do you see about restarts?

Let it crash a few times, then stop it yourself with `docker compose stop memory-hog`. Wait a moment. Did it restart? Why or why not?

Now imagine this scenario: your container has a bug that makes it crash immediately on startup. You don't want it restarting forever and consuming resources. 

**Your challenge:** Configure it to give up after 3 failed attempts instead of restarting indefinitely.

Test it - does it stop trying after 3 crashes?

**Think:** What's worse in production - a service that crashes and stays down, or one that keeps crashing and restarting forever?

**Part D: Graceful Shutdown**

If you haven't already, remove the resource constraints from `memory-hog.yml` so it doesn't crash.

**Time To Stop**

Start the memory-hog stack.

Stop the container and time how long it takes:

```bash
time docker compose -f memory-hog.yml stop memory-hog
```

**Think:** If this were a critical container, would this time be enough? What repercussions can there be if it's not?

**Your challenge:** Change the memory-hog.yml to wait for 10 more seconds before stopping the container.

**Success Criteria:** You see a real time longer than 10 seconds when running:

```bash
time docker compose -f memory-hog.yml stop memory-hog
```

**Kill Bill 2**

Start the container again.

Now KILL the undead:

```bash
time docker compose -f memory-hog.yml kill memory-hog
```

How long did it take to kill? Hope you learned the difference between mercy and show of force.

**Bonus:** Create the stack again, and kill it using classic `docker stop`.
 


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

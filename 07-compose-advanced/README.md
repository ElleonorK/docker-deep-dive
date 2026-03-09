# Exercise 07: Docker Compose Advanced

Master production-ready Docker Compose configurations with health checks, resource management, and operational best practices.

## Objectives

By the end of this exercise, you'll be comfortable with:
* Coordinating service startup and ensuring readiness
* Making services resilient to failures and resource constraints
* Implementing graceful shutdown and cleanup
* Securing containers with user isolation and filesystem restrictions
* Managing configuration and sensitive data declaratively
* Running different service combinations for different environments

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
 
**The Language of Signals**

When you ran `docker stop`, Docker sent a signal to the process inside the container. By default, it sends SIGTERM - a polite request to shut down. If the process doesn't respond within the grace period, Docker sends SIGKILL - which can't be ignored.

But not all applications speak the same language. Some expect different signals.

**Your challenge:** Look up what signal Docker Compose uses by default when stopping containers. Can you change it? Try configuring `memory-hog.yml` to send SIGINT (the same signal Ctrl+C sends) instead.

Test it:
```bash
docker compose -f memory-hog.yml up -d
docker compose -f memory-hog.yml stop memory-hog
```

Check the logs. Did anything change? Python's default behavior handles both signals similarly, but does your configuration work?

**Think:** When would you need to change the stop signal? What if your application only handles SIGINT? What about SIGHUP?

**Last Words**

Sometimes you need to run cleanup commands when your container receives the stop signal, before it fully dies.

**Think:** What could go wrong if your container dies without a chance to clean up properly? What state might be left behind?

**Your challenge:** Modify `memory-hog.yml` so that when the stop signal is received, it echoes "Goodbye, cruel world!" 

Start the container, then stop it:
```bash
docker compose -f memory-hog.yml up -d
docker compose -f memory-hog.yml stop memory-hog
```

**Success criteria:** You can see "Goodbye, cruel world!" in the container logs when it shuts down.

**Think:** When would you use this in production? What's the difference between handling shutdown in your application code vs running a separate command during shutdown?

### Task 4: Who Runs This Thing Anyway?

Your security team mandates all containers run as non-root. Configure memory-hog to run as UID 5000.

Add a volume mount `./hog-data:/data` and have the container create some files:

```bash
docker compose -f memory-hog.yml exec memory-hog sh -c 'echo "data" > /data/output.txt'
docker compose -f memory-hog.yml exec memory-hog sh -c 'echo "more" > /data/results.txt'
```

Now try to delete those files from your host:

```bash
rm hog-data/output.txt
```

What happens? Why?

**Your challenge:** Fix it so you can manage these files from your host, while still using a non-root user.

**Think:** Why shouldn't you do this in production? You run with user 5000, but what would happen if you run with 1000?

Now let's explore granular file permissions. Create three config files:

```bash
mkdir -p hog-config
echo "SHARED=true" > hog-config/shared.conf
echo "HOST=true" > hog-config/host-only.conf
echo "CONTAINER=true" > hog-config/container-only.conf
```

**Your challenge:** Mount these three files individually in `memory-hog.yml` with different permissions so that:
* `shared.conf` - both host and container (UID 5000) can modify
* `host-only.conf` - only host can modify, container can only read
* `container-only.conf` - only container can modify, host can only read

Test by trying to modify each file from both the host and container.

**Bonus:** Learn how to translate numeric permission codes into letter format (like `rwxr-xr--`).

Now try accessing the files as root from inside the container:

```bash
docker compose -f memory-hog.yml exec --user root memory-hog sh -c 'echo "I AM GROOT" >> /config/host-only.conf'
```

Did it work? Why or why not?

**Think:** Where did you learn how to prevent containers from running as root in the first place? (Hint: it wasn't in the compose file)

Your container's filesystem is fully writable. Try creating files in different locations:

```bash
docker compose -f memory-hog.yml exec memory-hog touch /tmp/test1
docker compose -f memory-hog.yml exec memory-hog touch /usr/local/test2
docker compose -f memory-hog.yml exec memory-hog touch /test3
```

All of these work. If an attacker compromises your container, they could modify files, create backdoors, or install malicious tools.

**Your challenge:** Lock down the filesystem so nothing can be written anywhere. The memory-hog script should still run successfully.

After making the change, test by trying to create files in those same locations. What happens? Does the script still run?

Run this command that modifies memory-hog to write into a file instead:

```bash
sed -i "s/print(f'🧠 Allocated {i\*5}MB...')/open('\/tmp\/memory.log', 'a').write(f'Allocated {i*5}MB\\\\n')/" memory-hog.yml
```

Does the container start now? Why?

**Your challenge:** While keeping the filesystem locked, find a way to give the container user write permissions specifically to the `/tmp` directory.

**Success criteria:** Container runs and writes logs to `/tmp/`, but you cannot create files anywhere else.

**Think:** An attacker compromises your container. With all these layers - non-root user, file permissions, locked filesystem - what can they still do? What can't they do anymore?

You've built multiple security layers. But there's one configuration that bypasses everything.

First, check what devices your container can see:

```bash
docker compose -f memory-hog.yml exec memory-hog ls /dev
```

Note how limited it is. Now add `privileged: true` to your memory-hog service and restart it. Run the same command:

```bash
docker compose -f memory-hog.yml exec memory-hog ls /dev
```

Where did all these additional devices come from? Figure it out.

**Success criteria:** You should start feeling very uneasy, maybe even scared 😰

**Think:** When would you legitimately need privileged mode? What's the risk if you use it unnecessarily?

> [!CAUTION]
> Never use privileged mode unless you are 100% sure what you are doing.

Remove `privileged: true` from your configuration.

### Task 5: The Final Boss Level

You've been through crashes, lock downs, zombie resurrections and intruders at your gates. You've learned a lot, but this doesn't cover everything Docker Compose can do - NOT EVEN CLOSE!

Best of all, new features keep getting added, and there's always more to explore.

**Your Final Quest ⚔️**

Time to bring what you've learned here to your full stack and add a few more production patterns.

**Part A: Practice Makes Perfect**

Switch back to your main stack from exercise 06 (api, database, simple-go-web).

**Your challenge:** Add to your main stack:
* Health checks for all services
* Restart policies
* Resource limits
* Non-root users where appropriate
* Volume permissions
* Read-only filesystems where possible

Start your stack and verify everything works. Check that services start in the correct order and restart automatically if they crash.

**Think:** Which services can run with read-only filesystems? Which need writable space?

**Part B: Better Composition**

Your simple-go-web app displays a message from `message.txt`. Currently, this file is copied into the image during build. Want to change the message? Rebuild the entire image. That's annoying.

**Your challenge:** Use Docker Compose configs to manage `message.txt` declaratively. Define it in your compose file and mount it into the container.

After making the change, modify the message and restart the service. The new message should appear without rebuilding.

**Success criteria:** You can change the message by editing your compose file and restarting, no rebuild needed.

**Think:** What's the advantage over copying files during build? When would you still want to copy files into the image?

**Part C: Stop Being So Obvious**

Your database password is just sitting there, in plain text for all to see. Your security team is not amused.

**Your challenge:** Move the database password to a Docker secret. Both the database and API need to read it.

**Success criteria:** The password doesn't appear in plain text in your compose file. The API connects to the database successfully.

>[!Note]
> You might need to modify the code a bit.

**Think:** What if someone gets shell access to your container? Can they still read the secret? How is this better than environment variables?

**Part D: To Each Their Own**

You have different teams working on different services and combinations.
Frontend developers work on simple-go-web and api. Backend developers only run the database. QA need to test connection between api and database. And sometimes you only need one.

**Your challenge:** Use profiles to make services optional:
* Default: all services run
* `frontend`: simple-go-web + api only
* `backend`: database only
* `goapp`: simple-go-web only
* `nodeapp`: api + database only

Test different combinations:
```bash
docker compose up                        # everything
docker compose --profile frontend up    # frontend stack
docker compose --profile goapp up       # just Go app
```

**Success criteria:** You can start different service combinations without editing the compose file.

**Think:** When would you use profiles vs separate compose files? What happens when you need a service in multiple profiles?

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
cd ../06-compose-basics
docker compose down -v
cd ../07-compose-advanced
docker compose -f memory-hog.yml down -v
```

Clean up any test directories created:
```bash
rm -rf hog-data hog-config
cd ../06-compose-basics
rm -rf api-logs api-config
```

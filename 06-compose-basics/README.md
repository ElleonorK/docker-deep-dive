## Exercise 06: Docker Compose Basics

Learn to orchestrate multi-container applications using Docker Compose.

## Objectives

By the end of this exercise, you'll be comfortable with:
* Defining multi-container applications in docker-compose.yml
* Building images with Compose using build context and arguments
* Scaling services and understanding replica behavior
* Using multiple compose files and overrides
* Managing container lifecycle with Compose commands
* Understanding the difference between Docker and Compose operations

## Prerequisites

You should have completed exercises 01-05. You'll need all three applications from the `apps/` directory, and you should have images already built for them from previous exercises.

## Tasks

### Task 1: The Big Challenge - Recreate Your Multi-Container Setup

In exercises 02-04, you manually connected containers with networks, exposed ports, and persisted data. You ran multiple commands to set everything up. Now do it all with one compose file.

Build a complete stack with three applications using your already-built images:
* **simple-go-web** - the Go web application (use your existing image)
* **api** - the Node.js API that connects to a database (use your existing image)
* **database** - postgres 16-alpine

Requirements:

**Port Exposure:**
* simple-go-web accessible at http://localhost:8080
* api accessible at http://localhost:3000
* database NOT accessible from host

**Network Isolation:**
* simple-go-web should NOT be able to reach the database
* api should be able to reach the database
* You'll need to define multiple networks explicitly

**Data Persistence:**
* Database data must survive `docker compose down` and `docker compose up`

Start everything with `docker compose up`. Verify:
* Go app is accessible at http://localhost:8080
* API is responding http://localhost:3000 and connects to the database
* Data in the database is persistent:
    * Create test data in the database:
    ```bash
    docker compose exec database psql -U postgres -c "CREATE TABLE test (id INT, name TEXT);"
    docker compose exec database psql -U postgres -c "INSERT INTO test VALUES (1, 'Alice');"
    ```
    * Destroy the stack and start it again
    * Query the data again - it should still be there:
    ```bash
    docker compose exec database psql -U postgres -c "SELECT * FROM test;"
    ```
* Try to connect to postgres from your host on port 5432 - should fail
* Exec into simple-go-web and try to ping database - should fail
* Exec into api and ping database - should work

### Task 2: Building Images with Compose

Up till now you used images already built. Will you have to build it yourself every time there is a small change?

Modify your compose file to build the images instead of using existing ones:
* simple-go-web should build from `apps/simple-go-web/`
* api should build from `apps/api/`

But here's the challenge - use build arguments:
* For simple-go-web, pass `APP_VERSION` as a build argument (set it to "2.0.0")
* For api, pass `NODE_VERSION` as a build argument (set it to "20")

You'll need to modify the Dockerfiles to accept these build arguments. Check the Dockerfiles and add `ARG` instructions where needed.

Remove your old images and run `docker compose up`. Compose should build the images automatically.

Verify:
* Visit http://localhost:8080 - the version should show "2.0.0"
* Exec into the api container and check Node version: `docker compose exec api node --version`

Now change the `APP_VERSION` to "3.0.0" in your compose file and run `docker compose up` again. Does it rebuild? Why or why not?

Try `docker compose up --build`. What's different?

**Think:** What's the difference between build arguments and environment variables? When is each set? What's the build context and why does it matter?

### Task 3: Scaling and Replicas

You want to run multiple instances of your services for high availability.

Try scaling the simple-go-web service to 3 replicas:
```bash
docker compose up --scale simple-go-web=3
```

What happens? You probably get a port conflict error.

Fix your compose file so you can scale simple-go-web. You'll need to change how ports are mapped.

Once it works, scale to 3 instances and list your containers:
```bash
docker compose ps
```

What are the containers named? Notice the pattern?

Now try scaling the database to 2 replicas:
```bash
docker compose up --scale database=2
```

Both database containers start, but what happens to your data? Create some data, then stop one database replica. Where did the data go?

**Think:** Why is scaling databases problematic? What happens when multiple database instances try to use the same volume? When would you want to scale a service vs when wouldn't you?

### Task 4: Multiple Compose Files and Overrides

Your compose file is getting complex. Split it up and learn about composition.

Create two separate compose files:

**docker-compose.yml** - Your base configuration with api and database only (remove simple-go-web)

**docker-compose.web.yml** - Just the simple-go-web service

Run both together:
```bash
docker compose -f docker-compose.yml -f docker-compose.web.yml up
```

All three services should start. What happens if you run just `docker compose up`? Which file does it use?

Now create a third file: **docker-compose.override.yml**

In this file, override some settings:
* Change the simple-go-web port from 8080 to 9090
* Change the `APP_VERSION` build argument to "4.0.0"
* Add a new environment variable to the api service

Run `docker compose up` (without specifying files). What happens?

Compose automatically merges docker-compose.yml and docker-compose.override.yml. Verify:
* The web app should be on port 9090 (not 8080)
* The version should be "4.0.0"

Now run with explicit files:
```bash
docker compose -f docker-compose.yml -f docker-compose.web.yml up
```

What port is the web app on now? The override file wasn't used.

**Think:** What's the precedence order when merging compose files? When would you use docker-compose.override.yml vs explicitly specifying files? How would you use this for dev vs prod environments?

### Task 5: Lifecycle Management - Compose vs Docker

Understand the difference between managing containers with `docker` commands vs `docker compose` commands.

Start your full stack with `docker compose up -d`.

List containers using both methods:
```bash
docker ps
docker compose ps
```

What's different about the output?

Now stop the api container using the docker command:
```bash
docker stop <api-container-name>
```

Check the status:
```bash
docker compose ps
```

What does Compose show? Start the container again using docker:
```bash
docker start <api-container-name>
```

Now stop the api using Compose:
```bash
docker compose stop api
```

What's different?

Try to remove a container using docker:
```bash
docker rm <api-container-name>
```

Does it work? Why or why not?

Now experiment with Compose commands:
```bash
docker compose stop        # Stop all services
docker compose start       # Start stopped services
docker compose restart     # Restart services
docker compose down        # Stop and remove containers
docker compose down -v     # Stop, remove containers AND volumes
```

Check logs:
```bash
docker compose logs api              # Logs from api service
docker compose logs -f               # Follow logs from all services
docker compose logs --tail=20 api    # Last 20 lines from api
```

After running `docker compose down`, try to start containers with `docker start`. What happens?

**Think:** What's the difference between `docker compose stop` and `docker compose down`? When would you use each? What does Compose track that plain Docker commands don't? What happens to networks and volumes with each command?

## Resources

* [Docker Compose overview](https://docs.docker.com/compose/)
* [Compose file reference](https://docs.docker.com/compose/compose-file/)
* [Compose build reference](https://docs.docker.com/compose/compose-file/build/)
* [Compose CLI reference](https://docs.docker.com/compose/reference/)
* [Networking in Compose](https://docs.docker.com/compose/networking/)
* [Volumes in Compose](https://docs.docker.com/compose/compose-file/07-volumes/)

## Verification Checklist

You've completed this exercise when:
* [ ] Can define multi-container applications with proper network isolation
* [ ] Can persist data across compose down/up cycles
* [ ] Can build images with Compose using build context and arguments
* [ ] Understand when Compose rebuilds images
* [ ] Can scale services and understand the limitations
* [ ] Understand replica naming patterns
* [ ] Can use multiple compose files together
* [ ] Understand docker-compose.override.yml behavior
* [ ] Know the precedence order when merging compose files
* [ ] Understand the difference between docker and docker compose commands
* [ ] Can manage container lifecycle with Compose
* [ ] Know when to use stop vs down
* [ ] Can view and follow logs with Compose

## Cleanup

Stop and remove all resources:
```bash
docker compose down -v
```

Remove any additional compose files you created during experimentation.

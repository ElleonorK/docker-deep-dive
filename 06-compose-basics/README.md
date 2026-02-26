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

In exercises 02-04, you ran multiple containers, had to configure two isolated networks, expose ports to localhost, and keep data persistent.

All of this you did manually and running many commands. That was tedious and error-prone. Now let's do it all at once and "as code".

Build a complete stack with three applications using your already-built images:
* **simple-go-web** - the Go web application
* **api** - the Node.js API that connects to a database
* **database** - postgres 16-alpine from Docker Hub

#### *Requirements*

**Port Exposure:**
* simple-go-web accessible at http://localhost:8080
* api accessible at http://localhost:3000
* database NOT accessible from host

**Network Isolation:**
* simple-go-web should NOT be able to reach the database
* api should be able to reach the database

**Data Persistence:**
* Database data must survive `docker compose down` and `docker compose up`

**Success Criteria:**
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

Modify the compose file so simple-go-web and api images are built at runtime instead of using existing ones.

Delete your old images and run `docker compose up`.

**Success Criteria:**

Once it's working, visit http://localhost:8080. The Go app should display a version. But which version does it show?

**Here's the next challenge:**
Make the version configurable through your compose file. Set `APP_VERSION` to "2.0.0" and make sure that's what you see.
Do it in two different ways:
1. Set it during build
2. Change it at runtime
Set `APP_VERSION` to "2.0.0".

**Hint:** You learned about build arguments vs environment variables in exercise 02.

### Task 3: Context is Key

Your compose file is in `06-compose-basics/` but your apps are in `apps/`. You've been running compose from the exercise directory, but what if you need to run it from somewhere else?

Try running compose from different locations:

```bash
# From the exercise directory (where you've been running it)
cd 06-compose-basics
docker compose up --build

# From the repository root
cd ..
docker compose -f 06-compose-basics/docker-compose.yml up --build

# From inside the apps directory
cd apps
docker compose -f ../06-compose-basics/docker-compose.yml up --build
```

What happens in each case? Does the build work from all locations? If some fail, what error do you see?

**Your challenge:** Fix your compose file so it works regardless of where you run the command from.

**Hint:** The build context path in your compose file is relative to something. Is it relative to the compose file location, or to where you run the command?

**Think:** In a real project, developers might run compose from different directories, or CI/CD might run it from the repo root. How do you make your setup portable?

Now that you understand build context paths, let's look at what's actually being sent. Your team keeps complaining that builds are slow. You notice the "Sending build context to Docker daemon" message takes forever.

Create some realistic development clutter in your `apps/api/` directory:
```bash
cd apps/api
npm install  # Creates node_modules
mkdir -p logs coverage .git
echo "debug log" > logs/app.log
echo "test coverage" > coverage/report.html
dd if=/dev/zero of=.git/large-blob bs=1M count=50
```

Now rebuild your api service from the repository root and watch carefully:
```bash
docker compose -f 06-compose-basics/docker-compose.yml build api
```

Look at the "Sending build context" line. How much data is being transferred? Time how long it takes.

Exec into the running api container and list what files are actually there. Does the container need `node_modules` from your host? Does it need logs? Coverage reports? Git history?

**Your challenge:** Create a `.dockerignore` file in `apps/api/` that excludes unnecessary files. Rebuild and measure the difference.

**Success criteria:** Build context should drop from megabytes to kilobytes. Build time should be noticeably faster.

**Think:** What files does a Node.js app need at runtime vs what accumulates during development? Why send files to the Docker daemon that won't even be used in the image?

### Task 4: Scaling and Replicas

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

### Task 5: Multiple Compose Files and Overrides

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

### Task 6: Lifecycle Management - Compose vs Docker

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

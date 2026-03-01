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
(Assuming your docker-compose.yml is inside this exercise directory)
```bash
# From the exercise directory
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

**Hint:** Look at the title of the task.

### Task 4: Reducing Clutter

**Mission:**
Let's look at a real scenario you may face as a DevOps Engineer.

The development team keeps complaining that starting their compose environment takes forever. You notice it is stuck on "Sending build context to Docker daemon" message for a long time.

Start the investigation 🔍 👀

**Reproduce the issue:**

Create some realistic clutter in your `apps/api/` directory that a developer might have:

```bash
cd apps/api
mkdir -p logs media
# Generate a large log file (more than 60MB)
for i in {1..1000000}; do echo "[2024-02-26 10:23:45] INFO: Processing request $i from user-$((i % 100))" >> logs/app.log; done

# Simulate some media files developers might have
# These are not real media files, don't try to play them :)
dd if=/dev/urandom of=media/demo.gif bs=100M count=5
dd if=/dev/urandom of=media/tutorial.mp4 bs=1M count=5000
```

Rebuild the api app using compose.
How much time does "Sending build context" take?

Exec into the running api container and list what files are actually there. Does the container contain logs? Screenshots? Videos?
Is it using any of them?

**Think:** You could change your Dockerfile to only copy certain files. Does it solve your time problem, though? Why?

**Your challenge:** Add something to `apps/api` folder that will stop docker from loading all these large files during build.

**Success criteria:** Build time should be noticeably faster.

**Bonus:** What if developers add a new garbage folder? Change what you added to only load the specific files you need (allow vs deny logic).

### Task 5: When One Isn't Enough

Your api service is getting hammered by traffic and the CPU utilization is through the roof!
You decide to double the capacity for handling requests.

Without reloading your docker compose, increase the number of `api` containers to 2.
Are you running into any issues?

**Your challenge:** Fix your compose file so you can run 2 `api` containers at the same time without collisions.

**Success criteria:** Both containers start successfully and you can reach them in browser. Check `docker compose ps` - you should see two api containers running.

What are the containers named? Notice the pattern?

Now that you `api` is stable and all requests are being fulfilled, the database starts to choke. In a real scenario, you will start seeing 502 errors.

So now you should increase the number of database containers to 3.
> [!Tip]
> Quite a few databases (and not just) require an uneven number of instances due to "leader election" logic.
> 
> Even though Postgres is not one of them, it's a good point to remember.

When all three start, create some test data:
```bash
docker compose exec database psql -U postgres -c "CREATE TABLE scaling_test (id INT, data TEXT);"
docker compose exec database psql -U postgres -c "INSERT INTO scaling_test VALUES (1, 'test data');"
```

Now list all your database containers and pick one. Delete it:
```bash
docker rm -f <one-database-container-name>
```

What happens? Check `docker compose ps`. What do you see?

Wait a moment and check again. Did Compose recreate the container? What's its name?

Query your test data from the recreated container. Is it still there?

Now exec into one of the database containers and create a file in the container's filesystem (not in a shared volume):
```bash
docker compose exec <one-database-container-name> touch /tmp/my-local-file.txt
docker compose exec <one-database-container-name> ls /tmp/
```

Delete that specific container and let Compose recreate it. Check if the file is still there in the new container.

**Think:** Does a container name match with it's data? What happens inside docker daemon when a container in a set vanishes?

### Task 6: Configuration

Your team wants to run different web app versions on different environments. Time to split things up.

First, let's split our services into 2 independent stacks.
Move simple-go-web to a separate file called `docker-compose-web.yml`. Keep api and database in `docker-compose.yml`.

**Your challenge:** Run all three services with a single command.

**Think:** What are the pros and cons of splitting into multiple files/stacks?

Now, let's allow our developers to change things per environment.
Create `docker-compose-web.dev.yml` that sets `APP_VERSION` to "3.0.0-dev" for simple-go-web.

Create `docker-compose-web.prod.yml` that sets `APP_VERSION` to "3.0.0" for simple-go-web.

> [!Warning]
> I hope you didn't fully copy-paste `docker-compose-web.yml` 😈

Run the web app with dev configuration. What version shows up? Try with prod configuration. What do you see?

**For the next challenge, switch to api and database stack.** Your production environment requires a secure database password. Make sure that in dev environment the api uses a default password (check `apps/api/README.md` again), while in prod you set it to "superSecureUnbreakablePass@123".

**Success criteria:** Run your full stack (api, database, and web) with production configuration. Visit http://localhost:3000/db-check to verify the api connects to the database with the production password.

**Another Real Scenario**

A new API developer on your team complains that when running the stack in production configuration, they are unable to connect the database with the production password.

You notice that:
* They are not actually using the api container, but running the source code manually. Compose is used only for the database. (It's a common way for developers to test changes before committing.)
* They are setting the necessary environment variables, including the correct password. You don't notice any values that are incorrect.
* They pass the files `docker-compose.prod.yml` first and `docker-compose.yml` second when running the compose command.

**Your Challenge:** You are not here to debate security with the developer or teach them. They need to test it NOW. Find the reasons for the problem and provide solutions.

**Success Criteria:** The developer can run a single docker compose command, use the production password and connect to the database created by the compose.

**Think:** Now that the release has come out and everyone is calm - what is the actual better way to address the issue?
This time there isn't a single correct answer, however I will urge you to research development best practices from a DevOps perspective.

### Task 7: To Compose OR To Not Compose?

You've been using `docker compose` commands, but you can also use plain `docker` commands on the same containers. What's the relationship between the two tools?

After an arduous journey, let's just play around. No challenge here.

Start your stack:
```bash
docker compose up -d
```

List containers both ways:
```bash
docker ps
docker compose ps
```

What's different about the output?

Now stop the api container using plain Docker:
```bash
docker stop <api-container-name>
```

Check what Compose thinks:
```bash
docker compose ps
```

What status does it show? Start it again with Docker:
```bash
docker start <api-container-name>
```

Now stop it using Compose:
```bash
docker compose stop api
```

Try to remove the stopped container:
```bash
docker rm <api-container-name>
```

Does it work? Why or why not?

Restart everything and try these Compose commands. After each one, check `docker compose ps` and see what changed:
```bash
docker compose stop
docker compose start
docker compose restart api
docker compose pause database
docker compose unpause database
```

Now the big one:
```bash
docker compose down
```

Run `docker ps -a`. What's left? Try to start the containers with `docker start`. What happens?

Check your logs before they're gone:
```bash
docker compose up -d
docker compose logs api
docker compose logs -f --tail=50
```

**Think:** What's the difference between `stop` and `down`? What does Compose manage that plain Docker can't? When would you use `docker compose down -v` vs just `docker compose down`?

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

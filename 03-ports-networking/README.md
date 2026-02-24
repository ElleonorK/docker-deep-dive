# Exercise 03: Ports and Networking

Learn how containers communicate with the host and with each other.

## Objectives

By the end of this exercise, you'll be comfortable with:
* Port mapping and exposure
* Container-to-container communication
* Custom networks
* Network isolation
* DNS resolution between containers

## Prerequisites

You should have completed exercise 02 and have a `simple-web` image built. If not, go back and build it first.

## Tasks

### Task 1: Port Mapping Basics

Let's start with the simple-web app you built in exercise 02.

**Part A:** Run your `simple-web` container and map its port so you can access it at `http://localhost:8080`

**Part B:** Now run a second instance of the same image, but map it to `http://localhost:3000` instead.

Both containers should run simultaneously without conflicts.

**Success criteria:** Both URLs work in your browser and show the app running simultaneously.

**Think:** How does Docker let multiple containers use the same internal port without conflicts?

### Task 2: Container-to-Container Communication

Now let's work with a more realistic scenario: an API that needs to talk to a database.

**Setup:**

First, get the postgres database running:
```bash
docker run -d \
  --name database \
  -e POSTGRES_PASSWORD=mysecretpassword \
  postgres:16-alpine
```

Next, build the API app. Navigate to `apps/api/` and check its README to understand what it does. Then create a Dockerfile for it and build an image called `api`.

**Your challenge:** Run the API container so it can connect to postgres using the hostname `database` (not an IP address).

**Success criteria:** The API starts without connection errors when you check the logs.

**Hint:** Do all containers know each other by container name?

### Task 3: Network Isolation

With your setup from Task 2 running:
- The API should be able to reach the database
- Your host machine should be able to reach the API (map a port to test this)
- Your host machine should NOT be able to reach the database directly

**Test it:** Try to connect to postgres from your host on port `5432`. It should fail (connection refused or timeout).

But when you access the API endpoint, it should successfully query the database.

**Think:** Why is this isolation useful in production? What security benefits does it provide?

### Task 4: Inspect the Network

Let's understand what Docker created for you.

Find out:
- What network are your API and database containers on?
- What are their IP addresses?
- Can you ping one container from another?

**Hints:** 
* Networks can be inspected with `docker network inspect <network-name>`
* You can execute commands inside containers with `docker exec`
* Try `docker exec <container-name> ping <other-container-name>`

### Task 5: Multiple Networks

Create a scenario with three containers:
* frontend (use your `simple-web` image)
* backend (your `api` image)
* database (postgres)

**Requirements:**
* frontend can reach backend
* backend can reach database
* frontend CANNOT reach database

This requires creating two separate networks and connecting containers to the right ones.

**Success criteria:** 
* You can exec into the frontend container and successfully curl the backend
* You can exec into the frontend container and fail to reach the database
* The backend can still query the database successfully

**Think:** How does this architecture pattern (frontend → backend → database) improve security?

## Resources

* [Docker networking overview](https://docs.docker.com/network/)
* [Container networking](https://docs.docker.com/config/containers/container-networking/)
* [Bridge networks](https://docs.docker.com/network/bridge/)

## Verification Checklist

You've completed this exercise when:
* [ ] Can access container services from host on mapped ports
* [ ] Can run multiple containers on different host ports
* [ ] Containers can reach each other by name (DNS)
* [ ] Understand how to isolate containers from host access
* [ ] Can inspect network configuration and container IPs
* [ ] Can create multi-network isolation scenarios

## Cleanup

Stop and remove all containers you created:
```bash
docker stop <container-name>
docker rm <container-name>
```

Remove any custom networks you created:
```bash
docker network rm <network-name>
```

# Exercise 01: Docker Basics

Learn fundamental Docker commands by working with existing images from Docker Hub.

## Objectives

By the end of this exercise, you'll be comfortable with:
- Pulling images from Docker Hub
- Running containers
- Inspecting running and stopped containers
- Viewing logs
- Executing commands inside containers
- Managing container lifecycle

## Tasks

### Task 1: Run Your First Container

Run an nginx web server container. Access it from your browser at http://localhost:8080

### Task 2: Explore a Running Container

While nginx is running, find out:
- What version of nginx is installed?
- What Linux distribution is the container using?
- What processes are running inside?

Hint: You can execute commands inside a running container.

### Task 3: Container Lifecycle

Start a postgres database container. Then:
- Stop it
- Start it again (same container, not a new one)
- Check its logs
- Remove it completely

### Task 4: Detached vs Interactive

Run a redis container in detached mode. Then run another redis container interactively and connect to it using redis-cli.

What's the difference between these two modes?

### Task 5: Inspect Everything

Pick any running container and find out:
- Its IP address
- What ports it exposes
- What volumes it has mounted
- What environment variables are set

Hint: There's a command that shows you all container details in JSON format.

### Task 6: Clean Up

List all containers (running and stopped). Remove all stopped containers.

List all images. Remove any images you no longer need.

## Resources

- [Docker run reference](https://docs.docker.com/engine/reference/run/)
- [Docker CLI documentation](https://docs.docker.com/engine/reference/commandline/cli/)
- [Docker Hub](https://hub.docker.com/) - search for nginx, postgres, redis

## Verification

You've completed this exercise when you can:
- ✓ Run containers in both detached and interactive modes
- ✓ Access a web server running in a container from your browser
- ✓ Execute commands inside running containers
- ✓ View and understand container logs
- ✓ Inspect container configuration
- ✓ Manage container lifecycle (start, stop, remove)
- ✓ Clean up unused containers and images

No solution file needed - these are standard Docker commands you'll use daily!

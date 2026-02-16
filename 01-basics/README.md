## Exercise 01: Docker Basics

Learn fundamental Docker commands by working with existing images from Docker Hub.

## Objectives

By the end of this exercise, you'll be comfortable with:

*   Pulling images from Docker Hub
*   Running containers in different modes
*   Viewing container logs
*   Inspecting container and image metadata
*   Understanding container states and transitions

## Tasks

### Task 1: Run Your First Container

Run an nginx web server container. Access it from your browser at `http://localhost:8080` (make it work!).

**Hint**: Nginx web server listens to port 80.

### Task 2: Run a Container with a Command

Run a container using the `devmisterio/fortune` image to get your fortune told. This container runs a command and exits immediately.

Then try the `rancher/cowsay` image and make it say something fun.

What's the difference in how you run these?

**Hint**: Try changing your fate.

### Task 3: Understanding Image Tags

Run `grafana/grafana` image exposing port 3000. Open it in browser and look at the bottom of the page. Which version did you get?

Now run `grafana/grafana:main-ubuntu` and then `grafana/grafana:11.6`. What's different about these?

Try running `ubuntu:22.04` and `ubuntu:20.04` side by side. How can you tell which is which when they're running?

**Challenge**: Find an image on Docker Hub that has a `latest` tag and at least 3 other version tags. What do the tags tell you?

### Task 4: Naming Your Containers

Run three nginx containers without naming them. List them. What are they called?

Now run three more nginx containers but give each a meaningful name. List them again.

Try to start a new container with a name that's already in use. What happens?

Remove a stopped container by its name. Remove another by its ID. Which is easier to remember?

### Task 5: Detached vs Interactive Modes

Try running the `danielkraic/asciiquarium` image. What happens? Can you see the aquarium?

Now try running it in a different way. Can you see the fish swimming now?

Next, run an nginx container so it keeps running in the background while you do other things. If you see logs - you are doing it wrong!

What's the difference between these two ways of running containers? When would you use each?

### Task 6: Understanding Container States

Run a `postgres` database container in the background with a memorable name. Transition it through these states:

*   Start it
*   Check status (running)
*   Pause it
*   Check status (paused)
*   Unpause it
*   Stop it
*   Check status (exited)
*   Start it again (same container)
*   Restart it (one command)
*   Delete it

What's the difference between stop/start and pause/unpause?
Did you run into problems while trying to delete the container? Why?

### Task 7: Working with Logs

Run the `chentex/random-logger` container in the background. Then explore its logs:

*   View all logs
*   Follow the logs in real-time
*   Stop following logs (Ctrl+C)
*   View only the last 10 lines of logs
*   View logs with timestamps
*   Wait for a few minutes. View logs from 2 minutes after you started the container until 1 minute ago.

Logs are crucial for debugging containerized applications!

### Task 8: Inspecting Containers and Images

Pick any running container and find out:

*   Its IP address
*   What ports it exposes
*   What environment variables are set
*   When it was created
*   Its current state and status

Now inspect the image that container is based on:

*   Who maintains this image?
*   What base image is it built from? (look for parent image)
*   What is the default CMD or ENTRYPOINT?
*   When was the image created?

**Hint**: `docker inspect` returns JSON. You can format output with `--format` or pipe to `jq` for easier reading.

### Task 9: Clean Up

List all containers (running and stopped). Remove all stopped containers.

List all images. Remove any images you no longer need.

Bonus: Try the `docker system prune` command to clean up everything at once (be careful!).

## Resources

*   [Docker run reference](https://docs.docker.com/engine/reference/run/)
*   [Docker logs documentation](https://docs.docker.com/engine/reference/commandline/logs/)
*   [Docker inspect documentation](https://docs.docker.com/engine/reference/commandline/inspect/)
*   [Docker Hub](https://hub.docker.com/) - search for images to experiment with
*   [Container lifecycle states](https://docs.docker.com/engine/reference/commandline/ps/)

## Verification

You've completed this exercise when you can:

- [ ] Run containers in both detached and interactive modes
- [ ] Run containers with custom commands
- [ ] Access a web server running in a container from your browser
- [ ] Transition containers through different states (run, pause, stop, start, restart)
- [ ] View, follow, and filter container logs
- [ ] Inspect both container and image metadata
- [ ] Clean up unused containers and images

No solution file needed - these are standard Docker commands you'll use daily!
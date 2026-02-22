# Exercise 02: Building Docker Images

Learn to create production-ready Docker images through hands-on practice.

## Objectives

By the end of this exercise, you'll be comfortable with:

*   Writing Dockerfiles from scratch
*   Making small images
*   Implementing security best practices
*   Configuring environment variables inside the image
*   Understanding startup commands

## The Mission

You're containerizing a Go web app. Your goal: build an image that's small, secure, and production-ready.

The app lives in `../apps/simple-go-web/` - check its README to see what it does and how to use it.

## Tasks

### Task 1: Build Your First Image

Create a `Dockerfile` in this directory that runs the simple-go-web app.

The app lives in `../apps/simple-go-web/` - explore the README to understand what it needs to run.

Your goal: build an image, run a container from it, and access the app at http://localhost:8080

**Success criteria:** You see JSON with app info when you visit the URL.

### Task 2: Show Different Versions

The app displays a version from the `APP_VERSION` environment variable. 

Set `APP_VERSION` to "1.0.0" in your Dockerfile, build an image tagged as `simple-web:v1`, and run it. Verify the version shows up at http://localhost:8080

Now change `APP_VERSION` to "2.0.0" in your Dockerfile, build another image tagged as `simple-web:v2`, and run it on a different port. Verify both containers are running with different versions.

Notice how image tags can be anything - not just numbers. Try building with tags like `simple-web:latest`, `simple-web:production`, or `simple-web:feature-xyz`.

What's annoying about modifying the Dockerfile for each version? There's a better way coming in Task 3.

### Task 3: Make Version Configurable

Remember Task 2 where you had to edit the Dockerfile for each version? That's not sustainable.

**Your goal:** build two images with versions "1.0.0" and "2.0.0" without modifying the Dockerfile between builds.
You will need to change something in the Dockerfile to make it work. 

**Success criteria:** Run each image and verify they show different versions.

**Hint:** The app needs the version as an environment variable at runtime. But you want to set it at build time. What do you do?

### Task 4: Make It User-Friendly

Try to run your container using Docker Desktop:
1. Open Docker Desktop
2. Go to Images
3. Find your image and click Run
4. Expand "Optional settings"
5. Look at the Ports section

What do you see? Can you map a port?

Now add something to your Dockerfile that makes the port to show up. Rebuild and try again - the Ports section should look different now.

**Think:** Your container runs fine without this. So why bother? Who benefits from knowing which ports an image uses? Are there other parts of the Dockerfile that similarly don't affect the running container?

### Task 5: Command Flexibility

Your Go app accepts a `test` argument that runs a quick health check instead of starting the server.

**Your challenge:** Set up your Dockerfile so that:
*   `docker run simple-web` runs the web server (default)
*   `docker run simple-web test` runs the test command

**Bonus challenge:** Can you override the entire command to run something else, like opening a shell?

**Think:** Why would you want to have a default command for an image? What happens when you don't have it? When would you want to override the base command entirely?

### Task 6: Separate Build from Runtime

**Preparation:** Start your container in detached mode so you can inspect it while it's running. Give it a memorable name to make the commands easier.

Now let's inspect what's inside your image.

**Check what's in your image:** Run `docker exec <container-name> which go`

Did you find Go inside? Why is that a security risk?

**Your challenge:** Restructure your Dockerfile to use two stages:
1. A build stage that compiles your Go app
2. A runtime stage that only contains the compiled binary and message file

After rebuilding, verify that `which go` returns an error (Go toolchain not found).

**Think:** What does a compiled Go binary need to run? Compare this to interpreted languages like Python.

### Task 7: Secure the Runtime

**Preparation:** Start your container in detached mode so you can inspect it while it's running. Give it a memorable name to make the commands easier.

**Experiment with your running container:**

Try to modify the application from inside the container:
```bash
docker exec -it <container-name> sh
```

Once inside, try to change the message that the app displays:
```bash
echo "HACKED!" > message.txt
```

Exit the shell and visit http://localhost:8080 in your browser (no need to restart).

What do you see? Were you able to modify the running application? In a production environment, if an attacker exploits your app, they shouldn't be able to change the functionality!

**Your challenge:** Secure your image so that an attacker who compromises the app can't modify the application files.

After rebuilding, try the same attack again. Can you still modify message.txt?

**Hints:**
* Who should run the application?
* What permissions should the files have?
* How does multi-stage build help you?

### Task 8: Minimize Image Size

**Check your image size:** Run `docker images simple-web`

**Your goal:** Get the image under 20MB.

You've already removed build tools with multi-stage builds. What else can you optimize?

**Think:**
*   Does the build stage size matter, or only the final stage?
*   What's the smallest possible base image you could use for a Go binary?
*   Go binaries are statically compiled - what does that mean for dependencies?

## Verification Checklist

You've completed this exercise when:
- [ ]   Image builds successfully
- [ ]   Version is configurable at build time (without modifying app code)
- [ ]   You can set the localhost port in Docker Desktop
- [ ]   You cannot modify message.txt in the running app
- [ ]   Image size is under 20MB

## Cleanup

Remove all containers and older versions of the images you created.

## Resources

*   [Dockerfile reference](https://docs.docker.com/engine/reference/builder/)
*   [Best practices for writing Dockerfiles](https://docs.docker.com/develop/develop-images/dockerfile_best-practices/)
*   [Multi-stage builds](https://docs.docker.com/build/building/multi-stage/)
*   [Go Docker images](https://hub.docker.com/_/golang)


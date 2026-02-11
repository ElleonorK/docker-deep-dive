# Exercise 02: Building Docker Images

Learn to create, optimize, and secure Docker images using Dockerfiles.

## Objectives

Master the art of building production-ready Docker images:
- Write Dockerfiles
- Optimize image size
- Implement security best practices
- Use build arguments
- Understand multi-stage builds
- Configure image metadata

## The Application

You'll be working with the Node.js app in `../apps/simple-web/`. 

Read its README to understand what it does. You don't need to modify the application code.

## Tasks

### Task 1: Build Your First Image

Create a Dockerfile in this directory that builds and runs the simple-web app.

Requirements:
- The image should build successfully
- Running a container from this image should start the web server
- You should be able to access http://localhost:8080 from your browser

Hint: You'll need to install dependencies and start the app.

### Task 2: Reduce Image Size

Check your image size with `docker images`. It's probably over 500MB.

Industry standard for a Node.js app like this is under 100MB. Reduce your image size to meet this standard.

Hint: Not all base images are created equal.

### Task 3: Remove Root Access

Run this command with your container: `docker exec <container-name> whoami`

If it says "root", that's a security risk. Fix your Dockerfile so the app runs as a non-root user.

Verify: `whoami` should return something other than "root".

### Task 4: Make Version Configurable

The app displays a version from the APP_VERSION environment variable. Currently, you'd have to hardcode this in the Dockerfile.

Make it so you can specify the version at build time:
```bash
docker build --build-arg VERSION=1.2.3 -t simple-web:1.2.3 .
docker build --build-arg VERSION=2.0.0 -t simple-web:2.0.0 .
```

Run each image and verify they show different versions at http://localhost:8080

### Task 5: Expose the Port

Try to run your container using Docker Desktop UI:
1. Open Docker Desktop
2. Go to Images
3. Find your image and click Run
4. Expand "Optional settings"
5. Try to add port 8080 in the Ports section

What happens? Fix your Dockerfile so Docker Desktop knows which port to expose.

### Task 6: Remove Build Tools

Your final image shouldn't contain npm or any build tools - only what's needed to run the app.

Verify by running: `docker exec <container-name> which npm`

It should return nothing (npm not found). But the app should still run!

Hint: You need to build in one stage and run in another.

### Task 7: Support Command Line Arguments

Your container should work in two ways:
- `docker run simple-web` - starts the web server
- `docker run simple-web --help` - shows help text (you can make this show anything)

This requires understanding the difference between ENTRYPOINT and CMD.

### Task 8: Fix File Permissions

Run: `docker exec <container-name> ls -la /app`

If files are owned by root, fix your Dockerfile so they're owned by the non-root user.

Hint: COPY has options for setting ownership.

## Resources

- [Dockerfile reference](https://docs.docker.com/engine/reference/builder/)
- [Best practices for writing Dockerfiles](https://docs.docker.com/develop/develop-images/dockerfile_best-practices/)
- [Multi-stage builds](https://docs.docker.com/build/building/multi-stage/)
- [Node.js Docker images](https://hub.docker.com/_/node)

## Verification Checklist

You've completed this exercise when:
- ✓ Image builds successfully
- ✓ Image size is under 100MB
- ✓ App runs as non-root user
- ✓ Version is configurable at build time
- ✓ Port is exposed (visible in Docker Desktop)
- ✓ No build tools in final image
- ✓ Container responds to both normal run and --help
- ✓ Files are owned by non-root user

This is the most challenging exercise - take your time and read the documentation!

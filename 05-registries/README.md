# Exercise 05: Container Registries

Learn how container registries work and understand image naming conventions across different registry types.

## Objectives

Master working with container registries:
- Run a local registry
- Push and pull images to/from local registry
- Understand image naming conventions
- Work with Docker Hub
- Work with third-party registries (ghcr.io, quay.io)
- Compare registry behaviors

## Tasks

### Task 1: Start a Local Registry

Run a container registry on your local machine at `localhost:5000`.

Hint: There's an official registry image on Docker Hub.

Verify: `curl http://localhost:5000/v2/_catalog` should return a JSON response (probably empty for now).

### Task 2: Push to Local Registry

Take the simple-web image you built in exercise 02. Tag it appropriately for your local registry and push it.

Verify: `curl http://localhost:5000/v2/_catalog` should now list your image.

### Task 3: Pull from Local Registry

Delete your local simple-web image completely. Then pull it back from your local registry.

Run `docker images` and look at the image name. What's different about it compared to images from Docker Hub?

### Task 4: Compare Image Names

Pull the same image from three different sources:
1. Your local registry (localhost:5000)
2. Docker Hub (docker.io or no prefix)
3. GitHub Container Registry (ghcr.io)

For Docker Hub and ghcr.io, you can use any public image (like nginx, or search for simple examples).

Run `docker images` and compare the REPOSITORY column. What pattern do you notice?

### Task 5: Understand the Naming Convention

Based on Task 4, answer these questions:
- Why do some images have hostnames in their names and others don't?
- What's the "default" registry?
- How does Docker know where to pull from?

Write your answers in a file called `registry-notes.txt`.

### Task 6: Registry Authentication

Try to push an image to Docker Hub (you'll need a Docker Hub account).

What's different about pushing to Docker Hub vs your local registry?

### Task 7: Pull from Multiple Registries

Pull images from at least three different registries:
- docker.io (Docker Hub)
- ghcr.io (GitHub Container Registry)
- quay.io (Red Hat Quay)

Compare:
- Image naming format
- Pull speed
- Any authentication requirements

### Task 8: Registry Inspection

Your local registry stores images somewhere. Find out:
- Where is the registry storing image data?
- How much space is it using?
- Can you see the image layers?

Hint: The registry is just a container with a volume.

## Resources

- [Docker Registry](https://docs.docker.com/registry/)
- [Deploy a registry server](https://docs.docker.com/registry/deploying/)
- [Image naming conventions](https://docs.docker.com/engine/reference/commandline/tag/)
- [Docker Hub](https://hub.docker.com/)
- [GitHub Container Registry](https://docs.github.com/en/packages/working-with-a-github-packages-registry/working-with-the-container-registry)

## Verification Checklist

You've completed this exercise when:
- ✓ Local registry running and accessible
- ✓ Can push images to local registry
- ✓ Can pull images from local registry
- ✓ Understand image naming with registry hostnames
- ✓ Can identify the default registry
- ✓ Can authenticate with Docker Hub
- ✓ Have pulled from at least 3 different registries
- ✓ Understand where registry data is stored

This exercise is about understanding the registry ecosystem, not just running commands!

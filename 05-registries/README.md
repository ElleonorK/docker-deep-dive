# Exercise 05: Container Registries

Learn where images come from, how they're named, and how to work with different registries.

## Objectives

By the end of this exercise, you'll be comfortable with:
* Working with multiple container registries
* Understanding image naming and identification
* Managing registry authentication
* Operating your own registry

## Tasks

### Task 1: The Mystery of Names

Pull the `alpine` image:
```bash
docker pull alpine
```

Now pull it again with its full name:
```bash
docker pull docker.io/library/alpine
```

Run `docker images` and look at both. How many alpine images do you have?

**Think:** What's the relationship between `alpine` and `docker.io/library/alpine`? What is this `docker.io` domain?

### Task 2: Layered Cakes

Pull the `alpine` image and watch what layers get downloaded:
```bash
docker pull alpine
```

Now pull an image based on alpine:
```bash
docker pull nginx:alpine
```

**Think:** Notice how quickly some layers finish downloading? What do you think happens here?

Next, pull several more alpine-based images:
```bash
docker pull redis:alpine
docker pull postgres:alpine
docker pull node:alpine
```

Now check the size of each image. Use Docker Desktop UI or run the following command in terminal:
```bash
docker images --format "table {{.Repository}}:{{.Tag}}\t{{.ID}}\t{{.Size}}"
```

Add up the sizes you get.

Now check your actual disk usage:
```bash
docker system df
```

Look at the Images row - is the number equal to the one you calculated?

Check the details per image:
```bash
docker system df -v
```

Look at the SHARED SIZE column for each alpine-based image. Can you now understand why the disk usage looks like that?

**Think:** If layers are shared, what happens when you delete one image that shares layers with another?

### Task 3: Alternative Registries

You've been pulling images from Docker Hub this whole time. But there are other registries out there.

**Your challenge:** Pull the Prometheus image, but get it from Red Hat's registry (Quay) instead of Docker Hub.

Once you have it, run `docker images` and look at the REPOSITORY column. Compare it to your other images.

Now pull aws-for-fluent-bit (AWS distribution of Fluent Bit) image from Amazon's ECR registry:
```bash
docker pull public.ecr.aws/aws-observability/aws-for-fluent-bit:latest
```

Watch how long it takes. Now pull the same tag from Docker Hub:
```bash
docker pull amazon/aws-for-fluent-bit:latest
```

How long did the second pull take? Run `docker images` and look at both entries.

**Think:** Why was the second pull so fast? What does this tell you about how Docker identifies images?

### Task 4: Authentication and Pushing

Create a free Docker Hub account at https://hub.docker.com if you don't have one.

Login:
```bash
docker login
```

Where did Docker store your credentials? Check `~/.docker/config.json`.

**Your challenge:** Take one of the images you built in previous exercises (like `simple-web` or `api`) and push it to your Docker Hub account.

**Hint:** When you make an account, Docker Hub allocates a registry with your username.

Watch what happens during the push. Are all layers being uploaded?

Now remove your local copy and pull it back from Docker Hub. What layers were downloaded?

**Think:** How does Docker Hub (or any registry) store image layers?

### Task 5: Multiple Registry Authentication

**Your challenge:** Set up authentication for both Docker Hub and GitHub Container Registry on the same machine.

For GitHub, you'll need a Personal Access Token with `read:packages` permission from https://github.com/settings/tokens

**Success criteria:** You should be able to push or pull an image to your Docker Hub account AND ghcr.io based on tag alone, without re-authenticating.

After setting this up, inspect `~/.docker/config.json`. How does Docker manage credentials for multiple registries?

**Think:** Why don't everyone just use Docker Hub? Why do companies push THE SAME IMAGES to multiple public registries (like the one from Task 3)?

### Task 6: Local Registry

**Your challenge:** Run your own container registry in a container and use it to store one of your images from previous exercises.

**Success criteria:**  You can see your image information in http://localhost:5000/v2/_catalog

**Think:** What are the Pros and Cons of maintaining your own registry vs using a SaaS offering?

### Task 7: Beyond Tags

Tags like `latest` or `v1.0` are convenient for humans, but there's another way Docker identifies images.

Pull nginx:
```bash
docker pull nginx:alpine
```

Now inspect it:
```bash
docker inspect nginx:alpine | grep -i digest
```

You'll see something like `sha256:abc123...`. 

**Think:** This digest is long and cryptic. What decides which digest refers to the image? Are tags (names) just friendly aliases for these digests? Or is there an actual difference in their meaning?

## Resources

* [Docker Registry Documentation](https://docs.docker.com/registry/)
* [Image Naming Conventions](https://docs.docker.com/engine/reference/commandline/tag/)
* [Content Addressable Storage](https://docs.docker.com/registry/spec/api/#content-digests)
* [Docker Hub](https://hub.docker.com/)
* [GitHub Container Registry](https://docs.github.com/en/packages/working-with-a-github-packages-registry/working-with-the-container-registry)
* [Understanding Image Layers](https://docs.docker.com/storage/storagedriver/)

## Verification Checklist

You've completed this exercise when:
* [ ] Can pull images from multiple registries
* [ ] Can push your own images to a registry
* [ ] Can authenticate with multiple registries
* [ ] Can run your own local registry
* [ ] Understand how Docker identifies images beyond just tags

## Cleanup

Stop and remove your local registry:
```bash
docker stop <registry-container-name>
docker rm <registry-container-name>
```

Delete images from your online registries. Storing them may incur costs.

# Docker Training for Junior DevOps Engineers

A hands-on Docker training program designed to take you from basic container commands to production-ready Docker Compose deployments.

## Prerequisites

- Docker Desktop installed and running
- Basic command line knowledge
- A text editor
- Curiosity and willingness to read documentation

## Learning Path

This training is structured as progressive exercises. Each builds on concepts from previous ones:

1. **basics** - Get comfortable with Docker commands using existing images
2. **building-images** - Create, optimize, and secure your own Docker images
3. **ports-networking** - Connect containers and expose services
4. **volumes-persistence** - Manage data that outlives containers
5. **registries** - Work with local and remote image registries
6. **compose-basics** - Orchestrate multi-container applications
7. **compose-advanced** - Production-ready deployments with scaling and security

## How to Use This Training

- Complete exercises in order
- Each exercise has a README with tasks
- Tasks have observable outcomes - you'll know when you've succeeded
- You'll need to read Docker documentation - links are provided
- Solutions are on the `solutions` branch (try without peeking first!)

## Getting Help

- Read the task carefully - the answer is often in what it's asking you to observe
- Check Docker documentation links provided in each exercise
- Use `docker --help` and `docker <command> --help`
- Search for error messages - they're usually informative
- If stuck for more than 30 minutes, check the solutions branch

## Tips

- `docker ps -a` shows all containers, not just running ones
- `docker images` shows all local images
- `docker logs <container>` is your friend for debugging
- `docker exec -it <container> sh` lets you explore inside containers
- Clean up regularly: `docker system prune`

Ready? Start with `01-basics/`!

# Exercise 03: Ports and Networking

Learn how containers communicate with the host and with each other.

## Objectives

Understand Docker networking:
- Port mapping and exposure
- Container-to-container communication
- Custom networks
- Network isolation
- DNS resolution between containers

## Tasks

### Task 1: Access Your App from the Host

Run your simple-web container (from exercise 02) and access it from your browser at http://localhost:8080

The container's port 8080 should be accessible from your host machine.

### Task 2: Run on a Different Port

Run another instance of simple-web, but this time access it at http://localhost:3000

Both containers should run simultaneously without conflicts.

### Task 3: Container-to-Container Communication

You'll need two containers:
1. A postgres database container
2. The API app from `../apps/api/`

Build and run the API container. It needs to connect to postgres using the hostname "database" (not an IP address).

Verify: The API should start without connection errors. Check the logs.

### Task 4: Network Isolation

With your setup from Task 3:
- The API should be able to reach the database
- Your host machine should be able to reach the API
- Your host machine should NOT be able to reach the database directly

Try to connect to postgres from your host on port 5432. It should fail.

But the API should still work fine.

### Task 5: Inspect the Network

Find out:
- What network are your containers on?
- What are their IP addresses?
- Can you ping one container from another?

Hint: Networks can be inspected, and you can execute commands inside containers.

### Task 6: Multiple Networks

Create a scenario with three containers:
- frontend (simple-web)
- backend (api)
- database (postgres)

Requirements:
- frontend can reach backend
- backend can reach database
- frontend CANNOT reach database

This requires two separate networks.

## Resources

- [Docker networking overview](https://docs.docker.com/network/)
- [Container networking](https://docs.docker.com/config/containers/container-networking/)
- [Bridge networks](https://docs.docker.com/network/bridge/)

## Verification Checklist

You've completed this exercise when:
- ✓ Can access container services from host on mapped ports
- ✓ Can run multiple containers on different host ports
- ✓ Containers can reach each other by name (DNS)
- ✓ Can isolate containers from host access
- ✓ Can inspect network configuration
- ✓ Can create multi-network isolation scenarios

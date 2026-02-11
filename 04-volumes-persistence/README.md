# Exercise 04: Volumes and Persistence

Learn how to manage data that needs to survive container restarts and how to share data between host and containers.

## Objectives

Master Docker data management:
- Named volumes for persistent data
- Bind mounts for host-container file sharing
- tmpfs mounts for temporary data
- Volume inspection and backup
- Understanding volume lifecycle

## Tasks

### Task 1: Data That Disappears

Run a postgres container and create some data:
```bash
# Connect to postgres and create a table with data
docker exec -it <container> psql -U postgres -c "CREATE TABLE test (id INT, name TEXT);"
docker exec -it <container> psql -U postgres -c "INSERT INTO test VALUES (1, 'data');"
```

Now stop and remove the container. Start a new postgres container.

Your data is gone. Fix this so data persists across container removals.

### Task 2: Verify Persistence

With your solution from Task 1:
1. Create some data in the database
2. Stop and remove the container
3. Start a new container
4. Your data should still be there

Verify by querying the database.

### Task 3: Inspect Volume Data

Find out:
- Where is your volume data actually stored on the host?
- How much space is it using?
- What other metadata does the volume have?

Hint: Volumes can be inspected just like containers and images.

### Task 4: Configuration Files

The API app needs a config file. Create a `config.json` file on your host:
```json
{
  "maxConnections": 100,
  "timeout": 5000,
  "logLevel": "debug"
}
```

Run the API container so that:
- This file appears at `/app/config.json` inside the container
- You can edit the file on your host and see changes immediately in the container
- The file should NOT be copied into the image

Verify: Edit the file on your host, then check it inside the container - changes should be instant.

### Task 5: Temporary Fast Storage

Some data needs to be fast but doesn't need to persist (like cache or temp files).

Run a container with a tmpfs mount at `/tmp/cache`. This storage:
- Lives in memory (very fast)
- Disappears when container stops
- Doesn't use disk space

Verify: Create files in `/tmp/cache`, check they exist, stop container, start again - files are gone.

### Task 6: Volume Backup

You have important data in a volume. Create a backup of it as a tar file on your host.

Then delete the volume and restore it from your backup.

Hint: You'll need to use a temporary container to access the volume data.

### Task 7: Read-Only Mounts

Mount a configuration file into a container as read-only. The container should be able to read it but not modify it.

Verify: Try to modify the file from inside the container - it should fail.

### Task 8: Multiple Containers, Same Volume

Run two containers that share the same volume. Write a file from one container and read it from the other.

This demonstrates how volumes can be used for container-to-container data sharing.

## Resources

- [Docker volumes](https://docs.docker.com/storage/volumes/)
- [Bind mounts](https://docs.docker.com/storage/bind-mounts/)
- [tmpfs mounts](https://docs.docker.com/storage/tmpfs/)
- [Volume backup and restore](https://docs.docker.com/storage/volumes/#back-up-restore-or-migrate-data-volumes)

## Verification Checklist

You've completed this exercise when:
- ✓ Database data persists across container removals
- ✓ Can inspect volume location and metadata
- ✓ Can mount host files into containers
- ✓ Changes to bind-mounted files are immediately visible
- ✓ Can use tmpfs for temporary storage
- ✓ Can backup and restore volume data
- ✓ Can mount files as read-only
- ✓ Can share volumes between containers

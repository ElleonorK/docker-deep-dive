# Exercise 04: Volumes and Persistence

Learn how to manage data that outlives containers and share files between host and containers.

## Objectives

By the end of this exercise, you'll be comfortable with:

*   Using named volumes for persistent data
*   Mounting host files and directories into containers
*   Understanding when data persists vs when it disappears
*   Backing up and restoring volume data
*   Sharing data between containers

## Tasks

### Task 1: The Disappearing Data Problem

Run a postgres database container:

```
docker run -d --name db -e POSTGRES_PASSWORD=secret postgres:16-alpine
```

Create some data in it:

```
docker exec -it db psql -U postgres -c "CREATE TABLE users (id INT, name TEXT);"
docker exec -it db psql -U postgres -c "INSERT INTO users VALUES (1, 'Alice'), (2, 'Bob');"
```

Verify your data exists:

```
docker exec -it db psql -U postgres -c "SELECT * FROM users;"
```

Now stop and remove the container:

```
docker stop db
docker rm db
```

Start a fresh postgres container with the same name and try to query your users table.

What happened to your data? Why?

### Task 2: Make Data Persist

**Your challenge:** Run a postgres container so that your data survives when you remove the container.

Postgres stores its data in `/var/lib/postgresql/data` inside the container. You need to make sure this directory persists.

Test it:

1.  Start a postgres container with a name like `db1`
2.  Create a table and insert some data
3.  Stop and remove the container
4.  Start a new postgres container with the same name (`db1`)
5.  Query your data - it should still be there

Now test with a different container name:

1.  Start another postgres container with a different name like `db2`
2.  Mount the same persisted data
3.  Query your data - it should be the same data from `db1`

**Success criteria:** You can access the same database data from containers with different names.

**Think:** Is the data tied to the container name or to something else?

### Task 3: Inspect Your Storage

Find out:

*   Where is your persisted data actually stored on your host machine?
*   How much disk space is it using?
*   What metadata can you discover about it?

**Hint:** Use `docker` commands to inspect what you created in Task 2.

### Task 4: Live Configuration Files

You need to provide a configuration file to a container, and you want to edit it without rebuilding the image or restarting the container.

Create a `config.json` file in this directory:

```
{
  "database": "postgres",
  "maxConnections": 100,
  "timeout": 5000
}
```

**Your challenge:** Run an nginx container so that:

*   This file appears at `/usr/share/nginx/html/config.json` inside the container
*   When you edit `config.json` on your host, the changes are immediately visible inside the container
*   You can access it at http://localhost:8080/config.json

**Success criteria:** Edit the file on your host (change `maxConnections` to 200), refresh the browser, and see the updated value without restarting the container.

**Think:** How is this different from the volume you used for postgres? When would you use each approach?

### Task 5: Temporary Fast Storage

Some data doesn't need to persist - cache files, temporary processing data, session storage. For this, you want speed, not durability.

**Your challenge:** Run an alpine container with a tmpfs mount at `/cache` that:

*   Lives entirely in RAM (no disk I/O)
*   Disappears when the container stops
*   Has a size limit of 100MB

Test it:

```
# Inside the container, create a file
docker exec <container> sh -c "echo 'temporary data' > /cache/test.txt"

# Verify it exists
docker exec <container> cat /cache/test.txt

# Restart the container
docker restart <container>

# Try to read the file again
docker exec <container> cat /cache/test.txt
```

What happened to the file after restart?

**Think:** Why would you want storage that disappears? What are the tradeoffs between tmpfs, volumes, and bind mounts?

### Task 6: Backup and Restore

You have critical data in a volume and need to back it up.

**Setup:** Create a volume with some data:

```
docker volume create important-data
docker run --rm -v important-data:/data alpine sh -c "echo 'critical information' > /data/backup-me.txt"
```

**Your challenge:**

1.  Create a backup of this volume as `backup.tar` on your host
2.  Delete the volume
3.  Restore the volume from your backup
4.  Verify the data is back

**Hint:** You'll need to use a temporary container to access the volume contents. Think about how you can mount both the volume and a host directory in the same container.

### Task 7: Read-Only Protection

Configuration files shouldn't be modified by the application - only read.

**Your challenge:** Create a `readonly-config.txt` file on your host, then run a container where:

*   The file is mounted inside the container
*   The container can read the file
*   The container CANNOT modify or delete the file

**Success criteria:** Try to modify the file from inside the container - it should fail with a "Read-only file system" error.

**Think:** Why is this useful for security? What could go wrong if an application could modify its own configuration?

### Task 8: Sharing Data Between Containers

Two containers need to share files - one writes logs, another processes them.

**Your challenge:** Set up two alpine containers:

*   Container 1 (writer): Continuously writes timestamps to `/shared/logs.txt`
*   Container 2 (reader): Can read `/shared/logs.txt`
*   Both containers should see the same file

**Success criteria:** Start the writer container, let it run for a few seconds, then exec into the reader container and see the log entries that the writer created.

**Hint:** Both containers need access to the same volume.

**Think:** How is this different from containers communicating over a network? When would you use shared volumes vs network communication?

## Resources

*   [Docker volumes documentation](https://docs.docker.com/storage/volumes/)
*   [Bind mounts documentation](https://docs.docker.com/storage/bind-mounts/)
*   [tmpfs mounts documentation](https://docs.docker.com/storage/tmpfs/)
*   [Backup and restore volumes](https://docs.docker.com/storage/volumes/#back-up-restore-or-migrate-data-volumes)

## Verification Checklist

You've completed this exercise when:

*   Database data persists across container removals
*   Can inspect volume location and metadata
*   Can mount host files and see live changes in containers
*   Can use tmpfs for temporary in-memory storage
*   Can backup and restore volume data
*   Can mount files as read-only
*   Can share volumes between multiple containers

## Cleanup

Remove all containers and volumes you created:

```
docker stop <container-name>
docker rm <container-name>
docker volume rm <volume-name>
```
# Exercise 04: Persistence and FileSystem

Learn how to manage data that outlives containers and share files between host and containers.

## Objectives

By the end of this exercise, you'll be comfortable with:
* Making container data persist across restarts
* Mounting host files and directories into containers
* Understanding when data persists vs when it disappears
* Understanding how Docker stores container data
* Sharing data between containers

## Tasks

### Task 1: The Disappearing Data Problem

Run a postgres database container:
```bash
docker run -d --name db -e POSTGRES_PASSWORD=secret postgres:16-alpine
```

Create some data in it:
```bash
docker exec -it db psql -U postgres -c "CREATE TABLE users (id INT, name TEXT);"
docker exec -it db psql -U postgres -c "INSERT INTO users VALUES (1, 'Alice'), (2, 'Bob');"
docker exec -it db psql -U postgres -c "SELECT * FROM users;"
```

Now stop and remove the container, then start a fresh postgres container with the same name and try to query your users table again.

What happened to your data? Why?

### Task 2: Make Data Persist

**Your challenge:** Run postgres containers so that your data survives when you remove the container.

**Hint:** Postgres stores its data in `/var/lib/postgresql/data` inside the container.

Start a container called `db1`, create a table with some data, then stop and remove it. Start a new container with the same name and verify your data is still there.

Now the real test - start another postgres container with a different name like `db2`. Can you make it see the same data from `db1`?

**Success criteria:** You can access the same database data from containers with different names.

**Think:** Is the data tied to the container name or to something else?

### Task 3: Understanding Container Storage

Let's explore how Docker actually stores data inside containers.

Pull and run a large image (this is a long download, ~3GB):
```bash
docker run -d --name storage-test kasmweb/rockylinux-9-desktop:1.18.0
```

Check the container's size with `docker ps -s`. Look at the SIZE column - you'll see two numbers. The first is the writable layer, the second is the total size including image layers (over 9GB!).

Now check what's taking up space inside the container `/opt` folder:
```bash
docker exec storage-test du -h --max-depth 1 /opt
```

You'll see applications like Telegram, Zoom, OnlyOffice taking up over 1GB. Delete them:
```bash
docker exec storage-test rm -rf /opt/Telegram /opt/zoom /opt/onlyoffice
```

Check `docker ps -s` again. Did the virtual size go down? Verify the folders are gone:
```bash
docker exec storage-test ls /opt
```

Stop and remove this container. 

**Your challenge:** Start a new container, but this time mount `/opt` to a directory on your host. Now delete those application folders (Telegram, Zoom, OnlyOffice) directly from your host filesystem. Check the size of the folder on your host with `du -sh`. Then check the container's disk usage with `docker ps -s`.
>[!Tip] If you are working inside WSL, you may have trouble mounting the directory to host. Search for 'drivers' in documentation.

**Success criteria:** The folder on your host got smaller (by over 1GB), but the container's disk usage didn't change.

**Think:** Why can't you reclaim space by deleting files that came with the image? What's the difference between the container's view of the filesystem and what's actually stored on disk? When would mounting to the host actually help you save space?

**Resources:** Read about [Docker storage drivers](https://docs.docker.com/storage/storagedriver/) and [OverlayFS](https://docs.docker.com/storage/storagedriver/overlayfs-driver/) to understand what's happening.

### Task 4: Live Code Editing

You're going to edit a live website without restarting anything.

**Your challenge:** Run an nginx container serving the frontend app from `apps/frontend/`. When you edit the HTML file on your host, the changes should appear instantly in the browser at `http://localhost:8080`.

Try making these changes and watch them appear live:
* Change the title from "Docker Training Application" to "My Awesome App"
* Change the background color in the CSS (find `background: #f5f5f5;` and try `#e3f2fd;`)
* Add a new button or change button text

**Success criteria:** Every edit you make appears instantly in the browser without restarting the container.

**Think:** How is this different from what you used for postgres? When would you use each approach? Why is this useful for development?

### Task 5: Configuration Management

You have a postgres database that needs different configuration for development vs production.

Create two config files in this directory:

`dev-postgres.conf`:
```
max_connections = 20
shared_buffers = 128MB
```

`prod-postgres.conf`:
```
max_connections = 200
shared_buffers = 2GB
```

**Your challenge:** Run two postgres containers - one for dev, one for prod - each using its respective config file. The config should be at `/etc/postgresql/postgresql.conf` inside each container.

**Success criteria:** Exec into each container and verify they're using different configurations:
```bash
docker exec <container> cat /etc/postgresql/postgresql.conf
```

**Think:** Why mount config files instead of baking them into the image? What happens when you need to change a setting?

### Task 6: Read-Only Protection

**Your challenge:** Run an nginx container serving the frontend app where the HTML files are mounted from your host but the container cannot modify them.

**Success criteria:** Try to modify the HTML file from inside the container - it should fail with a "Read-only file system" error.

**Think:** Why is this useful for security? What could go wrong if an application could modify its own files?

### Task 7: Log Aggregation

Set up a realistic scenario: nginx writing access logs that another container can analyze.

**Your challenge:** Set up two containers:
* Container 1: nginx serving the frontend app and writing logs
* Container 2: Use `busybox` or `alpine` to read and tail those nginx logs

**Success criteria:** Visit `http://localhost:8080` a few times, then exec into the second container and use `tail -f` to watch the nginx access logs in real-time.

**Think:** How is this different from containers communicating over a network? When would you use shared storage vs network communication? What are real-world examples of this pattern?

## Resources

* [Docker storage overview](https://docs.docker.com/storage/)
* [Storage drivers](https://docs.docker.com/storage/storagedriver/)
* [OverlayFS driver](https://docs.docker.com/storage/storagedriver/overlayfs-driver/)

## Verification Checklist

You've completed this exercise when:
* [ ] Database data persists across container removals
* [ ] Understand how Docker's storage layers work
* [ ] Can mount host files and see live changes in containers
* [ ] Can mount different config files to different containers
* [ ] Can mount files as read-only
* [ ] Can share data between multiple containers

## Cleanup

Remove all containers you created:
```bash
docker stop <container-name>
docker rm <container-name>
```
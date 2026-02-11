# Setup Instructions

## For Instructors/Maintainers

### Initial Git Setup

```bash
# Initialize repository
git init

# Add all files
git add .

# Create initial commit
git commit -m "Initial commit: Docker training exercises"

# Create solutions branch
git checkout -b solutions

# Add solution files (see below)
# ... add Dockerfiles, docker-compose.yml files, etc.

git add .
git commit -m "Add solutions for all exercises"

# Switch back to main
git checkout main

# Optional: Set up remote and push
git remote add origin <your-repo-url>
git push -u origin main
git push -u origin solutions
```

### Protecting the Solutions Branch

On GitHub/GitLab:
1. Go to repository settings
2. Navigate to Branches
3. Add branch protection rule for `solutions`
4. Enable "Restrict who can push to matching branches"
5. Only allow instructors/maintainers

## For Students

### Prerequisites

1. Install Docker Desktop: https://www.docker.com/products/docker-desktop
2. Verify installation:
   ```bash
   docker --version
   docker compose version
   ```
3. Clone this repository:
   ```bash
   git clone <repo-url>
   cd docker-training
   ```

### Getting Started

1. Read the main README.md
2. Start with `01-basics/`
3. Complete exercises in order
4. Try to solve without looking at solutions
5. If stuck, check the `solutions` branch:
   ```bash
   git checkout solutions
   # Look at the solution
   git checkout main  # Go back to exercises
   ```

### Tips

- Don't copy-paste solutions - type them out to learn
- Experiment and break things - that's how you learn
- Use `docker system prune` regularly to clean up
- Read error messages carefully - they're usually helpful

## Solution Files to Add

When creating the solutions branch, add these files:

### 02-building-images/
- `Dockerfile` - Complete, optimized Dockerfile with all best practices
- `.dockerignore` - Proper ignore patterns

### 03-ports-networking/
- `commands.sh` - Example commands for all tasks
- `notes.md` - Explanations of networking concepts

### 04-volumes-persistence/
- `commands.sh` - Example commands for volume operations
- `backup-restore.sh` - Scripts for backup/restore tasks

### 05-registries/
- `commands.sh` - Registry setup and operations
- `registry-notes.txt` - Answers to naming convention questions

### 06-compose-basics/
- `docker-compose.yml` - Complete compose file
- `.env.example` - Example environment variables

### 07-compose-advanced/
- `docker-compose.yml` - Base configuration
- `docker-compose.dev.yml` - Development overrides
- `docker-compose.prod.yml` - Production overrides
- `.env.example` - Example environment variables
- `README-SOLUTION.md` - Detailed explanations

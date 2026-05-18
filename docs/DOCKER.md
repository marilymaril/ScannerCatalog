# Docker Setup for Scanner Catalog

## Overview

This project uses Docker to run a PostgreSQL database container. The database is accessible on your local machine at `localhost:5432`.

## Current Container

| Property | Value |
|----------|-------|
| Container Name | `catalogproj-db` |
| Image | `postgres` |
| Status | Running |
| Port | `5432` (local) → `5432` (container) |

## Port Mapping

The Docker container maps ports as follows:
- **Host Port**: `0.0.0.0:5432` (accessible on your machine)
- **IPv6**: `[::]:5432` (IPv6 equivalent)
- **Container Port**: `5432` (internal PostgreSQL port)

This allows you to connect to PostgreSQL from your application using:
```
localhost:5432
```

## Common Docker Commands

### View Running Containers
```bash
docker ps
```

### Stop the Container
```bash
docker stop catalogproj-db
```

### Start the Container
```bash
docker start catalogproj-db
```

### View Container Logs
```bash
docker logs catalogproj-db
```

### Remove the Container
```bash
docker rm catalogproj-db
```

## Stopping Docker

### Stop Just the Container
```bash
docker stop catalogproj-db
```

### Stop the Docker Daemon (Linux)
```bash
sudo systemctl stop docker
```

### Restart Docker (Linux)
```bash
sudo systemctl start docker
```

## Connection String

To connect to the PostgreSQL database:

```
postgres://user:password@localhost:5432/dbname
```

Update credentials and database name as needed for your environment.

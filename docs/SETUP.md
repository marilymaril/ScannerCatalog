# Project Setup Guide

## Prerequisites

- Go 1.26.1 or later
- Docker (for database)
- Docker Compose (optional, for orchestration)

## Getting Started

### 1. Clone the Repository
```bash
git clone https://github.com/marilymaril/ScannerCatalog.git
cd ScannerCatalog
```

### 2. Check Go Installation
```bash
go version
```

### 3. Start Docker Container
The PostgreSQL database container is already configured to run:

```bash
docker ps
```

See [DOCKER.md](./DOCKER.md) for detailed Docker instructions.

### 4. Install Swaggo (for API documentation)
```bash
go install github.com/swaggo/swag/cmd/swag@latest
```

### 5. Generate Swagger Documentation
From the project root:

```bash
swag init -g ./cmd/scannerApp/main.go
```

This generates:
- `docs/docs.go`
- `docs/swagger.json`
- `docs/swagger.yaml`

### 6. Build and Run the Application

**Build**:
```bash
go build -o ./build/scannerApp ./cmd/scannerApp/main.go
```

**Run**:
```bash
go run ./cmd/scannerApp/main.go
```

The API will start on `http://localhost:8080`

## Project Structure

```
.
├── cmd/scannerApp/       # Main application entry point
├── internal/             # Internal packages (not exported)
├── pkg/                  # Public packages
├── api/                  # API definitions
├── configs/              # Configuration files
├── deploy/               # Deployment files
├── docs/                 # Documentation
├── build/                # Build output
├── test/                 # Tests
└── go.mod               # Go module definition
```

## Key Files

- `cmd/scannerApp/main.go` — Application entry point with API handlers
- `go.mod` — Go module dependencies

## Next Steps

1. Read [DOCKER.md](./DOCKER.md) for database setup
2. Uncomment the HTTP server code in `cmd/scannerApp/main.go` to start the API
3. Visit `http://localhost:8080/swagger/index.html` to view API docs (after uncommenting handler and starting server)

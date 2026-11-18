# ToDo App

Plan: A REST API built with Go and Gin which uses PostgreSQL as the database and TS with React and Vite / HTML

## Tech Stack

- **Backend:** Go, Gin Gonic planned for later
- **Database:** PostgreSQL planned for later
- **Tooling & Environment:** Docker & Docker Compose, Air (Live Reload) planned for later
- **Quality & Security:** SonarQube planned for later
- **Observability:** Prometheus & Grafana planned for later
- **Testing:** Mailpit (Email), Bruno (API Client) planned for later

## Infrastructure Ports

When running the environment, the following services are available:
- 🚀 **Go API:** `http://localhost:8080` (or your configured port)
- 📊 **Grafana:** `http://localhost:3000`
- 📈 **Prometheus:** `http://localhost:9090`
- 📬 **Mailpit:** `http://localhost:8025`
- 🦊 **SonarQube:** `http://localhost:9000`

## Getting Started

### Prerequisites

Make sure you have [Docker](https://docker.com) and [Air](https://github.com) installed:
```bash
go install ://github.com
```

### 1. Start Infrastructure
Spin up PostgreSQL, Mailpit, SonarQube, and the monitoring tools using Docker Compose:
```bash
docker compose up -d
```

### 2. Run the Application
Start the Go backend with live-reloading enabled:
```bash
air
```

### For now
swtich to
```bash
cd ./cmd/api
```
and type
```bash
go run .
```
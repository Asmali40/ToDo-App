# ToDo App

A REST API built with Go and Gin which uses PostgreSQL as the database.

## Tech Stack

- **Backend:** Go, Gin Gonic
- **Database:** PostgreSQL
- **Tooling & Environment:** Docker & Docker Compose, Air (Live Reload)
- **Quality & Security:** SonarQube
- **Observability:** Prometheus & Grafana
- **Testing:** Mailpit (Email), Bruno (API Client)

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
# MyGoWebserver

A RESTful API built with Go and PostgreSQL for managing posts.

## Prerequisites

- Go 1.21+
- Docker

## Setup

### 1. Clone the repository

```bash
git clone https://github.com/djedd1ne/MyGoWebserver.git
```
```bash
cd MyGoWebserver
```

### 2.Configure Environment Variables
Create a .env file in the root directory:
```.env
DB_HOST=localhost
DB_PORT=6543
DB_USER=postgres
DB_PASSWORD=
DB_NAME=postgres
```
### 3.Install Dependencies
```bash
make deps
```
Or manually:
```bash
go mod tidy
```
#### Makefile Commands

| Command | Description |
| :--- | :--- |
| `make deps` | Install Go dependencies |
| `make build` | Build the application |
| `make run` | Run the application |
| `make docker-build` | Build PostgreSQL Docker image |
| `make docker-run` | Build and run PostgreSQL container |
| `make docker-stop` | Stop PostgreSQL container |
| `make docker-logs` | View container logs |
| `make psql` | Open PostgreSQL shell |
| `make docker-clean` | Remove container and image |
| `make clean` | Remove container, image, and binary |
| `make setup` | Build and run database, then wait for ready |

## Run
### 1.Start database
```bash
make docker-run
```

### 2.Run server
```bash
make run
```

# REST API for backend-education coursework using Gin, Reform and PostgreSQL

This is a REST API written in Go using Gin for server creation, Reform for database interaction and PostgreSQL as our
database.

## Requirements

- Go (version 1.16 or higher)
- PostgreSQL (version 12 or higher)

## Installation

1. Clone the repository:
    ```bash
    git clone https://github.com/Loadngg/BackendEducationCoursework.git backend-education-coursework
    ```
2. Change to the project directory:
    ```bash
    cd backend-education-coursework
    ```
3. Install dependencies:
    ```bash
    go mod tidy
    ```
4. Create a `config.yaml` file in `config` dir with settings:

```yaml
env: "local"
server:
  address: localhost:5000
  timeout: 4s
  idle_timeout: 60s
db:
  host: 127.0.0.1
  port: port
  user: user
  password: password
  database: db_name
  ssl: disable

```

## Running

Run the application:

```sh
go run ./cmd/backend/main.go
```

The API will be available at `http://localhost:5000`.
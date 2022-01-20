# Tracker app

This app intended to track user statistics of valid/invalid API calls

# Install

API build on top of the Go, Postgres, Chi Router and sqlx. All infra tasks (start API, install deps, generater Swagger etc.) are managed by [Task](https://taskfile.dev/) cli. To install it run the following command:

```bash
sudo ./scripts/install-task.sh
```

This command installs `task` to `/user/local/bin/task` so sudo is needed.

1. Initiliaze `.env` file with development defaults by:

```bash
task init:env
```

2. Install additional tools:

```bash
task install:tools
```

3. Start database

```bash
task start:postgres
```

4. Apply migrations

```bash
go run cmd/migrate/main.go
```

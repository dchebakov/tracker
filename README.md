# Tracker app

This app is intended to track user statistics of valid/invalid API calls.

## Install

API build on top of the Go, Echo, sqlx and Postgres. All infra tasks (start API, install deps, etc.) are managed by [Task](https://taskfile.dev/) cli. The easiest way to start API is by using docker.

1. Initiliaze `.env` file with development defaults by:

```bash
cp env.example.docker .env
```

2. Start database:

```bash
docker-compose --env-file .env up -d postgres
```

3. Apply migrations:

```bash
docker-compose --env-file .env up --build migrate
```

4. Start API server

```bash
docker-compose --env-file .env up --build api
```

5. Check API is alive

```bash
curl localhost:6000/api/v1/health/
```

## API Methods

Examples of API calls can be found in the [examples](./examples) directory. Some of them are:

```bash
# Valid collector requst
curl --request POST \
  --url http://localhost:6000/api/v1/collect/ \
  --header 'content-type: application/json' \
  --data '{"customerID": 1,"tagID": 2,"userID": "aaaaaaaa-bbbb-cccc-1111-222222222222","remoteIP": "123.234.56.78","timestamp": 1500000000}'

# Return all stats
curl --request GET \
  --url http://localhost:6000/api/v1/stats/
```

Files in the examples directory with an extenstion `.http` are intended to use with the [REST Client](https://marketplace.visualstudio.com/items?itemName=humao.rest-client) extension for vscode.

Files in the examples directory with an extenstion `.md` contains curl examples.

## Development

Setup for the development guide can be found in the [docs](./docs/dev.md) directory.

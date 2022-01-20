### Health API

Health API requests examples using curl cli.

- Health server check

```bash
curl --request GET \
  --url http://localhost:6000/api/v1/health/
```

###

- DB connection check

```bash
curl --request GET \
  --url http://localhost:6000/api/v1/health/readiness
```

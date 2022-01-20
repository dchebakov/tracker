### Health API

Stats API requests examples using curl cli.

- Return all stats

```bash
curl --request GET \
  --url http://localhost:6000/api/v1/stats/
```

- Filter stats for a given customer

```bash
curl --request GET \
  --url 'http://localhost:6000/api/v1/stats/?customerID=1'
```

- Filter stats for the specific day

```bash
curl --request GET \
  --url 'http://localhost:6000/api/v1/stats/?day=2017-07-14'
```

- Filter stats both for the customer and day

```bash
curl --request GET \
  --url 'http://localhost:6000/api/v1/stats/?customerID=1&day=2017-07-14'
```

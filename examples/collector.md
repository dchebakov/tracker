### Collector API

Collector API requests examples using curl cli.

- Valid collector requst

```bash
curl --request POST \
  --url http://localhost:6000/api/v1/collect/ \
  --header 'content-type: application/json' \
  --data '{"customerID": 1,"tagID": 2,"userID": "aaaaaaaa-bbbb-cccc-1111-222222222222","remoteIP": "123.234.56.78","timestamp": 1500000000}'
```

- Malformed JSON call

```bash
curl --request POST \
  --url http://localhost:6000/api/v1/collect/ \
  --header 'content-type: application/json' \
  --data '"{"'
```

- Missed userID field, saves API call as invalid

```bash
curl --request POST \
  --url http://localhost:6000/api/v1/collect/ \
  --header 'content-type: application/json' \
  --data '{"customerID": 1,"tagID": 2,"remoteIP": "123.234.56.78","timestamp": 1500000000}'
```

- Missed customerID filed so the collector wouldn't update any DB record

```bash
curl --request POST \
  --url http://localhost:6000/api/v1/collect/ \
  --header 'content-type: application/json' \
  --data '{"tagID": 2,"userID": "aaaaaaaa-bbbb-cccc-1111-222222222222","remoteIP": "123.234.56.78","timestamp": 1500000000}'
```

- Non-existing customer so the collector wouldn't update any DB record

```bash
curl --request POST \
  --url http://localhost:6000/api/v1/collect/ \
  --header 'content-type: application/json' \
  --data '{"customerID": 1000000,"tagID": 2,"userID": "aaaaaaaa-bbbb-cccc-1111-222222222222","remoteIP": "123.234.56.78","timestamp": 1500000000}'
```

- Non-actie user so the request will be counted as invalid

```bash
curl --request POST \
  --url http://localhost:6000/api/v1/collect/ \
  --header 'content-type: application/json' \
  --data '{"customerID": 3,"tagID": 2,"userID": "aaaaaaaa-bbbb-cccc-1111-222222222222","remoteIP": "123.234.56.78","timestamp": 1500000000}'
```

- Blocked IP so the request will be counted as invalid

```bash
curl --request POST \
  --url http://localhost:6000/api/v1/collect/ \
  --header 'content-type: application/json' \
  --data '{"customerID": 1,"tagID": 2,"userID": "aaaaaaaa-bbbb-cccc-1111-222222222222","remoteIP": "213.070.64.33","timestamp": 1500000000}'
```

- Blocked User Agent so the request will be counted as invalid

```bash
curl --request POST \
  --url http://localhost:6000/api/v1/collect/ \
  --header 'content-type: application/json' \
  --header 'user-agent: Googlebot' \
  --data '{"customerID": 1,"tagID": 2,"userID": "aaaaaaaa-bbbb-cccc-1111-222222222222","remoteIP": "123.234.56.78","timestamp": 1500000000}'
```

- Making valid API call for the next hour (2017-07-14 03:16:00), so new DB record will be created

```bash
curl --request POST \
  --url http://localhost:6000/api/v1/collect/ \
  --header 'content-type: application/json' \
  --data '{"customerID": 1,"tagID": 2,"userID": "aaaaaaaa-bbbb-cccc-1111-222222222222","remoteIP": "123.234.56.78","timestamp": 1500002160}'
```

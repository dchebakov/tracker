package repository

const createHourStatsQuery = `--sql
INSERT INTO hourly_stats (customer_id, time, request_count, invalid_count)
VALUES (
    $1,
    date_trunc('hour', to_timestamp($2)),
    $3,
    $4
  )
RETURNING *;
`

const getHourStatsQuery = `--sql
SELECT id,
  customer_id,
  time,
  request_count,
  invalid_count
FROM hourly_stats
WHERE customer_id = $1
  AND time >= date_trunc('hour', to_timestamp($2))
  AND time < date_trunc('hour', to_timestamp($2)) + INTERVAL '1' HOUR;
`

const updateStatsQuery = `--sql
UPDATE hourly_stats
SET request_count = request_count + $2,
  invalid_count = invalid_count + $3
WHERE id = $1;
`

const getAllStats = `--sql
SELECT id,
  customer_id,
  time,
  request_count,
  invalid_count
FROM hourly_stats
ORDER BY time;
`

const getCustomerStats = `--sql
SELECT id,
  customer_id,
  time,
  request_count,
  invalid_count
FROM hourly_stats
WHERE customer_id = $1
ORDER BY time;
`

const getDayStats = `--sql
SELECT id,
  customer_id,
  time,
  request_count,
  invalid_count
FROM hourly_stats
WHERE time >= date_trunc('day', to_timestamp($1))
  AND time < date_trunc('day', to_timestamp($1)) + INTERVAL '1' DAY;
`

const getCustomerDayStats = `--sql
SELECT id,
  customer_id,
  time,
  request_count,
  invalid_count
FROM hourly_stats
WHERE customer_id = $1
  AND time >= date_trunc('day', to_timestamp($2))
  AND time < date_trunc('day', to_timestamp($2)) + INTERVAL '1' DAY;
`

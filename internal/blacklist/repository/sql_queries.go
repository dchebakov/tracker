package repository

const getIpQuery = `--sql
SELECT ip
FROM ip_blacklist
WHERE ip = $1;
`

const getUaQuery = `--sql
SELECT ua
FROM ua_blacklist
WHERE ua = $1;
`

package repository

const getCustomerQuery = `--sql
SELECT id,
  name,
  active
FROM customer
WHERE id = $1;
`

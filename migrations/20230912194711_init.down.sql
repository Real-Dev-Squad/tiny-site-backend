BEGIN;

-- Check and drop the tiny_url table first if it exists.
DROP TABLE IF EXISTS tiny_url;

-- Now it is safe to drop the users table.
DROP TABLE IF EXISTS users;

COMMIT;

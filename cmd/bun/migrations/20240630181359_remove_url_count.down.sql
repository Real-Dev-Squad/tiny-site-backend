BEGIN;
--bun:split

ALTER TABLE users DROP COLUMN url_count;
--bun:split

COMMIT;
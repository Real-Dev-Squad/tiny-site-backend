BEGIN;
--bun:split

ALTER TABLE USERS DROP COLUMN url_count;
--bun:split

COMMIT;
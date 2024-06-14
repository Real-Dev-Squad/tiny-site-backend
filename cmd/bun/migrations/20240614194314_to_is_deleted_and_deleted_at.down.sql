BEGIN;
--bun:split

ALTER TABLE tiny_url DROP COLUMN is_deleted;
--bun:split

ALTER TABLE tiny_url DROP COLUMN  deleted_at;

COMMIT;
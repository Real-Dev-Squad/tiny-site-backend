BEGIN;

--bun:split

ALTER TABLE tiny_url ADD  is_deleted bool null DEFAULT FALSE;

--bun:split

ALTER TABLE tiny_url ADD  deleted_at timestamp null;

COMMIT;
BEGIN;

--bun:split

ALTER TABLE USERS ADD url_count  int NOT NULL DEFAULT 0;

--bun:split

COMMIT;
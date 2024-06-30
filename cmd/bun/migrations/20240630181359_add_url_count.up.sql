BEGIN;

--bun:split

ALTER TABLE USERS ADD url_count bool  int NOT NULL DEFAULT 0;

--bun:split

COMMIT;
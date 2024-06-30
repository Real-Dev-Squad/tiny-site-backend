BEGIN;

--bun:split

ALTER TABLE users ADD  url_count bool  int NOT NULL DEFAULT 0;

--bun:split

COMMIT;
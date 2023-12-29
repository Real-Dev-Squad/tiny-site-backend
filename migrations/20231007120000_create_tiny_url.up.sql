BEGIN;

CREATE TABLE tiny_url (
  id bigserial PRIMARY KEY,
  original_url text NOT NULL,
  short_url text UNIQUE NOT NULL,
  comment text,
  user_id int NOT NULL REFERENCES users(id), 
  expired_at timestamp NOT NULL,
  created_at timestamp DEFAULT (NOW() AT TIME ZONE 'UTC'),
  created_by text NOT NULL,
  access_count bigint DEFAULT 0,
  last_accessed_at timestamp DEFAULT (NOW() AT TIME ZONE 'UTC')
);

COMMIT;

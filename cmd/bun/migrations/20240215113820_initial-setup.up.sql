BEGIN;

--bun:split

CREATE TABLE users (
  id bigserial PRIMARY KEY,
  username varchar(256) UNIQUE NOT NULL,
  email varchar UNIQUE NOT NULL,
  is_verified boolean DEFAULT false,
  password varchar(128) NOT NULL,
  created_at timestamp DEFAULT (NOW() AT TIME ZONE 'UTC'),
  updated_at timestamp DEFAULT (NOW() AT TIME ZONE 'UTC'),
  is_deleted boolean DEFAULT false,
  is_onboarding BOOLEAN NOT NULL DEFAULT TRUE
);

--bun:split

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

--bun:split

COMMIT;

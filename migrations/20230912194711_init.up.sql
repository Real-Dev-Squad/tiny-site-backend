BEGIN;

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

COMMIT;

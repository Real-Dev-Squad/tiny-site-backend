BEGIN;

-- Create users table if not exists
CREATE TABLE IF NOT EXISTS users (
  id bigserial PRIMARY KEY,
  username varchar(256) UNIQUE NOT NULL,
  email varchar(255) UNIQUE NOT NULL,
  is_verified boolean DEFAULT false,
  password varchar(128) NOT NULL,
  is_deleted boolean DEFAULT false,
  is_onboarding BOOLEAN NOT NULL DEFAULT TRUE,
  created_at timestamp WITH TIME ZONE DEFAULT (NOW() AT TIME ZONE 'UTC'),
  updated_at timestamp WITH TIME ZONE DEFAULT (NOW() AT TIME ZONE 'UTC')
);

-- Insert data into the users table
INSERT INTO users (username, email, is_verified, password, is_onboarding)
VALUES ('JohnDoe', 'john.doe@example.com', true, 'hashed_password', true);

-- Create tiny_url table if not exists
CREATE TABLE IF NOT EXISTS tiny_url (
  id bigserial PRIMARY KEY,
  original_url text NOT NULL,
  short_url text UNIQUE NOT NULL,
  comment text,
  user_id int NOT NULL REFERENCES users(id), 
  expired_at timestamp WITH TIME ZONE NOT NULL, -- Consider allowing NULL or setting a default
  created_at timestamp WITH TIME ZONE DEFAULT (NOW() AT TIME ZONE 'UTC'),
  created_by text NOT NULL,
  access_count bigint DEFAULT 0,
  last_accessed_at timestamp DEFAULT (NOW() AT TIME ZONE 'UTC')
  isDeleted bool null default false
  deleted_at timestamp WITH TIME ZONE NULL
);

-- Insert data into the tiny_url table
INSERT INTO tiny_url (original_url, short_url, comment, user_id, expired_at, created_by, access_count, last_accessed_at)
VALUES ('https://www.example.com/1', '37fff', 'Some comment', 1, '2023-01-01T00:00:00Z', 'JohnDoe', 0, '2023-01-01T00:00:00Z');

COMMIT;

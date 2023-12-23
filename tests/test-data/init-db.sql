BEGIN;

-- Create users table if not exists
CREATE TABLE IF NOT EXISTS users (
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

-- Insert data into the users table
INSERT INTO users (id, username, email, is_verified, password, is_onboarding)
VALUES (1, 'JohnDoe', 'john.doe@example.com', true, 'hashed_password', true);

-- Create tiny_url table if not exists
CREATE TABLE IF NOT EXISTS tiny_url (
  id bigserial PRIMARY KEY,
  original_url text NOT NULL,
  short_url text UNIQUE NOT NULL,
  comment text,
  user_id int NOT NULL REFERENCES users(id), 
  expired_at timestamp NOT NULL,
  created_at timestamp DEFAULT (NOW() AT TIME ZONE 'UTC'),
  created_by text NOT NULL
);

-- Insert data into the tiny_url table
INSERT INTO tiny_url ( id, original_url, short_url, comment, user_id, expired_at, created_by)
VALUES (1, 'https://www.example.com', '37fff02c', 'Some comment', 1, '2023-01-01', 'JohnDoe');

COMMIT;

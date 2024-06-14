BEGIN;

INSERT INTO users (id, username, email, password) VALUES (3, 'another user', 'another@gmail.com', 'anotherpasswordhere');
INSERT INTO users (id, username, email, password) VALUES (4, 'karla', 'karla@gmail.com', 'karlaisanewuser');

COMMIT;

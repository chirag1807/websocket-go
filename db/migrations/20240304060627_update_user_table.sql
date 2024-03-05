-- migrate:up
ALTER TABLE users ALTER COLUMN Password TYPE VARCHAR(50);

-- migrate:down
DROP TABLE IF EXISTS users;

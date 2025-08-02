CREATE TABLE IF NOT EXISTS users (
    id VARCHAR PRIMARY KEY,
    name text NOT NULL,
    email text NOT NULL,
    password text NOT NULL,
    unique(email)

);

CREATE TABLE IF NOT EXISTS tasks(
    id VARCHAR PRIMARY KEY,
    title text NOT NULL,
    description text NOT NULL,
    status text NOT NULL DEFAULT 'Новая',
    user_id VARCHAR NOT NULL,
    unique(title)
);


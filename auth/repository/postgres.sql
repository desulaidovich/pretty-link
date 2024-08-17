DROP TABLE IF EXISTS Account CASCADE;

CREATE TABLE Account (
    id SERIAL PRIMARY KEY,
    email CHARACTER VARYING(255) UNIQUE,
    password CHARACTER VARYING(255),
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    CHECK (
        (email != '')
        AND (password != '')
    )
);

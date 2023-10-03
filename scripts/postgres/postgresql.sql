CREATE EXTENSION IF NOT EXISTS "pg_trgm";

DROP TABLE IF EXISTS people;

CREATE TABLE IF NOT EXISTS
    people (
        id UUID PRIMARY KEY NOT NULL,
        nickname VARCHAR(32) UNIQUE NOT NULL,
        "name" VARCHAR(100) NOT NULL,
        birthdate DATE NOT NULL,
        search TEXT NULL,
        stack TEXT NULL
    );

CREATE INDEX CONCURRENTLY IF NOT EXISTS people_index ON people USING
    gist (
        search gist_trgm_ops(siglen = 64)
    );
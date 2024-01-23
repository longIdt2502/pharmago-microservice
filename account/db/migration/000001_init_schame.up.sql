CREATE TABLE "accounts" (
    "id" serial PRIMARY KEY,
    "username" varchar UNIQUE NOT NULL,
    "hashed_password" varchar NOT NULL,
    "full_name" varchar NOT NULL,
    "email" varchar UNIQUE NOT NULL,
    "type" serial NOT NULL,
    "is_verify" boolean NOT NULL DEFAULT false,
    "password_changed_at" timestamptz NOT NULL DEFAULT '0001-01-01 00:00:00Z',
    "created_at" timestamptz NOT NULL DEFAULT (now())
);

CREATE TABLE "account_type" (
    "id" serial PRIMARY KEY,
    "code" varchar NOT NULL,
    "title" varchar NOT NULL
);

CREATE TABLE "verifies" (
    "id" serial PRIMARY KEY,
    "username" varchar NOT NULL,
    "email" varchar NOT NULL,
    "secret_code" varchar NOT NULL,
    "is_used" bool NOT NULL DEFAULT false,
    "created_at" timestamptz NOT NULL DEFAULT (now()),
    "expired_at" timestamptz NOT NULL DEFAULT ((now() + interval '15 minutes'))
);

ALTER TABLE "verifies" ADD FOREIGN KEY ("username") REFERENCES "accounts" ("username") ON DELETE CASCADE;
ALTER TABLE "accounts" ADD FOREIGN KEY ("type") REFERENCES "account_type" ("id") ON DELETE CASCADE;
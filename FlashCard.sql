CREATE TABLE "user" (
  "id" serial PRIMARY KEY,
  "email" varchar,
  "name" varchar,
  "password" varchar,
  "created_at" timestamptz DEFAULT (now()),
  "updated_at" timestamptz DEFAULT (now())
);

CREATE TABLE "vocabulary" (
  "id" serial PRIMARY KEY,
  "new_word" varchar,
  "user_id" serial,
  "status" integer,
  "count_status" integer,
  "image" varchar,
  "meaning" text,
  "created_at" timestamptz DEFAULT (now()),
  "updated_at" timestamptz DEFAULT (now())
);

CREATE TABLE "setting" (
  "id" serial PRIMARY KEY,
  "user_id" serial,
  "allow_send_email" boolean,
  "repetition_rule" jsonb,
  "vocabulary_number" integer,
  "created_at" timestamptz DEFAULT (now()),
  "updated_at" timestamptz DEFAULT (now())
);

ALTER TABLE
  "vocabulary"
ADD
  FOREIGN KEY ("user_id") REFERENCES "user" ("id");

ALTER TABLE
  "setting"
ADD
  FOREIGN KEY ("user_id") REFERENCES "user" ("id");
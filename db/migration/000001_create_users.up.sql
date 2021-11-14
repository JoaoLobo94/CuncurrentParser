CREATE TABLE "public"."users" (
    "id" integer GENERATED ALWAYS AS IDENTITY,
    "name" varchar NOT NULL,
    "created_at" timestamp without time zone NOT NULL DEFAULT (now()),
    PRIMARY KEY ("id")
);
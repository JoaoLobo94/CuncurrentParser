CREATE TABLE "public"."users" (
    "id" integer GENERATED ALWAYS AS IDENTITY,
    "name" varchar,
    "created_at" timestamp without time zone NOT NULL,
    PRIMARY KEY ("id")
);
CREATE TABLE "public"."batches" (
    "id" integer GENERATED ALWAYS AS IDENTITY,
    "dispatched" boolean NOT NULL DEFAULT FALSE,
    "amount" float NOT NULL DEFAULT '0.0',
    "user_id" integer NOT NULL,
    "created_at" timestamp without time zone NOT NULL DEFAULT (now()),
    PRIMARY KEY ("id"),
    FOREIGN KEY ("user_id") REFERENCES "public"."users"("id")
);

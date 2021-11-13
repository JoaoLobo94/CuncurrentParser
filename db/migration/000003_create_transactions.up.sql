CREATE TABLE "public"."transactions" (
    "id" integer GENERATED ALWAYS AS IDENTITY,
    "amount" float,
    "user_id" bigint,
    "created_at" timestamp without time zone NOT NULL DEFAULT (now()),
    PRIMARY KEY ("id"),
    CONSTRAINT "user_id" FOREIGN KEY ("user_id") REFERENCES "public"."users"("id")
);

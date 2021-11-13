CREATE TABLE "public"."bank_transactions" (
    "id" integer GENERATED ALWAYS AS IDENTITY,
    "amount" float,
    "user_id" integer NOT NULL,
    "created_at" timestamp without time zone NOT NULL DEFAULT (now()),
    PRIMARY KEY ("id"),
    FOREIGN KEY ("user_id") REFERENCES "public"."users"("id")
);

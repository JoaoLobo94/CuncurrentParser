CREATE TABLE "public"."bankTransactions" (
    "id" integer GENERATED ALWAYS AS IDENTITY,
    "amount" float,
    "user_id" integer,
    PRIMARY KEY ("id"),
    CONSTRAINT "user_id" FOREIGN KEY ("user_id") REFERENCES "public"."users"("id")
);

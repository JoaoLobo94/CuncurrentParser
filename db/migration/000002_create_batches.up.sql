CREATE TABLE "public"."batches" (
    "id" integer GENERATED ALWAYS AS IDENTITY,
    "dispatches" boolean,
    "amount" integer,
    "user_id" bigint,
    PRIMARY KEY ("id"),
    CONSTRAINT "user_id" FOREIGN KEY ("user_id") REFERENCES "public"."users"("id")
);

CREATE TABLE "public"."batches" (
    "id" integer GENERATED ALWAYS AS IDENTITY,
    "dispatched" boolean,
    "amount" float NOT NULL DEFAULT '0.0',
    "user_id" bigint, 
    "created_at" timestamp without time zone NOT NULL,
    PRIMARY KEY ("id"),
    CONSTRAINT "user_id" FOREIGN KEY ("user_id") REFERENCES "public"."users"("id")
);

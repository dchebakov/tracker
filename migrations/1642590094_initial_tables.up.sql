CREATE TABLE "customer" (
  "id" serial NOT NULL,
  "name" varchar(255) NOT NULL,
  "active" boolean NOT NULL DEFAULT TRUE,
  PRIMARY KEY ("id")
);

CREATE TABLE "ip_blacklist" (
  "ip" bigint NOT NULL,
  PRIMARY KEY ("ip")
);

CREATE TABLE "ua_blacklist" (
  "ua" varchar(255) NOT NULL,
  PRIMARY KEY ("ua")
);

CREATE TABLE "hourly_stats" (
  "id" serial NOT NULL,
  "customer_id" bigint NOT NULL,
  "time" timestamp NOT NULL,
  "request_count" bigint NOT NULL DEFAULT '0',
  "invalid_count" bigint NOT NULL DEFAULT '0',
  PRIMARY KEY ("id"),
  UNIQUE ("customer_id", "time"),
  CONSTRAINT "hourly_stats_customer_id" FOREIGN KEY ("customer_id") REFERENCES "customer" ("id") ON DELETE CASCADE ON UPDATE NO ACTION
);

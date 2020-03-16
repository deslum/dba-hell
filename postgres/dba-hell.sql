CREATE SCHEMA IF NOT EXISTS dba_test;

CREATE TABLE IF NOT EXISTS "dba_test.threads" (
	"id" BIGINT NOT NULL UNIQUE,
	"name" TEXT NOT NULL,
	"number" INTEGER NOT NULL,
	"body" TEXT NOT NULL,
	"timestamp" TIMESTAMP NOT NULL
);

	CREATE INDEX id_index ON "dba_test.threads" (id);
	CREATE INDEX name_index ON "dba_test.threads" (name);
	CREATE INDEX timestamp_index ON "dba_test.threads" (timestamp);
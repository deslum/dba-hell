CREATE SCHEMA IF NOT EXISTS dba_test;

CREATE TABLE IF NOT EXISTS "dba_test.threads" (
	"id" bigint NOT NULL UNIQUE,
	"body" TEXT(50) NOT NULL,
	"timestamp" TIMESTAMP NOT NULL,

	PRIMARY KEY(id, body, timestamp)
);
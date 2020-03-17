CREATE SCHEMA IF NOT EXISTS dba_test;

CREATE TABLE IF NOT EXISTS dba_test.threads (
	id bigint,
	name varchar(50) NOT NULL,
	number integer NOT NULL,
	body varchar(50) NOT NULL,
	ts timestamp NOT NULL,

	PRIMARY KEY(id)
);

	CREATE INDEX name_index ON dba_test.threads (name);
	CREATE INDEX timestamp_index ON dba_test.threads (ts);
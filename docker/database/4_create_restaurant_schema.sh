#!/bin/bash
set -e

psql -v ON_ERROR_STOP=1 --username "$POSTGRES_USER" --dbname "ftgogo" <<-EOSQL
    CREATE SCHEMA consumer;

    CREATE TABLE consumer.consumers(
        id          text NOT NULL,
        name        text NOT NULL,
        addresses   bytea NOT NULL,
        created_at  timestamptz NOT NULL DEFAULT NOW(),
        updated_at  timestamptz NOT NULL DEFAULT NOW(),
        PRIMARY KEY (id)
    );

    CREATE TRIGGER created_at_consumer_trgr BEFORE UPDATE ON consumer.consumers FOR EACH ROW EXECUTE PROCEDURE created_at_trigger();
    CREATE TRIGGER updated_at_consumer_trgr BEFORE UPDATE ON consumer.consumers FOR EACH ROW EXECUTE PROCEDURE updated_at_trigger();

    GRANT USAGE ON SCHEMA consumer TO ftgogo_user;
    GRANT INSERT, UPDATE, DELETE, SELECT ON ALL TABLES IN SCHEMA consumer TO ftgogo_user;
EOSQL

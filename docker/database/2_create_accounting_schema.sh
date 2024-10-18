#!/bin/bash
set -e

psql -v ON_ERROR_STOP=1 --username "$POSTGRES_USER" --dbname "ftgogo" <<-EOSQL
    CERATE SCHEMA accounting;

    CREATE TABLE accounting.accounts(
        id          text NOT NULL,
        name        text NOT NULL,
        enabled     bool NOT NULL,
        created_at  timestamptz NOT NULL DEFAULT NOW(),
        updated_at  timestamptz NOT NULL DEFAULT NOW(),
        PRIMARY KEY (id)
    );

    CREATE TRIGGER created_at_accounting_trgr BEFORE UPDATE ON accounting.accounts FOR EACH ROW EXECUTE PROCEDURE created_at_trigger();
    CREATE TRIGGER updated_at_accounting_trgr BEFORE UPDATE ON accounting.accounts FOR EACH ROW EXECUTE PROCEDURE updated_at_trigger();

    GRANT USAGE ON SCHEMA accounting TO ftgogo_user;
    GRANT INSERT, UPDATE, DELETE, SELECT ON ALL TABLES IN SCHEMA accounting TO ftgogo_user;
EOSQL

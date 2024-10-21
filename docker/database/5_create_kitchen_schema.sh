#!/bin/bash
set -e

psql -v ON_ERROR_STOP=1 --username "$POSTGRES_USER" --dbname "ftgogo" <<-EOSQL
    CREATE SCHEMA kitchen;

    CREATE TABLE kitchen.tickets(
        id                   text NOT NULL,
        restaurant_id        text NOT NULL,
        line_items           bytea NOT NULL,
        accepted_at          timestamptz NOT NULL,
        preparing_time       timestamptz NOT NULL,
        ready_for_pick_up_at timestamptz NOT NULL,
        picked_up_at         timestamptz NOT NULL,
        status               text NOT NULL,
        pervious_status      text NOT NULL,
        created_at           timestamptz NOT NULL DEFAULT NOW(),
        updated_at           timestamptz NOT NULL DEFAULT NOW(),
        PRIMARY KEY (id)
    );

    CREATE TRIGGER created_at_tickets_trgr BEFORE UPDATE ON kitchen.tickets FOR EACH ROW EXECUTE PROCEDURE created_at_trigger();
    CREATE TRIGGER updated_at_tickets_trgr BEFORE UPDATE ON kitchen.tickets FOR EACH ROW EXECUTE PROCEDURE updated_at_trigger();

    GRANT USAGE ON SCHEMA kitchen TO ftgogo_user;
    GRANT INSERT, UPDATE, DELETE, SELECT ON ALL TABLES IN SCHEMA kitchen TO ftgogo_user;
EOSQL

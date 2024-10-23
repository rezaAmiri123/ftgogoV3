#!/bin/bash
set -e

psql -v ON_ERROR_STOP=1 --username "$POSTGRES_USER" --dbname "ftgogo" <<-EOSQL
    CREATE SCHEMA orders;

    CREATE TABLE orders.orders(
        id             text NOT NULL,
        consumer_id    text NOT NULL,
        restaurant_id  text NOT NULL,
        ticket_id      text NOT NULL,
        line_items     bytea NOT NULL,
        status         text NOT NULL,
        deliver_at     timestamptz NOT NULL,
        deliver_to     bytea NOT NULL,
        created_at     timestamptz NOT NULL DEFAULT NOW(),
        updated_at     timestamptz NOT NULL DEFAULT NOW(),
        PRIMARY KEY (id)
    );

    CREATE TRIGGER created_at_orders_trgr BEFORE UPDATE ON orders.orders FOR EACH ROW EXECUTE PROCEDURE created_at_trigger();
    CREATE TRIGGER updated_at_orders_trgr BEFORE UPDATE ON orders.orders FOR EACH ROW EXECUTE PROCEDURE updated_at_trigger();

    GRANT USAGE ON SCHEMA orders TO ftgogo_user;
    GRANT INSERT, UPDATE, DELETE, SELECT ON ALL TABLES IN SCHEMA orders TO ftgogo_user;
EOSQL

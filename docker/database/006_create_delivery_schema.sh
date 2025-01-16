#!/bin/bash
set -e

psql -v ON_ERROR_STOP=1 --username "$POSTGRES_USER" --dbname "ftgogo" <<-EOSQL
    CREATE SCHEMA delivery;

    CREATE TABLE delivery.deliveries(
        id                   text NOT NULL,
        restaurant_id        text NOT NULL,
        assigned_courier_id  text NOT NULL,
        pick_up_address      bytea NOT NULL,
        delivery_address     bytea NOT NULL,
        pick_up_time         timestamptz NOT NULL,
        ready_by             timestamptz NOT NULL,
        status               text NOT NULL,
        created_at           timestamptz NOT NULL DEFAULT NOW(),
        updated_at           timestamptz NOT NULL DEFAULT NOW(),
        PRIMARY KEY (id)
    );

    CREATE TRIGGER created_at_deliveries_trgr BEFORE UPDATE ON delivery.deliveries FOR EACH ROW EXECUTE PROCEDURE created_at_trigger();
    CREATE TRIGGER updated_at_deliveries_trgr BEFORE UPDATE ON delivery.deliveries FOR EACH ROW EXECUTE PROCEDURE updated_at_trigger();

    CREATE TABLE delivery.couriers(
        id                   text NOT NULL,
        plan                 bytea NOT NULL,
        available            bool NOT NULL,
        created_at           timestamptz NOT NULL DEFAULT NOW(),
        updated_at           timestamptz NOT NULL DEFAULT NOW(),
        PRIMARY KEY (id)
    );

    CREATE TRIGGER created_at_couriers_trgr BEFORE UPDATE ON delivery.couriers FOR EACH ROW EXECUTE PROCEDURE created_at_trigger();
    CREATE TRIGGER updated_at_couriers_trgr BEFORE UPDATE ON delivery.couriers FOR EACH ROW EXECUTE PROCEDURE updated_at_trigger();

    GRANT USAGE ON SCHEMA delivery TO ftgogo_user;
    GRANT INSERT, UPDATE, DELETE, SELECT ON ALL TABLES IN SCHEMA delivery TO ftgogo_user;
EOSQL

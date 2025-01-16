#!/bin/bash
set -e

psql -v ON_ERROR_STOP=1 --username "$POSTGRES_USER" --dbname "ftgogo" <<-EOSQL
    CREATE SCHEMA restaurant;

    CREATE TABLE restaurant.restaurants(
        id          text NOT NULL,
        name        text NOT NULL,
        address     bytea NOT NULL,
        menu_items  bytea NOT NULL,
        created_at  timestamptz NOT NULL DEFAULT NOW(),
        updated_at  timestamptz NOT NULL DEFAULT NOW(),
        PRIMARY KEY (id)
    );

    CREATE TRIGGER created_at_restaurant_trgr BEFORE UPDATE ON restaurant.restaurants FOR EACH ROW EXECUTE PROCEDURE created_at_trigger();
    CREATE TRIGGER updated_at_restaurant_trgr BEFORE UPDATE ON restaurant.restaurants FOR EACH ROW EXECUTE PROCEDURE updated_at_trigger();

    GRANT USAGE ON SCHEMA restaurant TO ftgogo_user;
    GRANT INSERT, UPDATE, DELETE, SELECT ON ALL TABLES IN SCHEMA restaurant TO ftgogo_user;
EOSQL

-- +goose Up
CREATE SCHEMA delivery;
SET SEARCH_PATH TO delivery, public;

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

CREATE TABLE delivery.inbox
(
  id          text NOT NULL,
  name        text NOT NULL,
  subject     text NOT NULL,
  data        bytea NOT NULL,
  received_at timestamptz NOT NULL,
  PRIMARY KEY (id)
);

CREATE TABLE delivery.outbox
(
  id           text NOT NULL,
  name         text NOT NULL,
  subject      text NOT NULL,
  data         bytea NOT NULL,
  published_at timestamptz,
  PRIMARY KEY (id)
);

CREATE INDEX delivery_unpublished_idx ON delivery.outbox (published_at) WHERE published_at IS NULL;

GRANT USAGE ON SCHEMA delivery TO ftgogo_user;
GRANT INSERT, UPDATE, DELETE, SELECT ON ALL TABLES IN SCHEMA delivery TO ftgogo_user;

-- +goose Down
DROP SCHEMA IF EXISTS delivery CASCADE;
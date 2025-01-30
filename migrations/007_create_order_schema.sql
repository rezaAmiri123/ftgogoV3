-- +goose Up
CREATE SCHEMA orders;
SET SEARCH_PATH TO orders, public;

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

CREATE TABLE orders.events
(
    stream_id      text        NOT NULL,
    stream_name    text        NOT NULL,
    stream_version int         NOT NULL,
    event_id       text        NOT NULL,
    event_name     text        NOT NULL,
    event_data     bytea       NOT NULL,
    occurred_at    timestamptz NOT NULL DEFAULT CURRENT_TIMESTAMP,
    PRIMARY KEY (stream_id, stream_name, stream_version)
);

CREATE TABLE orders.snapshots
(
    stream_id        text        NOT NULL,
    stream_name      text        NOT NULL,
    stream_version   int         NOT NULL,
    snapshot_name    text        NOT NULL,
    snapshot_data    bytea       NOT NULL,
    updated_at       timestamptz NOT NULL DEFAULT CURRENT_TIMESTAMP,
    PRIMARY KEY (stream_id, stream_name)
);

CREATE TRIGGER updated_at_snapshots_trgr BEFORE UPDATE ON orders.snapshots FOR EACH ROW EXECUTE PROCEDURE updated_at_trigger();

CREATE TABLE orders.inbox
(
  id          text NOT NULL,
  name        text NOT NULL,
  subject     text NOT NULL,
  data        bytea NOT NULL,
  received_at timestamptz NOT NULL,
  PRIMARY KEY (id)
);

CREATE TABLE orders.outbox
(
  id           text NOT NULL,
  name         text NOT NULL,
  subject      text NOT NULL,
  data         bytea NOT NULL,
  published_at timestamptz,
  PRIMARY KEY (id)
);

CREATE INDEX orders_unpublished_idx ON orders.outbox (published_at) WHERE published_at IS NULL;

GRANT USAGE ON SCHEMA orders TO ftgogo_user;
GRANT INSERT, UPDATE, DELETE, SELECT ON ALL TABLES IN SCHEMA orders TO ftgogo_user;

-- +goose Down
DROP SCHEMA IF EXISTS orders CASCADE;
-- +goose Up
CREATE SCHEMA consumer;
SET SEARCH_PATH TO consumer, public;

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

CREATE TABLE consumer.inbox
(
  id          text NOT NULL,
  name        text NOT NULL,
  subject     text NOT NULL,
  data        bytea NOT NULL,
  received_at timestamptz NOT NULL,
  PRIMARY KEY (id)
);

CREATE TABLE consumer.outbox
(
  id           text NOT NULL,
  name         text NOT NULL,
  subject      text NOT NULL,
  data         bytea NOT NULL,
  published_at timestamptz,
  PRIMARY KEY (id)
);

CREATE INDEX consumer_unpublished_idx ON consumer.outbox (published_at) WHERE published_at IS NULL;

GRANT USAGE ON SCHEMA consumer TO ftgogo_user;
GRANT INSERT, UPDATE, DELETE, SELECT ON ALL TABLES IN SCHEMA consumer TO ftgogo_user;

-- +goose Down
DROP SCHEMA IF EXISTS consumer CASCADE;
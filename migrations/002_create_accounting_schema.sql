-- +goose Up
CREATE SCHEMA accounting;
SET SEARCH_PATH TO accounting, public;

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

CREATE TABLE accounting.inbox
(
  id          text NOT NULL,
  name        text NOT NULL,
  subject     text NOT NULL,
  data        bytea NOT NULL,
  received_at timestamptz NOT NULL,
  PRIMARY KEY (id)
);

CREATE TABLE accounting.outbox
(
  id           text NOT NULL,
  name         text NOT NULL,
  subject      text NOT NULL,
  data         bytea NOT NULL,
  published_at timestamptz,
  PRIMARY KEY (id)
);

CREATE INDEX accounting_unpublished_idx ON accounting.outbox (published_at) WHERE published_at IS NULL;

GRANT USAGE ON SCHEMA accounting TO ftgogo_user;
GRANT INSERT, UPDATE, DELETE, SELECT ON ALL TABLES IN SCHEMA accounting TO ftgogo_user;

-- +goose Down
DROP SCHEMA IF EXISTS accounting CASCADE;

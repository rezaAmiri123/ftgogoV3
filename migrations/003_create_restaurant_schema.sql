-- +goose Up
CREATE SCHEMA restaurant;
SET SEARCH_PATH TO restaurant, public;

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

CREATE TABLE restaurant.inbox
(
  id          text NOT NULL,
  name        text NOT NULL,
  subject     text NOT NULL,
  data        bytea NOT NULL,
  received_at timestamptz NOT NULL,
  PRIMARY KEY (id)
);

CREATE TABLE restaurant.outbox
(
  id           text NOT NULL,
  name         text NOT NULL,
  subject      text NOT NULL,
  data         bytea NOT NULL,
  published_at timestamptz,
  PRIMARY KEY (id)
);

CREATE INDEX restaurant_unpublished_idx ON restaurant.outbox (published_at) WHERE published_at IS NULL;

GRANT USAGE ON SCHEMA restaurant TO ftgogo_user;
GRANT INSERT, UPDATE, DELETE, SELECT ON ALL TABLES IN SCHEMA restaurant TO ftgogo_user;

-- +goose Down
DROP SCHEMA IF EXISTS restaurant CASCADE;
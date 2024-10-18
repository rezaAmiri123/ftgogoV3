#!/bin/bash
set -e

psql -v ON_ERROR_STOP=1 --username "$POSTGRES_USER" --dbname "mallbots" <<-EOSQL
    INSERT INTO accounting.accounts (id, name, enabled)
        VALUES ('1', 'test_name', true);
EOSQL

#!/bin/bash
set -e

# Run the SQL commands to create a table and load data from the txt file
psql -v ON_ERROR_STOP=1 --username "$POSTGRES_USER" --dbname "$POSTGRES_DB" <<-EOSQL
    CREATE TABLE IF NOT EXISTS my_table (
        id SERIAL PRIMARY KEY,
        word TEXT
    );

    \COPY my_table(word) FROM '/docker-entrypoint-initdb.d/data.txt' CSV;
EOSQL
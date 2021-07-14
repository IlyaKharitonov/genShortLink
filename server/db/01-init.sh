#!/bin/bash
set -e
export PGPASSWORD=$POSTGRES_PASSWORD;
psql -v ON_ERROR_STOP=1 --username "$POSTGRES_USER" --dbname "$POSTGRES_DB" <<-EOSQL
  CREATE USER $APP_DB_USER WITH PASSWORD '$APP_DB_PASS';
  CREATE DATABASE $APP_DB_NAME;
  GRANT ALL PRIVILEGES ON DATABASE $APP_DB_NAME TO $APP_DB_USER;
  \connect $APP_DB_NAME $APP_DB_USER
  BEGIN;
    create table links
    (
        id integer generated by default as identity
        constraint links_pkey primary key,
        longlink  varchar(250) not null,
        shortlink varchar(30)  not null
    );
	COMMIT;
EOSQL
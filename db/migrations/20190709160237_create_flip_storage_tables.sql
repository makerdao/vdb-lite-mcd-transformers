-- +goose Up
-- SQL in this section is executed when the migration is applied.
CREATE TABLE maker.flip_bid_bid
(
    id               SERIAL PRIMARY KEY,
    block_number     BIGINT,
    block_hash       TEXT,
    contract_address TEXT,
    bid_id           NUMERIC NOT NULL,
    bid              NUMERIC NOT NULL,
    UNIQUE (block_number, block_hash, bid_id, contract_address, bid)
);

CREATE INDEX flip_bid_bid_block_number_index
    ON maker.flip_bid_bid (block_number);
CREATE INDEX flip_bid_bid_bid_id_index
    ON maker.flip_bid_bid (bid_id);
CREATE INDEX flip_bid_bid_contract_address_index
    ON maker.flip_bid_bid (contract_address);

CREATE TABLE maker.flip_bid_lot
(
    id               SERIAL PRIMARY KEY,
    block_number     BIGINT,
    block_hash       TEXT,
    contract_address TEXT,
    bid_id           NUMERIC NOT NULL,
    lot              NUMERIC NOT NULL,
    UNIQUE (block_number, block_hash, bid_id, contract_address, lot)
);

CREATE INDEX flip_bid_lot_block_number_index
    ON maker.flip_bid_lot (block_number);
CREATE INDEX flip_bid_lot_bid_id_index
    ON maker.flip_bid_lot (bid_id);
CREATE INDEX flip_bid_lot_contract_address_index
    ON maker.flip_bid_lot (contract_address);

CREATE TABLE maker.flip_bid_guy
(
    id               SERIAL PRIMARY KEY,
    block_number     BIGINT,
    block_hash       TEXT,
    contract_address TEXT,
    bid_id           NUMERIC NOT NULL,
    guy              TEXT,
    UNIQUE (block_number, block_hash, bid_id, contract_address, guy)
);

CREATE INDEX flip_bid_guy_block_number_index
    ON maker.flip_bid_guy (block_number);
CREATE INDEX flip_bid_guy_bid_id_index
    ON maker.flip_bid_guy (bid_id);
CREATE INDEX flip_bid_guy_contract_address_index
    ON maker.flip_bid_guy (contract_address);

CREATE TABLE maker.flip_bid_tic
(
    id               SERIAL PRIMARY KEY,
    block_number     BIGINT,
    block_hash       TEXT,
    contract_address TEXT,
    bid_id           NUMERIC NOT NULL,
    tic              BIGINT  NOT NULL,
    UNIQUE (block_number, block_hash, bid_id, contract_address, tic)
);

CREATE INDEX flip_bid_tic_block_number_index
    ON maker.flip_bid_tic (block_number);
CREATE INDEX flip_bid_tic_bid_id_index
    ON maker.flip_bid_tic (bid_id);
CREATE INDEX flip_bid_tic_contract_address_index
    ON maker.flip_bid_tic (contract_address);

CREATE TABLE maker.flip_bid_end
(
    id               SERIAL PRIMARY KEY,
    block_number     BIGINT,
    block_hash       TEXT,
    contract_address TEXT,
    bid_id           NUMERIC NOT NULL,
    "end"            BIGINT  NOT NULL,
    UNIQUE (block_number, block_hash, bid_id, contract_address, "end")
);

CREATE INDEX flip_bid_end_block_number_index
    ON maker.flip_bid_end (block_number);
CREATE INDEX flip_bid_end_bid_id_index
    ON maker.flip_bid_end (bid_id);
CREATE INDEX flip_bid_end_contract_address_index
    ON maker.flip_bid_end (contract_address);

CREATE TABLE maker.flip_bid_usr
(
    id               SERIAL PRIMARY KEY,
    block_number     BIGINT,
    block_hash       TEXT,
    contract_address TEXT,
    bid_id           NUMERIC NOT NULL,
    usr              TEXT,
    UNIQUE (block_number, block_hash, bid_id, contract_address, usr)
);

CREATE INDEX flip_bid_usr_block_number_index
    ON maker.flip_bid_usr (block_number);
CREATE INDEX flip_bid_usr_bid_id_index
    ON maker.flip_bid_usr (bid_id);
CREATE INDEX flip_bid_usr_contract_address_index
    ON maker.flip_bid_usr (contract_address);

CREATE TABLE maker.flip_bid_gal
(
    id               SERIAL PRIMARY KEY,
    block_number     BIGINT,
    block_hash       TEXT,
    contract_address TEXT,
    bid_id           NUMERIC NOT NULL,
    gal              TEXT,
    UNIQUE (block_number, block_hash, bid_id, contract_address, gal)
);

CREATE INDEX flip_bid_gal_block_number_index
    ON maker.flip_bid_gal (block_number);
CREATE INDEX flip_bid_gal_bid_id_index
    ON maker.flip_bid_gal (bid_id);
CREATE INDEX flip_bid_gal_contract_address_index
    ON maker.flip_bid_gal (contract_address);

CREATE TABLE maker.flip_bid_tab
(
    id               SERIAL PRIMARY KEY,
    block_number     BIGINT,
    block_hash       TEXT,
    contract_address TEXT,
    bid_id           NUMERIC NOT NULL,
    tab              NUMERIC NOT NULL,
    UNIQUE (block_number, block_hash, bid_id, contract_address, tab)
);

CREATE INDEX flip_bid_tab_block_number_index
    ON maker.flip_bid_tab (block_number);
CREATE INDEX flip_bid_tab_bid_id_index
    ON maker.flip_bid_tab (bid_id);
CREATE INDEX flip_bid_tab_contract_address_index
    ON maker.flip_bid_tab (contract_address);

CREATE TABLE maker.flip_vat
(
    id               SERIAL PRIMARY KEY,
    block_number     BIGINT,
    block_hash       TEXT,
    contract_address TEXT,
    vat              TEXT,
    UNIQUE (block_number, block_hash, contract_address, vat)
);

CREATE TABLE maker.flip_ilk
(
    id               SERIAL PRIMARY KEY,
    block_number     BIGINT,
    block_hash       TEXT,
    contract_address TEXT,
    ilk_id           INTEGER NOT NULL REFERENCES maker.ilks (id) ON DELETE CASCADE,
    UNIQUE (block_number, block_hash, contract_address, ilk_id)
);

CREATE INDEX flip_ilk_ilk_id_index
    ON maker.flip_ilk (ilk_id);
CREATE INDEX flip_ilk_block_number_index
    ON maker.flip_ilk (block_number);

CREATE TABLE maker.flip_beg
(
    id               SERIAL PRIMARY KEY,
    block_number     BIGINT,
    block_hash       TEXT,
    contract_address TEXT,
    beg              NUMERIC NOT NULL,
    UNIQUE (block_number, block_hash, contract_address, beg)
);

CREATE TABLE maker.flip_ttl
(
    id               SERIAL PRIMARY KEY,
    block_number     BIGINT,
    block_hash       TEXT,
    contract_address TEXT,
    ttl              NUMERIC NOT NULL,
    UNIQUE (block_number, block_hash, contract_address, ttl)
);

CREATE TABLE maker.flip_tau
(
    id               SERIAL PRIMARY KEY,
    block_number     BIGINT,
    block_hash       TEXT,
    contract_address TEXT,
    tau              NUMERIC NOT NULL,
    UNIQUE (block_number, block_hash, contract_address, tau)
);

CREATE TABLE maker.flip_kicks
(
    id               SERIAL PRIMARY KEY,
    block_number     BIGINT,
    block_hash       TEXT,
    contract_address TEXT,
    kicks            NUMERIC NOT NULL,
    UNIQUE (block_number, block_hash, contract_address, kicks)
);

CREATE INDEX flip_kicks_block_number_index
    ON maker.flip_kicks (block_number);
CREATE INDEX flip_kicks_kicks_index
    ON maker.flip_kicks (kicks);
CREATE INDEX flip_kicks_contract_address_index
    ON maker.flip_kicks (contract_address);

-- prevent naming conflict with maker.flip_kick in postgraphile
COMMENT ON TABLE maker.flip_kicks IS E'@name flipKicksStorage';

-- +goose Down
-- SQL in this section is executed when the migration is rolled back.
DROP INDEX maker.flip_kicks_contract_address_index;
DROP INDEX maker.flip_kicks_kicks_index;
DROP INDEX maker.flip_kicks_block_number_index;
DROP INDEX maker.flip_ilk_block_number_index;
DROP INDEX maker.flip_ilk_ilk_id_index;
DROP INDEX maker.flip_bid_tab_contract_address_index;
DROP INDEX maker.flip_bid_tab_bid_id_index;
DROP INDEX maker.flip_bid_tab_block_number_index;
DROP INDEX maker.flip_bid_gal_contract_address_index;
DROP INDEX maker.flip_bid_gal_bid_id_index;
DROP INDEX maker.flip_bid_gal_block_number_index;
DROP INDEX maker.flip_bid_usr_contract_address_index;
DROP INDEX maker.flip_bid_usr_bid_id_index;
DROP INDEX maker.flip_bid_usr_block_number_index;
DROP INDEX maker.flip_bid_end_contract_address_index;
DROP INDEX maker.flip_bid_end_bid_id_index;
DROP INDEX maker.flip_bid_end_block_number_index;
DROP INDEX maker.flip_bid_tic_contract_address_index;
DROP INDEX maker.flip_bid_tic_bid_id_index;
DROP INDEX maker.flip_bid_tic_block_number_index;
DROP INDEX maker.flip_bid_guy_contract_address_index;
DROP INDEX maker.flip_bid_guy_bid_id_index;
DROP INDEX maker.flip_bid_guy_block_number_index;
DROP INDEX maker.flip_bid_lot_contract_address_index;
DROP INDEX maker.flip_bid_lot_bid_id_index;
DROP INDEX maker.flip_bid_lot_block_number_index;
DROP INDEX maker.flip_bid_bid_contract_address_index;
DROP INDEX maker.flip_bid_bid_bid_id_index;
DROP INDEX maker.flip_bid_bid_block_number_index;

DROP TABLE maker.flip_kicks;
DROP TABLE maker.flip_tau;
DROP TABLE maker.flip_ttl;
DROP TABLE maker.flip_beg;
DROP TABLE maker.flip_ilk;
DROP TABLE maker.flip_vat;
DROP TABLE maker.flip_bid_tab;
DROP TABLE maker.flip_bid_gal;
DROP TABLE maker.flip_bid_usr;
DROP TABLE maker.flip_bid_end;
DROP TABLE maker.flip_bid_tic;
DROP TABLE maker.flip_bid_guy;
DROP TABLE maker.flip_bid_lot;
DROP TABLE maker.flip_bid_bid;

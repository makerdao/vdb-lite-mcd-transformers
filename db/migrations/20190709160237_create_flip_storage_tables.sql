-- +goose Up
-- SQL in this section is executed when the migration is applied.
CREATE TABLE maker.flip_bid_bid
(
    id         SERIAL PRIMARY KEY,
    diff_id    BIGINT  NOT NULL REFERENCES storage_diff (id) ON DELETE CASCADE,
    header_id  INTEGER NOT NULL REFERENCES headers (id) ON DELETE CASCADE,
    address_id INTEGER NOT NULL REFERENCES addresses (id) ON DELETE CASCADE,
    bid_id     NUMERIC NOT NULL,
    bid        NUMERIC NOT NULL,
    UNIQUE (diff_id, header_id, bid_id, address_id, bid)
);

CREATE INDEX flip_bid_bid_header_id_index
    ON maker.flip_bid_bid (header_id);
COMMENT ON TABLE maker.flip_bid_bid
    IS E'@omit';

CREATE INDEX flip_bid_bid_bid_id_index
    ON maker.flip_bid_bid (bid_id);
CREATE INDEX flip_bid_bid_address_index
    ON maker.flip_bid_bid (address_id);

CREATE TABLE maker.flip_bid_lot
(
    id         SERIAL PRIMARY KEY,
    diff_id    BIGINT  NOT NULL REFERENCES storage_diff (id) ON DELETE CASCADE,
    header_id  INTEGER NOT NULL REFERENCES headers (id) ON DELETE CASCADE,
    address_id INTEGER NOT NULL REFERENCES addresses (id) ON DELETE CASCADE,
    bid_id     NUMERIC NOT NULL,
    lot        NUMERIC NOT NULL,
    UNIQUE (diff_id, header_id, bid_id, address_id, lot)
);

CREATE INDEX flip_bid_lot_header_id_index
    ON maker.flip_bid_lot (header_id);
CREATE INDEX flip_bid_lot_bid_id_index
    ON maker.flip_bid_lot (bid_id);
CREATE INDEX flip_bid_lot_address_index
    ON maker.flip_bid_lot (address_id);
COMMENT ON TABLE maker.flip_bid_lot
    IS E'@omit';

CREATE TABLE maker.flip_bid_guy
(
    id         SERIAL PRIMARY KEY,
    diff_id    BIGINT  NOT NULL REFERENCES storage_diff (id) ON DELETE CASCADE,
    header_id  INTEGER NOT NULL REFERENCES headers (id) ON DELETE CASCADE,
    address_id INTEGER NOT NULL REFERENCES addresses (id) ON DELETE CASCADE,
    bid_id     NUMERIC NOT NULL,
    guy        TEXT,
    UNIQUE (diff_id, header_id, bid_id, address_id, guy)
);

CREATE INDEX flip_bid_guy_header_id_index
    ON maker.flip_bid_guy (header_id);
CREATE INDEX flip_bid_guy_bid_id_index
    ON maker.flip_bid_guy (bid_id);
CREATE INDEX flip_bid_guy_address_index
    ON maker.flip_bid_guy (address_id);
COMMENT ON TABLE maker.flip_bid_guy
    IS E'@omit';

CREATE TABLE maker.flip_bid_tic
(
    id         SERIAL PRIMARY KEY,
    diff_id    BIGINT  NOT NULL REFERENCES storage_diff (id) ON DELETE CASCADE,
    header_id  INTEGER NOT NULL REFERENCES headers (id) ON DELETE CASCADE,
    address_id INTEGER NOT NULL REFERENCES addresses (id) ON DELETE CASCADE,
    bid_id     NUMERIC NOT NULL,
    tic        BIGINT  NOT NULL,
    UNIQUE (diff_id, header_id, bid_id, address_id, tic)
);

CREATE INDEX flip_bid_tic_header_id_index
    ON maker.flip_bid_tic (header_id);
CREATE INDEX flip_bid_tic_bid_id_index
    ON maker.flip_bid_tic (bid_id);
CREATE INDEX flip_bid_tic_address_index
    ON maker.flip_bid_tic (address_id);
COMMENT ON TABLE maker.flip_bid_tic
    IS E'@omit';

CREATE TABLE maker.flip_bid_end
(
    id         SERIAL PRIMARY KEY,
    diff_id    BIGINT  NOT NULL REFERENCES storage_diff (id) ON DELETE CASCADE,
    header_id  INTEGER NOT NULL REFERENCES headers (id) ON DELETE CASCADE,
    address_id INTEGER NOT NULL REFERENCES addresses (id) ON DELETE CASCADE,
    bid_id     NUMERIC NOT NULL,
    "end"      BIGINT  NOT NULL,
    UNIQUE (diff_id, header_id, bid_id, address_id, "end")
);

CREATE INDEX flip_bid_end_header_id_index
    ON maker.flip_bid_end (header_id);
CREATE INDEX flip_bid_end_bid_id_index
    ON maker.flip_bid_end (bid_id);
CREATE INDEX flip_bid_end_address_index
    ON maker.flip_bid_end (address_id);
COMMENT ON TABLE maker.flip_bid_end
    IS E'@omit';

CREATE TABLE maker.flip_bid_usr
(
    id         SERIAL PRIMARY KEY,
    diff_id    BIGINT  NOT NULL REFERENCES storage_diff (id) ON DELETE CASCADE,
    header_id  INTEGER NOT NULL REFERENCES headers (id) ON DELETE CASCADE,
    address_id INTEGER NOT NULL REFERENCES addresses (id) ON DELETE CASCADE,
    bid_id     NUMERIC NOT NULL,
    usr        TEXT,
    UNIQUE (diff_id, header_id, bid_id, address_id, usr)
);

CREATE INDEX flip_bid_usr_header_id_index
    ON maker.flip_bid_usr (header_id);
CREATE INDEX flip_bid_usr_bid_id_index
    ON maker.flip_bid_usr (bid_id);
CREATE INDEX flip_bid_usr_address_index
    ON maker.flip_bid_usr (address_id);
COMMENT ON TABLE maker.flip_bid_usr
    IS E'@omit';

CREATE TABLE maker.flip_bid_gal
(
    id         SERIAL PRIMARY KEY,
    diff_id    BIGINT  NOT NULL REFERENCES storage_diff (id) ON DELETE CASCADE,
    header_id  INTEGER NOT NULL REFERENCES headers (id) ON DELETE CASCADE,
    address_id INTEGER NOT NULL REFERENCES addresses (id) ON DELETE CASCADE,
    bid_id     NUMERIC NOT NULL,
    gal        TEXT,
    UNIQUE (diff_id, header_id, bid_id, address_id, gal)
);

CREATE INDEX flip_bid_gal_header_id_index
    ON maker.flip_bid_gal (header_id);
CREATE INDEX flip_bid_gal_bid_id_index
    ON maker.flip_bid_gal (bid_id);
CREATE INDEX flip_bid_gal_address_index
    ON maker.flip_bid_gal (address_id);
COMMENT ON TABLE maker.flip_bid_gal
    IS E'@omit';

CREATE TABLE maker.flip_bid_tab
(
    id         SERIAL PRIMARY KEY,
    diff_id    BIGINT  NOT NULL REFERENCES storage_diff (id) ON DELETE CASCADE,
    header_id  INTEGER NOT NULL REFERENCES headers (id) ON DELETE CASCADE,
    address_id INTEGER NOT NULL REFERENCES addresses (id) ON DELETE CASCADE,
    bid_id     NUMERIC NOT NULL,
    tab        NUMERIC NOT NULL,
    UNIQUE (diff_id, header_id, bid_id, address_id, tab)
);

CREATE INDEX flip_bid_tab_header_id_index
    ON maker.flip_bid_tab (header_id);
CREATE INDEX flip_bid_tab_bid_id_index
    ON maker.flip_bid_tab (bid_id);
CREATE INDEX flip_bid_tab_address_index
    ON maker.flip_bid_tab (address_id);
COMMENT ON TABLE maker.flip_bid_tab
    IS E'@omit';

CREATE TABLE maker.flip_vat
(
    id         SERIAL PRIMARY KEY,
    diff_id    BIGINT  NOT NULL REFERENCES storage_diff (id) ON DELETE CASCADE,
    header_id  INTEGER NOT NULL REFERENCES headers (id) ON DELETE CASCADE,
    address_id INTEGER NOT NULL REFERENCES addresses (id) ON DELETE CASCADE,
    vat        TEXT,
    UNIQUE (diff_id, header_id, address_id, vat)
);

CREATE INDEX flip_vat_header_id_index
    ON maker.flip_vat (header_id);
CREATE INDEX flip_vat_address_index
    ON maker.flip_vat (address_id);
COMMENT ON TABLE maker.flip_vat
    IS E'@omit';

CREATE TABLE maker.flip_ilk
(
    id         SERIAL PRIMARY KEY,
    diff_id    BIGINT  NOT NULL REFERENCES storage_diff (id) ON DELETE CASCADE,
    header_id  INTEGER NOT NULL REFERENCES headers (id) ON DELETE CASCADE,
    address_id INTEGER NOT NULL REFERENCES addresses (id) ON DELETE CASCADE,
    ilk_id     INTEGER NOT NULL REFERENCES maker.ilks (id) ON DELETE CASCADE,
    UNIQUE (diff_id, header_id, address_id, ilk_id)
);

CREATE INDEX flip_ilk_header_id_index
    ON maker.flip_ilk (header_id);
CREATE INDEX flip_ilk_ilk_id_index
    ON maker.flip_ilk (ilk_id);
CREATE INDEX flip_ilk_address_index
    ON maker.flip_ilk (address_id);
COMMENT ON TABLE maker.flip_ilk
    IS E'@omit';

CREATE TABLE maker.flip_beg
(
    id         SERIAL PRIMARY KEY,
    diff_id    BIGINT  NOT NULL REFERENCES storage_diff (id) ON DELETE CASCADE,
    header_id  INTEGER NOT NULL REFERENCES headers (id) ON DELETE CASCADE,
    address_id INTEGER NOT NULL REFERENCES addresses (id) ON DELETE CASCADE,
    beg        NUMERIC NOT NULL,
    UNIQUE (diff_id, header_id, address_id, beg)
);

CREATE INDEX flip_beg_header_id_index
    ON maker.flip_beg (header_id);
CREATE INDEX flip_beg_address_index
    ON maker.flip_beg (address_id);
COMMENT ON TABLE maker.flip_beg
    IS E'@omit';

CREATE TABLE maker.flip_ttl
(
    id         SERIAL PRIMARY KEY,
    diff_id    BIGINT  NOT NULL REFERENCES storage_diff (id) ON DELETE CASCADE,
    header_id  INTEGER NOT NULL REFERENCES headers (id) ON DELETE CASCADE,
    address_id INTEGER NOT NULL REFERENCES addresses (id) ON DELETE CASCADE,
    ttl        NUMERIC NOT NULL,
    UNIQUE (diff_id, header_id, address_id, ttl)
);

CREATE INDEX flip_ttl_header_id_index
    ON maker.flip_ttl (header_id);
CREATE INDEX flip_ttl_address_index
    ON maker.flip_ttl (address_id);
COMMENT ON TABLE maker.flip_ttl
    IS E'@omit';

CREATE TABLE maker.flip_tau
(
    id         SERIAL PRIMARY KEY,
    diff_id    BIGINT  NOT NULL REFERENCES storage_diff (id) ON DELETE CASCADE,
    header_id  INTEGER NOT NULL REFERENCES headers (id) ON DELETE CASCADE,
    address_id INTEGER NOT NULL REFERENCES addresses (id) ON DELETE CASCADE,
    tau        NUMERIC NOT NULL,
    UNIQUE (diff_id, header_id, address_id, tau)
);

CREATE INDEX flip_tau_header_id_index
    ON maker.flip_tau (header_id);
CREATE INDEX flip_tau_address_index
    ON maker.flip_tau (address_id);
COMMENT ON TABLE maker.flip_tau
    IS E'@omit';

CREATE TABLE maker.flip_kicks
(
    id         SERIAL PRIMARY KEY,
    diff_id    BIGINT  NOT NULL REFERENCES storage_diff (id) ON DELETE CASCADE,
    header_id  INTEGER NOT NULL REFERENCES headers (id) ON DELETE CASCADE,
    address_id INTEGER NOT NULL REFERENCES addresses (id) ON DELETE CASCADE,
    kicks      NUMERIC NOT NULL,
    UNIQUE (diff_id, header_id, address_id, kicks)
);

CREATE INDEX flip_kicks_header_id_index
    ON maker.flip_kicks (header_id);
CREATE INDEX flip_kicks_address_index
    ON maker.flip_kicks (address_id);
-- prevent naming conflict with maker.flip_kick in postgraphile
COMMENT ON TABLE maker.flip_kicks
    IS E'@name flipKicksStorage\n@omit';


-- +goose Down
-- SQL in this section is executed when the migration is rolled back.
DROP INDEX maker.flip_kicks_address_index;
DROP INDEX maker.flip_kicks_header_id_index;
DROP INDEX maker.flip_tau_address_index;
DROP INDEX maker.flip_tau_header_id_index;
DROP INDEX maker.flip_ttl_address_index;
DROP INDEX maker.flip_ttl_header_id_index;
DROP INDEX maker.flip_beg_address_index;
DROP INDEX maker.flip_beg_header_id_index;
DROP INDEX maker.flip_ilk_address_index;
DROP INDEX maker.flip_ilk_ilk_id_index;
DROP INDEX maker.flip_ilk_header_id_index;
DROP INDEX maker.flip_vat_address_index;
DROP INDEX maker.flip_vat_header_id_index;
DROP INDEX maker.flip_bid_tab_address_index;
DROP INDEX maker.flip_bid_tab_bid_id_index;
DROP INDEX maker.flip_bid_tab_header_id_index;
DROP INDEX maker.flip_bid_gal_address_index;
DROP INDEX maker.flip_bid_gal_bid_id_index;
DROP INDEX maker.flip_bid_gal_header_id_index;
DROP INDEX maker.flip_bid_usr_address_index;
DROP INDEX maker.flip_bid_usr_bid_id_index;
DROP INDEX maker.flip_bid_usr_header_id_index;
DROP INDEX maker.flip_bid_end_address_index;
DROP INDEX maker.flip_bid_end_bid_id_index;
DROP INDEX maker.flip_bid_end_header_id_index;
DROP INDEX maker.flip_bid_tic_address_index;
DROP INDEX maker.flip_bid_tic_bid_id_index;
DROP INDEX maker.flip_bid_tic_header_id_index;
DROP INDEX maker.flip_bid_guy_address_index;
DROP INDEX maker.flip_bid_guy_bid_id_index;
DROP INDEX maker.flip_bid_guy_header_id_index;
DROP INDEX maker.flip_bid_lot_address_index;
DROP INDEX maker.flip_bid_lot_bid_id_index;
DROP INDEX maker.flip_bid_lot_header_id_index;
DROP INDEX maker.flip_bid_bid_address_index;
DROP INDEX maker.flip_bid_bid_bid_id_index;
DROP INDEX maker.flip_bid_bid_header_id_index;

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

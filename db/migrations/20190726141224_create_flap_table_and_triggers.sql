-- +goose Up
CREATE TABLE maker.flap
(
    id           SERIAL PRIMARY KEY,
    -- TODO: remove hash?
    -- Can we replace block number + hash with header_id?
    -- Tricky because there might be multiple diffs contributing to a single row
    block_number BIGINT  DEFAULT NULL,
    block_hash   TEXT    DEFAULT NULL,
    address_id   INTEGER NOT NULL REFERENCES addresses (id) ON DELETE CASCADE,
    bid_id       NUMERIC DEFAULT NULL,
    guy          TEXT    DEFAULT NULL,
    tic          BIGINT  DEFAULT NULL,
    "end"        BIGINT  DEFAULT NULL,
    lot          NUMERIC DEFAULT NULL,
    bid          NUMERIC DEFAULT NULL,
    created      TIMESTAMP,
    updated      TIMESTAMP,
    UNIQUE (block_number, bid_id)
);

CREATE INDEX flap_address_index
    ON maker.flap (address_id);

COMMENT ON TABLE maker.flap
    IS E'@omit';

CREATE FUNCTION get_latest_flap_bid_guy(bid_id numeric) RETURNS TEXT AS
$$
SELECT guy
FROM maker.flap
WHERE guy IS NOT NULL
  AND flap.bid_id = bid_id
ORDER BY block_number DESC
LIMIT 1
$$
    LANGUAGE sql;

COMMENT ON FUNCTION get_latest_flap_bid_guy
    IS E'@omit';

CREATE FUNCTION get_latest_flap_bid_bid(bid_id numeric) RETURNS NUMERIC AS
$$
SELECT bid
FROM maker.flap
WHERE bid IS NOT NULL
  AND flap.bid_id = bid_id
ORDER BY block_number DESC
LIMIT 1
$$
    LANGUAGE sql;

COMMENT ON FUNCTION get_latest_flap_bid_bid
    IS E'@omit';

CREATE FUNCTION get_latest_flap_bid_tic(bid_id numeric) RETURNS BIGINT AS
$$
SELECT tic
FROM maker.flap
WHERE tic IS NOT NULL
  AND flap.bid_id = bid_id
ORDER BY block_number DESC
LIMIT 1
$$
    LANGUAGE sql;

COMMENT ON FUNCTION get_latest_flap_bid_tic
    IS E'@omit';

CREATE FUNCTION get_latest_flap_bid_end(bid_id numeric) RETURNS BIGINT AS
$$
SELECT "end"
FROM maker.flap
WHERE "end" IS NOT NULL
  AND flap.bid_id = bid_id
ORDER BY block_number DESC
LIMIT 1
$$
    LANGUAGE sql;

COMMENT ON FUNCTION get_latest_flap_bid_end
    IS E'@omit';

CREATE FUNCTION get_latest_flap_bid_lot(bid_id numeric) RETURNS NUMERIC AS
$$
SELECT lot
FROM maker.flap
WHERE lot IS NOT NULL
  AND flap.bid_id = bid_id
ORDER BY block_number DESC
LIMIT 1
$$
    LANGUAGE sql;

COMMENT ON FUNCTION get_latest_flap_bid_lot
    IS E'@omit';

-- +goose StatementBegin
CREATE OR REPLACE FUNCTION maker.insert_updated_flap_bid() RETURNS TRIGGER
AS
$$
BEGIN
    WITH created AS (
        SELECT created
        FROM maker.flap
        WHERE flap.bid_id = NEW.bid_id
        ORDER BY flap.block_number
        LIMIT 1
    ),
         diff_block AS (
             SELECT block_number, hash, block_timestamp
             FROM public.headers
             WHERE id = NEW.header_id
         )
    INSERT
    INTO maker.flap(bid_id, address_id, block_number, block_hash, bid, guy, tic, "end", lot, updated,
                    created)
    VALUES (NEW.bid_id, NEW.address_id, (SELECT block_number FROM diff_block), (SELECT hash FROM diff_block), NEW.bid,
            (SELECT get_latest_flap_bid_guy(NEW.bid_id)),
            (SELECT get_latest_flap_bid_tic(NEW.bid_id)),
            (SELECT get_latest_flap_bid_end(NEW.bid_id)),
            (SELECT get_latest_flap_bid_lot(NEW.bid_id)),
            (SELECT api.epoch_to_datetime(block_timestamp) FROM diff_block),
            (SELECT created FROM created))
    ON CONFLICT (bid_id, block_number) DO UPDATE SET bid = NEW.bid;
    return NEW;
END
$$
    LANGUAGE plpgsql;
-- +goose StatementEnd

-- +goose StatementBegin
CREATE OR REPLACE FUNCTION maker.insert_updated_flap_guy() RETURNS TRIGGER
AS
$$
BEGIN
    WITH created AS (
        SELECT created
        FROM maker.flap
        WHERE flap.bid_id = NEW.bid_id
        ORDER BY flap.block_number
        LIMIT 1
    ),
         diff_block AS (
             SELECT block_number, hash, block_timestamp
             FROM public.headers
             WHERE id = NEW.header_id
         )
    INSERT
    INTO maker.flap(bid_id, address_id, block_number, block_hash, guy, bid, tic, "end", lot, updated,
                    created)
    VALUES (NEW.bid_id, NEW.address_id, (SELECT block_number FROM diff_block), (SELECT hash FROM diff_block), NEW.guy,
            (SELECT get_latest_flap_bid_bid(NEW.bid_id)),
            (SELECT get_latest_flap_bid_tic(NEW.bid_id)),
            (SELECT get_latest_flap_bid_end(NEW.bid_id)),
            (SELECT get_latest_flap_bid_lot(NEW.bid_id)),
            (SELECT api.epoch_to_datetime(block_timestamp) FROM diff_block),
            (SELECT created FROM created))
    ON CONFLICT (bid_id, block_number) DO UPDATE SET guy = NEW.guy;
    return NEW;
END
$$
    LANGUAGE plpgsql;
-- +goose StatementEnd

-- +goose StatementBegin
CREATE OR REPLACE FUNCTION maker.insert_updated_flap_tic() RETURNS TRIGGER
AS
$$
BEGIN
    WITH created AS (
        SELECT created
        FROM maker.flap
        WHERE flap.bid_id = NEW.bid_id
        ORDER BY flap.block_number
        LIMIT 1
    ),
         diff_block AS (
             SELECT block_number, hash, block_timestamp
             FROM public.headers
             WHERE id = NEW.header_id
         )
    INSERT
    INTO maker.flap(bid_id, address_id, block_number, block_hash, tic, bid, guy, "end", lot, updated,
                    created)
    VALUES (NEW.bid_id, NEW.address_id, (SELECT block_number FROM diff_block), (SELECT hash FROM diff_block), NEW.tic,
            (SELECT get_latest_flap_bid_bid(NEW.bid_id)),
            (SELECT get_latest_flap_bid_guy(NEW.bid_id)),
            (SELECT get_latest_flap_bid_end(NEW.bid_id)),
            (SELECT get_latest_flap_bid_lot(NEW.bid_id)),
            (SELECT api.epoch_to_datetime(block_timestamp) FROM diff_block),
            (SELECT created FROM created))
    ON CONFLICT (bid_id, block_number) DO UPDATE SET tic = NEW.tic;
    return NEW;
END
$$
    LANGUAGE plpgsql;
-- +goose StatementEnd

-- +goose StatementBegin
CREATE OR REPLACE FUNCTION maker.insert_updated_flap_end() RETURNS TRIGGER
AS
$$
BEGIN
    WITH created AS (
        SELECT created
        FROM maker.flap
        WHERE flap.bid_id = NEW.bid_id
        ORDER BY flap.block_number
        LIMIT 1
    ),
         diff_block AS (
             SELECT block_number, hash, block_timestamp
             FROM public.headers
             WHERE id = NEW.header_id
         )
    INSERT
    INTO maker.flap(bid_id, address_id, block_number, block_hash, "end", bid, guy, tic, lot, updated,
                    created)
    VALUES (NEW.bid_id, NEW.address_id, (SELECT block_number FROM diff_block), (SELECT hash FROM diff_block), NEW."end",
            (SELECT get_latest_flap_bid_bid(NEW.bid_id)),
            (SELECT get_latest_flap_bid_guy(NEW.bid_id)),
            (SELECT get_latest_flap_bid_tic(NEW.bid_id)),
            (SELECT get_latest_flap_bid_lot(NEW.bid_id)),
            (SELECT api.epoch_to_datetime(block_timestamp) FROM diff_block),
            (SELECT created FROM created))
    ON CONFLICT (bid_id, block_number) DO UPDATE SET "end" = NEW."end";
    return NEW;
END
$$
    LANGUAGE plpgsql;
-- +goose StatementEnd

-- +goose StatementBegin
CREATE OR REPLACE FUNCTION maker.insert_updated_flap_lot() RETURNS TRIGGER
AS
$$
BEGIN
    WITH created AS (
        SELECT created
        FROM maker.flap
        WHERE flap.bid_id = NEW.bid_id
        ORDER BY flap.block_number
        LIMIT 1
    ),
         diff_block AS (
             SELECT block_number, hash, block_timestamp
             FROM public.headers
             WHERE id = NEW.header_id
         )
    INSERT
    INTO maker.flap(bid_id, address_id, block_number, block_hash, lot, bid, guy, tic, "end", updated,
                    created)
    VALUES (NEW.bid_id, NEW.address_id, (SELECT block_number FROM diff_block), (SELECT hash FROM diff_block), NEW.lot,
            (SELECT get_latest_flap_bid_bid(NEW.bid_id)),
            (SELECT get_latest_flap_bid_guy(NEW.bid_id)),
            (SELECT get_latest_flap_bid_tic(NEW.bid_id)),
            (SELECT get_latest_flap_bid_end(NEW.bid_id)),
            (SELECT api.epoch_to_datetime(block_timestamp) FROM diff_block),
            (SELECT created FROM created))
    ON CONFLICT (bid_id, block_number) DO UPDATE SET lot = NEW.lot;
    return NEW;
END
$$
    LANGUAGE plpgsql;
-- +goose StatementEnd

-- +goose StatementBegin
CREATE OR REPLACE FUNCTION maker.flap_created() RETURNS TRIGGER
AS
$$
BEGIN
    WITH block_info AS (
        SELECT block_number, hash, api.epoch_to_datetime(headers.block_timestamp) AS datetime
        FROM public.headers
        WHERE headers.id = NEW.header_id
        LIMIT 1
    )
    INSERT
    INTO maker.flap(bid_id, address_id, block_number, block_hash, created, updated, bid, guy, tic, "end", lot)
    VALUES (NEW.bid_id, NEW.address_id,
            (SELECT block_number FROM block_info),
            (SELECT hash FROM block_info),
            (SELECT datetime FROM block_info),
            (SELECT datetime FROM block_info),
            (SELECT get_latest_flap_bid_bid(NEW.bid_id)),
            (SELECT get_latest_flap_bid_guy(NEW.bid_id)),
            (SELECT get_latest_flap_bid_tic(NEW.bid_id)),
            (SELECT get_latest_flap_bid_end(NEW.bid_id)),
            (SELECT get_latest_flap_bid_lot(NEW.bid_id)))
    ON CONFLICT (bid_id, block_number) DO UPDATE SET created = (SELECT datetime FROM block_info),
                                                     updated = (SELECT datetime FROM block_info);
    return NEW;
END
$$
    LANGUAGE plpgsql;
-- +goose StatementEnd

CREATE TRIGGER flap_bid_bid
    AFTER INSERT OR UPDATE
    ON maker.flap_bid_bid
    FOR EACH ROW
EXECUTE PROCEDURE maker.insert_updated_flap_bid();

CREATE TRIGGER flap_bid_guy
    AFTER INSERT OR UPDATE
    ON maker.flap_bid_guy
    FOR EACH ROW
EXECUTE PROCEDURE maker.insert_updated_flap_guy();

CREATE TRIGGER flap_bid_tic
    AFTER INSERT OR UPDATE
    ON maker.flap_bid_tic
    FOR EACH ROW
EXECUTE PROCEDURE maker.insert_updated_flap_tic();

CREATE TRIGGER flap_bid_end
    AFTER INSERT OR UPDATE
    ON maker.flap_bid_end
    FOR EACH ROW
EXECUTE PROCEDURE maker.insert_updated_flap_end();

CREATE TRIGGER flap_bid_lot
    AFTER INSERT OR UPDATE
    ON maker.flap_bid_lot
    FOR EACH ROW
EXECUTE PROCEDURE maker.insert_updated_flap_lot();

CREATE TRIGGER flap_created_trigger
    AFTER INSERT
    ON maker.flap_kick
    FOR EACH ROW
EXECUTE PROCEDURE maker.flap_created();

-- +goose Down
DROP TRIGGER flap_bid_bid ON maker.flap_bid_bid;
DROP TRIGGER flap_bid_guy ON maker.flap_bid_guy;
DROP TRIGGER flap_bid_tic ON maker.flap_bid_tic;
DROP TRIGGER flap_bid_end ON maker.flap_bid_end;
DROP TRIGGER flap_bid_lot ON maker.flap_bid_lot;
DROP TRIGGER flap_created_trigger ON maker.flap_kick;

DROP FUNCTION maker.insert_updated_flap_bid();
DROP FUNCTION maker.insert_updated_flap_guy();
DROP FUNCTION maker.insert_updated_flap_tic();
DROP FUNCTION maker.insert_updated_flap_end();
DROP FUNCTION maker.insert_updated_flap_lot();
DROP FUNCTION maker.flap_created();
DROP FUNCTION get_latest_flap_bid_guy(numeric);
DROP FUNCTION get_latest_flap_bid_bid(numeric);
DROP FUNCTION get_latest_flap_bid_tic(numeric);
DROP FUNCTION get_latest_flap_bid_end(numeric);
DROP FUNCTION get_latest_flap_bid_lot(numeric);

DROP INDEX maker.flap_address_index;
DROP TABLE maker.flap;
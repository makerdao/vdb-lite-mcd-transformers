-- +goose Up
CREATE TABLE maker.vat_grab
(
    id        SERIAL PRIMARY KEY,
    header_id INTEGER NOT NULL REFERENCES headers (id) ON DELETE CASCADE,
    log_id    BIGINT  NOT NULL REFERENCES header_sync_logs (id) ON DELETE CASCADE,
    urn_id    INTEGER NOT NULL REFERENCES maker.urns (id) ON DELETE CASCADE,
    v         TEXT,
    w         TEXT,
    dink      NUMERIC,
    dart      NUMERIC,
    UNIQUE (header_id, log_id)
);

CREATE INDEX vat_grab_header_index
    ON maker.vat_grab (header_id);
CREATE INDEX vat_grab_log_index
    ON maker.vat_grab (log_id);
CREATE INDEX vat_grab_urn_index
    ON maker.vat_grab (urn_id);


-- +goose Down
DROP INDEX maker.vat_grab_header_index;
DROP INDEX maker.vat_grab_log_index;
DROP INDEX maker.vat_grab_urn_index;

DROP TABLE maker.vat_grab;

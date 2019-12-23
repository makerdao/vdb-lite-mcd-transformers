-- +goose Up
CREATE TABLE public.full_sync_receipts
(
    id                  SERIAL PRIMARY KEY,
    transaction_id      INTEGER NOT NULL REFERENCES full_sync_transactions (id) ON DELETE CASCADE,
    contract_address_id INTEGER NOT NULL REFERENCES addresses (id) ON DELETE CASCADE,
    cumulative_gas_used NUMERIC,
    gas_used            NUMERIC,
    state_root          VARCHAR(66),
    status              INTEGER,
    tx_hash             VARCHAR(66)
);

CREATE INDEX full_sync_receipts_contract_address
    ON full_sync_receipts (contract_address_id);

COMMENT ON TABLE public.full_sync_receipts
    IS E'@omit';

-- +goose Down
DROP INDEX full_sync_receipts_contract_address;
DROP TABLE full_sync_receipts;

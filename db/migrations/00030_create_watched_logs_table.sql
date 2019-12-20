-- +goose Up
-- SQL in this section is executed when the migration is applied.
CREATE TABLE public.watched_logs
(
    id               SERIAL PRIMARY KEY,
    contract_address VARCHAR(42),
    topic_zero       VARCHAR(66)
);

COMMENT ON TABLE public.watched_logs
    IS E'@omit';

COMMENT ON TABLE public.goose_db_version
    IS E'@omit';

-- +goose Down
-- SQL in this section is executed when the migration is rolled back.
DROP TABLE public.watched_logs;

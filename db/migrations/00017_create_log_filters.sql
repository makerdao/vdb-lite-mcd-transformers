-- +goose Up
CREATE TABLE public.log_filters (
  id         SERIAL,
  name       VARCHAR NOT NULL CHECK (name <> ''),
  from_block BIGINT CHECK (from_block >= 0),
  to_block   BIGINT CHECK (from_block >= 0),
  address    VARCHAR(66),
  topic0     VARCHAR(66),
  topic1     VARCHAR(66),
  topic2     VARCHAR(66),
  topic3     VARCHAR(66),
  CONSTRAINT name_uc UNIQUE (name)
);

COMMENT ON TABLE public.log_filters
    IS E'@omit';

-- +goose Down
DROP TABLE log_filters;

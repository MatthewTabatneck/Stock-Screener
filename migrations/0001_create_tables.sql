-- +goose Up
CREATE TABLE tickers (
  symbol TEXT PRIMARY KEY,
  updated_at TIMESTAMPTZ NOT NULL DEFAULT now(),
  is_processed BOOLEAN NOT NULL DEFAULT true
);

CREATE INDEX IF NOT EXISTS idx_tickers_processed
    ON tickers (is_processed);


CREATE TABLE snapshots (
  id BIGSERIAL PRIMARY KEY,
  symbol TEXT REFERENCES tickers(symbol) ON DELETE CASCADE,
  fetched_at TIMESTAMPTZ NOT NULL,
  price NUMERIC(18,6),
  change_pct NUMERIC(9,4),
  volume BIGINT,
  source TEXT NOT NULL,
  status TEXT CHECK (status IN ('OK','TIMEOUT','ERROR'))
);

-- +goose Down
DROP TABLE snapshots;
DROP TABLE tickers;

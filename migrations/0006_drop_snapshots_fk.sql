-- +goose Up
ALTER TABLE snapshots DROP CONSTRAINT IF EXISTS snapshots_symbol_fkey;

-- +goose Down
ALTER TABLE snapshots
  ADD CONSTRAINT snapshots_symbol_fkey
  FOREIGN KEY (symbol) REFERENCES tickers(symbol) ON DELETE RESTRICT;

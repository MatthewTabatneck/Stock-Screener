-- +goose Up
ALTER TABLE snapshots DROP CONSTRAINT IF EXISTS snapshots_symbol_fkey;
ALTER TABLE snapshots
  ADD CONSTRAINT snapshots_symbol_fkey
  FOREIGN KEY (symbol) REFERENCES tickers(symbol) ON DELETE RESTRICT;

-- +goose Down
ALTER TABLE snapshots DROP CONSTRAINT IF EXISTS snapshots_symbol_fkey;
ALTER TABLE snapshots
  ADD CONSTRAINT snapshots_symbol_fkey
  FOREIGN KEY (symbol) REFERENCES tickers(symbol) ON DELETE CASCADE;

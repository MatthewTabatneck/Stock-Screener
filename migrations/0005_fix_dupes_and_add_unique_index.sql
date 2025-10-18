-- +goose Up
-- remove duplicate rows by keeping the newest (by fetched_at) per (symbol, latest_trading_day, source)
WITH ranked AS (
  SELECT
    ctid,
    ROW_NUMBER() OVER (
      PARTITION BY symbol, latest_trading_day, source
      ORDER BY fetched_at DESC, ctid DESC
    ) AS rn
  FROM snapshots
  WHERE latest_trading_day IS NOT NULL
)
DELETE FROM snapshots s
USING ranked r
WHERE s.ctid = r.ctid
  AND r.rn > 1;

-- now create the unique index
CREATE UNIQUE INDEX IF NOT EXISTS ux_snapshots_symbol_day_source
  ON snapshots(symbol, latest_trading_day, source);

-- +goose Down
DROP INDEX IF EXISTS ux_snapshots_symbol_day_source;

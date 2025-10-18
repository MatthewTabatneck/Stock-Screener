-- +goose Up
-- 1) Remove duplicates: keep the newest row (by fetched_at then id) per (symbol, latest_trading_day, source)
WITH ranked AS (
  SELECT
    id,
    ROW_NUMBER() OVER (
      PARTITION BY symbol, latest_trading_day, source
      ORDER BY fetched_at DESC, id DESC
    ) AS rn
  FROM snapshots
),
to_delete AS (
  SELECT id FROM ranked WHERE rn > 1
)
DELETE FROM snapshots s
USING to_delete d
WHERE s.id = d.id;

-- 2) Now enforce uniqueness going forward
CREATE UNIQUE INDEX IF NOT EXISTS ux_snapshots_symbol_day_source
  ON snapshots(symbol, latest_trading_day, source);

-- +goose Down
DROP INDEX IF EXISTS ux_snapshots_symbol_day_source;

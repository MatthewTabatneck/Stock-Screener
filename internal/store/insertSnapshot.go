package store

import (
	"context"
	"database/sql"
	"errors"

	"github.com/MatthewTabatneck/stock-screener/internal/models"
)

var db *sql.DB

func SetDB(d *sql.DB) { db = d }

var errNoDB = errors.New("store: DB is nil; call store.SetDB first")

func InsertSnapshot(ctx context.Context, s models.Snapshot) error {
	if db == nil {
		return errNoDB
	}

	const q = `
        INSERT INTO snapshots (
            symbol, fetched_at, price, change_pct, volume, source, status, latest_trading_day
        ) VALUES ($1,$2,$3,$4,$5,$6,$7,$8)
        ON CONFLICT (symbol, latest_trading_day, source) DO NOTHING;
    `
	_, err := db.ExecContext(ctx, q,
		s.Symbol,
		s.FetchedAt,
		s.Price,
		s.ChangePct,
		s.Volume,
		s.Source,
		s.Status,
		s.LatestTradingDay,
	)
	return err
}

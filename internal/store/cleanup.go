package store

import (
	"context"
	"database/sql"
)

func CleanupProcessedTickers(ctx context.Context, db *sql.DB) error {
	_, err := db.ExecContext(ctx, `
		DELETE FROM tickers
		WHERE is_processed = true;
	`)
	return err
}

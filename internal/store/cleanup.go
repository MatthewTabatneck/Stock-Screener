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

// Optional: run automatic cleanup on an interval (e.g., every 10 minutes).
// func StartAutoCleanupProcessedTickers(ctx context.Context, db *sql.DB, interval time.Duration) {
// 	t := time.NewTicker(interval)
// 	go func() {
// 		defer t.Stop()
// 		for {
// 			select {
// 			case <-ctx.Done():
// 				return
// 			case <-t.C:
// 				_ = CleanupProcessedTickers(ctx, db) // ignore errors or log as you prefer
// 			}
// 		}
// 	}()
// }

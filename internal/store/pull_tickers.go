package store

import (
	"context"
	"database/sql"
)

func GetAllTickers(ctx context.Context, db *sql.DB) ([]string, error) {
	const q = `
		WITH next AS (
			SELECT symbol
			FROM tickers
			WHERE is_processed = false
			ORDER BY updated_at
			FOR UPDATE SKIP LOCKED
		)
		UPDATE tickers t
		SET is_processed = true,
			updated_at   = now()
		FROM next
		WHERE t.symbol = next.symbol
		RETURNING t.symbol;
	`

	tx, err := db.BeginTx(ctx, nil)
	if err != nil {
		return nil, err
	}
	defer tx.Rollback()

	rows, err := tx.QueryContext(ctx, q)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var symbols []string
	for rows.Next() {
		var sym string
		if err := rows.Scan(&sym); err != nil {
			return nil, err
		}
		symbols = append(symbols, sym)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	if err := tx.Commit(); err != nil {
		return nil, err
	}

	return symbols, nil
}

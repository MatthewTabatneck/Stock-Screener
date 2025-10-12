package store

import (
	"database/sql"
	"log"

	"github.com/lib/pq"
)

func InsertTickers(db *sql.DB, tickers []string) {
	//Delete rows not in the new list
	_, err := db.Exec(`
        WITH new(symbol) AS (SELECT UNNEST($1::text[]))
        DELETE FROM tickers t
        WHERE NOT EXISTS (SELECT 1 FROM new n WHERE n.symbol = t.symbol);
    `, pq.Array(tickers))
	if err != nil {
		log.Fatal(err)
	}

	//Insert missing ones
	_, err = db.Exec(`
        INSERT INTO tickers (symbol)
        SELECT UNNEST($1::text[])
        ON CONFLICT (symbol) DO NOTHING;
    `, pq.Array(tickers))
	if err != nil {
		log.Fatal(err)
	}
}

package main

import (
	"context"
	"database/sql"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/MatthewTabatneck/stock-screener/internal/config"
	_ "github.com/lib/pq"
)

func main() {
	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	defer stop()

	config := config.Load()

	db, err := sql.Open("postgres", config.DatabaseURL)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	tickers, err := GetAllTickers(ctx)
}

package main

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/MatthewTabatneck/stock-screener/internal/screener"
	"github.com/MatthewTabatneck/stock-screener/internal/store"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

func main() {
	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	defer stop()

	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found, using environment variables")
	}

	file, err := os.Open("tickers.csv")
	if err != nil {
		fmt.Println("Error opening file: ", err)
		return
	}
	defer file.Close()

	tickers, err := screener.LoadtickersCSV(file)
	if err != nil {
		fmt.Println(err)
		return
	}

	dbURL := os.Getenv("DATABASE_URL")
	if dbURL == "" {
		log.Fatal("DATABASE_URL not set")
	}

	db, err := sql.Open("postgres", dbURL)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	if err = db.Ping(); err != nil {
		log.Fatal(err)
	}
	store.SetDB(db)
	store.InsertTickers(ctx, db, tickers)

}

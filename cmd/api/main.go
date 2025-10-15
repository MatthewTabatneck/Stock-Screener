package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/MatthewTabatneck/stock-screener/internal/screener"
	"github.com/MatthewTabatneck/stock-screener/internal/store"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

// func main() {
// 	mux := http.NewServeMux()
// 	mux.HandleFunc("/healthz", func(w http.ResponseWriter, r *http.Request) {
// 		fmt.Fprintln(w, "ok")
// 	})
// 	port := "8080"
// 	log.Printf("listening on :%s", port)
// 	log.Fatal(http.ListenAndServe(":"+port, mux))
// }

func main() {
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

	store.InsertTickers(db, tickers)

}

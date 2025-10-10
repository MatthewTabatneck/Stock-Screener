package main

import (
	"fmt"
	"os"

	"github.com/MatthewTabatneck/stock-screener/internal/screener"
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
	file, err := os.Open("sp500_tickers.csv")
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

	fmt.Println(tickers)

}

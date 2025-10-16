package provider

import (
	"context"
	"log"
	"time"

	"github.com/MatthewTabatneck/stock-screener/internal/config"
	"github.com/MatthewTabatneck/stock-screener/internal/provider/alpha"
	"golang.org/x/time/rate"
)

type Result struct {
	Symbol string
	Snap   alpha.Snapshot
	Err    error
}

// This func needs to do the following:
// - run a single thread go routine that will every 13 seconds call the alpa api
// the data wil be stored within the struct results
func GetAlpha(ctx context.Context, cfg config.Config, symbols []string) error {

	limiter := rate.NewLimiter(rate.Every(cfg.MinInterval()*time.Second), 1)

	results := make(chan Result)

	go func() {
		for _, sym := range symbols {
			if err := limiter.Wait(ctx); err != nil {
				results <- Result{Symbol: sym, Err: err}
				continue
			}

			snap, err := alpha.FetchAlpha(ctx, sym, cfg.AlphaKey)

			results <- Result{Symbol: sym, Snap: snap, Err: err}
		}
	}()

	for i := 0; i < cfg.Processors; i++ {
		go func(id int) {
			for r := range results {
				if r.Err != nil {
					log.Printf("[proc %d] error fetching %s: %v", id, r.Symbol, r.Err)
					continue
				}

			}
		}(i)
	}
}

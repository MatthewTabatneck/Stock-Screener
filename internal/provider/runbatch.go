package provider

import (
	"context"
	"log"
	"sync"
	"time"

	"github.com/MatthewTabatneck/stock-screener/internal/config"
	"github.com/MatthewTabatneck/stock-screener/internal/models"
	"github.com/MatthewTabatneck/stock-screener/internal/provider/alpha"
	"github.com/MatthewTabatneck/stock-screener/internal/store"
	"golang.org/x/time/rate"
)

type Result struct {
	Symbol string
	Snap   models.Snapshot
	Err    error
}

// This func needs to do the following:
// - run a single thread go routine that will every 13 seconds call the alpa api
// the data wil be stored within the struct results
func GetAlpha(ctx context.Context, cfg config.Config, symbols []string) error {

	limiter := rate.NewLimiter(rate.Every(cfg.MinInterval()*time.Second), 1)

	results := make(chan Result)

	go func() {
		defer close(results)
		for _, sym := range symbols {
			if err := limiter.Wait(ctx); err != nil {
				results <- Result{Symbol: sym, Err: err}
				continue
			}

			snap, err := alpha.FetchAlpha(ctx, sym, cfg.AlphaKey)

			results <- Result{Symbol: sym, Snap: snap, Err: err}
		}
	}()

	var wg sync.WaitGroup
	for i := 0; i < cfg.Processors; i++ {
		wg.Add(1)
		go func(id int) {
			for r := range results {
				if r.Err != nil {
					log.Printf("[proc %d] error fetching %s: %v", id, r.Symbol, r.Err)
					continue
				}
				if err := store.InsertSnapshot(ctx, r.Snap); err != nil {
					log.Printf("[proc %d] insert %s: %v", id, r.Symbol, err)
					continue
				}
			}
		}(i)
	}
	wg.Wait()
	return nil
}

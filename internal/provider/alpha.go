package provider


limiter := rate.NewLimiter(rate.Every(13*time.Second), 1)

func AlphaProvider(ctx context.Contex, )
// 1) Fetcher: single goroutine respecting the limiter
go func() {
    for sym := range fetchQueue {
        if err := limiter.Wait(ctx); err != nil { continue }
        snap, err := alpha.Fetch(ctx, sym) // one external call
        results <- Result{Sym: sym, Snap: snap, Err: err}
    }
}()

// 2) Processors: small pool for local work (DB, transforms)
for i := 0; i < 3; i++ {
    go func() {
        for r := range results {
			// check if fetch call haad an error and for which ticker
			if r.Err != nil {
				fmt.Printf("Error fetching %s: %v\n", r.Sym, r.Err)
				failed := Snapshot{Symbol: r.Sym, Status: "ERROR", Source: "alpha_vantage"}
				_ = store.InsertSnapshot(ctx, failed)
        	continue
    		}

			fmt.Printf("%s %.2f (%.2f%%) Vol %d\n", r.Sym, r.Snap.Price, r.Snap.ChangePct, r.Snap.Volume)
            // parse/validate, compute derived fields, insert into DB
            _ = store.InsertSnapshot(ctx, r.SnapOrFailure())
        }
    }()
}
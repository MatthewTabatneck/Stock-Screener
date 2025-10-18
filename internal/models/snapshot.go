// internal/models/snapshot.go
package models

import "time"

type Snapshot struct {
	Symbol           string
	Price            float64
	ChangePct        float64
	Volume           int64
	Source           string
	Status           string
	FetchedAt        time.Time
	LatestTradingDay *time.Time // pointer so it can be nil
}

package alpha

import (
	"context"
	"time"
)

type Snapshot struct {
	Symbol    string    `json:"symbol"`
	Price     float64   `json:"price,omitempty"`
	ChangePct float64   `json:"change_pct,omitempty"`
	Volume    int64     `json:"volume,omitempty"`
	FetchedAt time.Time `json:"fetched_at"`
	Source    string    `json:"source"`
	Status    string    `json:"status"` // "OK", "TIMEOUT", "ERROR", "RATELIMIT"
}

func FetchAlpha(ctx context.Context, symbol, apiKey string) (Snapshot, error) {

}

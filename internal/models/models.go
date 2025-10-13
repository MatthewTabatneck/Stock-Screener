package models

import "time"

type Snapshot struct {
	Symbol    string    `json:"symbol"`
	Price     float64   `json:"price,omitempty"`
	ChangePct float64   `json:"change_pct,omitempty"`
	Volume    int64     `json:"volume,omitempty"`
	FetchedAt time.Time `json:"fetched_at"`
	Source    string    `json:"source"`
	Status    string    `json:"status"` // "OK", "TIMEOUT", "ERROR", "RATELIMIT"
}

type Result struct {
	Sym  string
	Snap Snapshot
	Err  error
}

func (r Result) SnapOrFailure(source string) Snapshot {
	if r.Err == nil {
		return r.Snap
	}
	return Snapshot{
		Symbol:    r.Sym,
		Source:    source,
		Status:    "ERROR",
		FetchedAt: time.Now().UTC(),
	}
}

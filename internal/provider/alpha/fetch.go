package alpha

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/MatthewTabatneck/stock-screener/internal/models"
)

var (
	ErrRateLimited = errors.New("alpha vantage: rate limited")
	alphaBaseURL   = "https://www.alphavantage.co/query"
)

func FetchAlpha(ctx context.Context, symbol, apiKey string) (models.Snapshot, error) {
	req, _ := http.NewRequestWithContext(ctx, http.MethodGet, alphaBaseURL, nil)
	q := req.URL.Query()
	q.Set("function", "GLOBAL_QUOTE")
	q.Set("symbol", symbol)
	q.Set("apikey", apiKey)
	req.URL.RawQuery = q.Encode()

	client := &http.Client{Timeout: 5 * time.Second}
	resp, err := client.Do(req)
	if err != nil {
		return models.Snapshot{Symbol: symbol, Status: "ERROR"}, err
	}
	defer resp.Body.Close()

	var body map[string]any
	if err := json.NewDecoder(resp.Body).Decode(&body); err != nil {
		return models.Snapshot{Symbol: symbol, Status: "ERROR"}, err
	}

	// Alpha Vantage rate-limit messages
	if note, ok := body["Note"].(string); ok && note != "" {
		return models.Snapshot{Symbol: symbol, Status: "RATE_LIMIT"}, fmt.Errorf("%w: %s", ErrRateLimited, note)
	}
	if info, ok := body["Information"].(string); ok && info != "" {
		return models.Snapshot{Symbol: symbol, Status: "RATE_LIMIT"}, fmt.Errorf("%w: %s", ErrRateLimited, info)
	}

	raw, ok := body["Global Quote"].(map[string]any)
	if !ok || len(raw) == 0 {
		return models.Snapshot{Symbol: symbol, Status: "ERROR"}, fmt.Errorf("unexpected JSON: %v", body)
	}

	// Helpers to pull strings safely
	getStr := func(k string) string {
		if v, ok := raw[k].(string); ok {
			return v
		}
		return ""
	}

	// Parse fields (best-effort)
	price, _ := strconv.ParseFloat(getStr("05. price"), 64)
	vol, _ := strconv.ParseInt(getStr("06. volume"), 10, 64)
	chStr := strings.TrimSuffix(getStr("10. change percent"), "%")
	chPct, _ := strconv.ParseFloat(chStr, 64)

	// Parse latest trading day as pointer
	var ltdPtr *time.Time
	if dayStr := getStr("07. latest trading day"); dayStr != "" {
		if t, err := time.Parse("2006-01-02", dayStr); err == nil {
			ltdPtr = &t // <-- pointer value
		}
	}

	return models.Snapshot{
		Symbol:           symbol,
		Price:            price,
		ChangePct:        chPct,
		Volume:           vol,
		LatestTradingDay: ltdPtr, // matches *time.Time in your model
		FetchedAt:        time.Now().UTC(),
		Source:           "alpha_vantage",
		Status:           "OK",
	}, nil
}

package screener

import (
	"encoding/csv"
	"errors"
	"io"
	"strings"
)

// func will read  csv file, store those values into a struct to get rid of duplicates, then transform into a []string for use
func LoadtickersCSV(r io.Reader) ([]string, error) {
	cr := csv.NewReader(r)
	records, err := cr.ReadAll()
	if err != nil {
		return nil, err
	}

	//ends program early if the csv is empty
	if len(records) == 0 {
		return nil, errors.New("input cannot be empty")
	}

	seen := make(map[string]struct{}, len(records))

	for _, rec := range records {
		if len(rec) == 0 {
			continue
		}
		sym := strings.TrimSpace(rec[0])
		if sym == "" {
			continue
		}
		sym = strings.ToUpper(strings.ReplaceAll(sym, " ", "")) //removes spaces then upper cases all characters
		seen[sym] = struct{}{}                                  //sets tickers into struct and skips duplicates
	}

	// Emit final []string (order is does not matter in this step)
	out := make([]string, 0, len(seen))
	for sym := range seen {
		out = append(out, sym)
	}
	return out, nil
}

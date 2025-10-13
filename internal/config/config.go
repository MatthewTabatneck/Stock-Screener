package config

import (
	"os"
	"strconv"
	"time"
)

type Config struct {
	AlphaKey       string
	MinIntervalSec int
	Processors     int
	DatabaseURL    string
}

func Load() Config {
	return Config{
		AlphaKey:       os.Getenv("ALPHA_KEY"),
		MinIntervalSec: mustInt(os.Getenv("MIN_INTERVAL_SEC"), 13),
		Processors:     mustInt(os.Getenv("PROCESSORS"), 3),
		DatabaseURL:    os.Getenv("DATABASE_URL"),
	}
}

func (c Config) MinInterval() time.Duration {
	return time.Duration(c.MinIntervalSec) * time.Second
}

func mustInt(s string, def int) int {
	if s == "" {
		return def
	}
	i, err := strconv.Atoi(s)
	if err != nil {
		return def
	}
	return i
}

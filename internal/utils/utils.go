package utils

import (
	"os"
)

func EnvVarExists(key string) bool {
	val, ok := os.LookupEnv(key)
	if ok {
		return true
	}
	return len(val) != 0
}

func EnvVarOrFallback(key, fallback string) string {
	if EnvVarExists(key) {
		return os.Getenv(key)
	}
	return fallback
}

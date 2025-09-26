package utils

import "os"

func GetEnv(key string, fallback ...string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	if len(fallback) > 0 {
		return fallback[0]
	}
	return ""
}

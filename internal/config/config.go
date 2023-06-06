package config

import "os"

type Config struct {
	OsrmURL string
	Port    string
}

func LoadConfig() Config {
	return Config{
		OsrmURL: getEnvOrDefault("OSRM_URL", "http://router.project-osrm.org/route/v1/driving"),
		Port:    getEnvOrDefault("PORT", "8080"),
	}
}

func getEnvOrDefault(key, defaultValue string) string {
	value := os.Getenv(key)
	if value == "" {
		return defaultValue
	}
	return value
}

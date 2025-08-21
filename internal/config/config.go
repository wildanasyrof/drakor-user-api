package config

import (
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

// Config holds all the configuration variables for the application.
// We use a struct to provide type safety and easy access.
type Config struct {
	Port               int
	DatabaseURL        string
	JWTSecret          string
	ScraperBaseURL     string
	AccessTokenMinutes int
	RefreshTokenDays   int
	UploadDir          string
	Env                string
}

// LoadConfig reads configuration from environment variables and returns a Config struct.
func LoadConfig() (*Config, error) {
	// Attempt to load .env file. It's OK if it doesn't exist, as it might be
	// running in a production environment with env variables already set.
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found, reading from environment variables")
	}

	// Helper function to get a string environment variable with a default value.
	getEnv := func(key, defaultValue string) string {
		if value, exists := os.LookupEnv(key); exists {
			return value
		}
		return defaultValue
	}

	// Helper function to get an int environment variable.
	getEnvAsInt := func(key string) (int, error) {
		valueStr := os.Getenv(key)
		if value, err := strconv.Atoi(valueStr); err == nil {
			return value, nil
		}
		// Return 0 and an error if the key doesn't exist or is not an integer.
		return 0, nil
	}

	port, err := getEnvAsInt("PORT")
	if err != nil {
		log.Fatal("PORT must be a valid integer")
	}

	accessTokenMinutes, err := getEnvAsInt("ACCESS_TOKEN_MINUTES")
	if err != nil {
		log.Fatal("ACCESS_TOKEN_MINUTES must be a valid integer")
	}

	refreshTokenDays, err := getEnvAsInt("REFRESH_TOKEN_DAYS")
	if err != nil {
		log.Fatal("REFRESH_TOKEN_DAYS must be a valid integer")
	}

	cfg := &Config{
		Port:               port,
		DatabaseURL:        getEnv("DATABASE_URL", ""),
		JWTSecret:          getEnv("JWT_SECRET", ""),
		ScraperBaseURL:     getEnv("SCRAPER_BASE_URL", ""),
		AccessTokenMinutes: accessTokenMinutes,
		RefreshTokenDays:   refreshTokenDays,
		UploadDir:          getEnv("UPLOAD_DIR", "./uploads"),
		Env:                getEnv("ENV", "development"),
	}

	// Example of a basic validation.
	if cfg.DatabaseURL == "" {
		log.Fatal("DATABASE_URL is not set")
	}

	if cfg.JWTSecret == "" {
		log.Fatal("JWT_SECRET is not set")
	}

	return cfg, nil
}

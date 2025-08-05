package config

import (
	"log"
	"os"
	"strings"
	"time"

	"github.com/joho/godotenv"
)

type Config struct {
	// Database
	DBHost     string
	DBPort     string
	DBName     string
	DBUser     string
	DBPassword string

	// JWT
	JWTSecret        string
	JWTAccessExpiry  time.Duration
	JWTRefreshExpiry time.Duration

	// Server
	Port string
	Env  string

	// CORS
	CORSOrigins string

	// QR Code
	QRCodeSize int

	// Timezone
	Timezone string

	// Resend
	ResendAPIKey string
	ResendFrom   string
	ResendCc     []string
	ResendBcc    []string

	// Frontend
	FrontendURL string

	// Password Reset
	PasswordResetTokenLifespan time.Duration
}

func Load() *Config {
	// Load .env file if it exists
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found, using environment variables")
	}

	// Parse JWT expiry durations
	accessExpiry, err := time.ParseDuration(getEnv("JWT_ACCESS_EXPIRY", "15m"))
	if err != nil {
		log.Fatal("Invalid JWT_ACCESS_EXPIRY format:", err)
	}

	refreshExpiry, err := time.ParseDuration(getEnv("JWT_REFRESH_EXPIRY", "168h"))
	if err != nil {
		log.Fatal("Invalid JWT_REFRESH_EXPIRY format:", err)
	}

	passwordResetLifespan, err := time.ParseDuration(getEnv("PASSWORD_RESET_TOKEN_LIFESPAN", "24h"))
	if err != nil {
		log.Fatal("Invalid PASSWORD_RESET_TOKEN_LIFESPAN format:", err)
	}

	return &Config{
		DBHost:           getEnv("DB_HOST", "localhost"),
		DBPort:           getEnv("DB_PORT", "27017"),
		DBName:           getEnv("DB_NAME", "church_attendance_db"),
		DBUser:           getEnv("DB_USER", ""),
		DBPassword:       getEnv("DB_PASSWORD", ""),
		JWTSecret:        getEnv("JWT_SECRET", "your-super-secret-jwt-key-change-this-in-production"),
		JWTAccessExpiry:  accessExpiry,
		JWTRefreshExpiry: refreshExpiry,
		Port:             getEnv("PORT", "8080"),
		Env:              getEnv("ENV", "development"),
		CORSOrigins:      getEnv("CORS_ORIGINS", "http://localhost:3000,http://localhost:8080"),
		QRCodeSize:       256,
		Timezone:         getEnv("TIMEZONE", "Africa/Lagos"),
		ResendAPIKey:     getEnv("RESEND_API_KEY", ""),
		ResendFrom:       getEnv("RESEND_FROM", ""),
		ResendCc:                   getEnvAsSlice("RESEND_CC", []string{}),
		ResendBcc:                  getEnvAsSlice("RESEND_BCC", []string{}),
		FrontendURL:                getEnv("FRONTEND_URL", "http://localhost:3000"),
		PasswordResetTokenLifespan: passwordResetLifespan,
	}
}

func getEnvAsSlice(name string, defaultVal []string) []string {
	valStr := getEnv(name, "")
	if valStr == "" {
		return defaultVal
	}
	return strings.Split(valStr, ",")
}

func getEnv(key, defaultValue string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return defaultValue
}

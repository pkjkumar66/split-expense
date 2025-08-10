package config

import "os"

type Config struct {
	DatabaseURL         string
	JWTSecret           string
	ServerPort          string
	WhatsAppToken       string
	WhatsAppVerifyToken string
}

func Load() *Config {
	return &Config{
		DatabaseURL:         getEnv("DATABASE_URL", "postgresql://myuser:mypassword@localhost:5432/mydb?sslmode=disable"),
		JWTSecret:           getEnv("JWT_SECRET", "your-secret-key"),
		ServerPort:          getEnv("PORT", "8080"),
		WhatsAppToken:       getEnv("WHATSAPP_TOKEN", ""),
		WhatsAppVerifyToken: getEnv("WHATSAPP_VERIFY_TOKEN", ""),
	}
}

func getEnv(key, defaultValue string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return defaultValue
}

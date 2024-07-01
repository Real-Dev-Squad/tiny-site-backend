package config

import (
	"fmt"
	"os"
	"path/filepath"
	"strconv"

	"github.com/Real-Dev-Squad/tiny-site-backend/logger"
	"github.com/joho/godotenv"
)

var (
	Env                 string
	UserMaxUrlCount         int
	Domain              string
	AuthRedirectUrl     string
	DbUrl               string
	DbMaxOpenConnections int
	GoogleClientId      string
	GoogleClientSecret  string
	GoogleRedirectUrl   string
	TokenValidity     int
	JwtSecret           string
	JwtValidity         int
	JwtIssuer           string
	AllowedOrigin		string
	Port                string
)

func findAndLoadEnv(envFile string) {
	cwd, err := os.Getwd()
	if err != nil {
		logger.Fatal("Could not get current working directory:", err)
	}

	for {
		envPath := filepath.Join(cwd, envFile)
		if _, err := os.Stat(envPath); err == nil {
			if err := godotenv.Load(envPath); err != nil {
				logger.Fatal("Error loading .env file:", err)
			}
			logger.Info("Loaded environment variables from:", envPath)
			return
		}

		parent := filepath.Dir(cwd)
		if parent == cwd {
			break
		}
		cwd = parent
	}

	logger.Fatal("Could not find .env file at:", envFile)
}

func loadEnv() {
	env := os.Getenv("ENV")
	if env == "production" || env == "staging" {
		return
	}

	envFile := ".env"
	if env == "test" {
		envFile = "environments/test.env"
	}

	findAndLoadEnv(envFile)
}

func init() {
	loadEnv()

	env := os.Getenv("ENV")
	if env == "" {
		Env = "dev"
	} else {
		Env = env
	}

	loadConfig()
	logger.Info("Loaded environment variables")
}

func loadConfig() {
	JwtSecret = getEnvVar("JWT_SECRET")
	JwtIssuer = getEnvVar("JWT_ISSUER")

	Domain = getEnvVar("DOMAIN")
	AuthRedirectUrl = getEnvVar("AUTH_REDIRECT_URL")

	DbUrl = getEnvVar("DB_URL")
	DbMaxOpenConnections = getEnvInt("DB_MAX_OPEN_CONNECTIONS")

	GoogleClientId = getEnvVar("GOOGLE_CLIENT_ID")
	GoogleClientSecret = getEnvVar("GOOGLE_CLIENT_SECRET")
	GoogleRedirectUrl = getEnvVar("GOOGLE_REDIRECT_URL")

	UserMaxUrlCount = getEnvInt("USER_MAX_URL_COUNT")
	TokenValidity = getEnvInt("TOKEN_VALIDITY_IN_SECONDS")

	AllowedOrigin = getEnvVar("ALLOWED_ORIGINS")

	Port = getEnvVar("PORT")
	JwtValidity = getEnvInt("JWT_VALIDITY_IN_HOURS")
}

func getEnvVar(key string) string {
	value := os.Getenv(key)
	if value == "" {
		logger.Fatal(fmt.Sprintf("Environment variable %s not set", key))
	}
	return value
}

func getEnvInt(key string) int {
	valueStr := os.Getenv(key)
	value, err := strconv.Atoi(valueStr)
	if err != nil {
		logger.Fatal(fmt.Sprintf("Error parsing environment variable %s: %v", key, err))
	}
	return value
}

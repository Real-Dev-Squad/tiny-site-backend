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
	MaxUrlCount         int
	Domain              string
	AuthRedirectUrl     string
	DbUrl               string
	DbMaxOpenConnections int
	GoogleClientId      string
	GoogleClientSecret  string
	GoogleRedirectUrl   string
	TokenExpiration     int
	JwtSecret           string
	JwtValidity         int
	JwtIssuer           string
	AllowedOrigin		string
)

// findAndLoadEnv attempts to load the .env file from the current directory or any parent directory.
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
	// If the environment is production or staging, we don't need to load the .env file
	// we assume that the environment variables are already set
	if env == "production" || env == "staging" {
		return
	}

	envFile := ".env"
	if env == "test" {
		envFile = "environments/env.test"
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

	MaxUrlCount = getEnvInt("Max_Url_Count")
	TokenExpiration = getEnvInt("TokenExpiration")

	AllowedOrigin = getEnvVar("ALLOWED_ORIGINS")
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

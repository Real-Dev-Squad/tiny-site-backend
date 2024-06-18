package config

import (
	"os"
	"path"
	"runtime"
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
	JwtSecret			string
	JwtValidity		    int
	JwtIssuer			string

)

func loadEnv() {
	env := os.Getenv("ENV")

	// If the environment is production, we don't need to load the .env file
	// we assume that the environment variables are already set
	if env == "production" || env == "staging" {
		return
	}

	if env == "test" {
		_, filename, _, _ := runtime.Caller(0)
		dir := path.Join(path.Dir(filename), "../..")

		if err := os.Chdir(dir); err != nil {
			panic(err)
		}

		if err := godotenv.Load(".env"); err != nil {
			logger.Error("Error loading .env file.", err)
		}

		return
	}

	if err := godotenv.Load(".env"); err != nil {
		logger.Fatal("Error loading .env file")
	}
}

func init() {
	loadEnv()

	env := os.Getenv("ENV")

	if env == "" {
		Env = "dev"
	} else {
		Env = env
	}

	JwtSecret = os.Getenv("JWT_SECRET")
	JwtValidity, _ = strconv.Atoi(os.Getenv("JWT_VALIDITY_IN_DAYS"))
	JwtIssuer = os.Getenv("JWT_ISSUER")

	Domain = os.Getenv("DOMAIN")
	AuthRedirectUrl = os.Getenv("AUTH_REDIRECT_URL")

	DbUrl = os.Getenv("DB_URL")
	DbMaxOpenConnections, _ = strconv.Atoi(os.Getenv("DB_MAX_OPEN_CONNECTIONS"))

	GoogleClientId = os.Getenv("GOOGLE_CLIENT_ID")
	GoogleClientSecret = os.Getenv("GOOGLE_CLIENT_SECRET")
	GoogleRedirectUrl = os.Getenv("GOOGLE_REDIRECT_URL")

	MaxUrlCount, _ =  strconv.Atoi(os.Getenv("Max_Url_Count"))
	TokenExpiration, _ =  strconv.Atoi(os.Getenv("TokenExpiration"))

	logger.Info("Loaded environment variables")
}

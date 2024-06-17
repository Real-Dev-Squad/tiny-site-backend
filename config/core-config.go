package config

import "os"

var Max_Url_Count string

func loadEnv() {
	env := os.Getenv(".env")
	if env == "production" || env == "staging" {
		return
	}
}

func init() {
	loadEnv()

	env := os.Getenv(".env")
	if env == "" {

	}

	Max_Url_Count = os.Getenv("MAX_URL_COUNT")

}

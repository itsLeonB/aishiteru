package config

import (
	"log"
	"time"

	"github.com/itsLeonB/ezutil"
)

func Defaults() ezutil.Config {
	timeout, _ := time.ParseDuration("10s")
	tokenDuration, _ := time.ParseDuration("24h")
	cookieDuration, _ := time.ParseDuration("24h")
	secretKey, err := ezutil.GenerateRandomString(32)
	if err != nil {
		log.Fatalf("error generating secret key: %v", err)
	}

	appConfig := ezutil.App{
		Env:        "debug",
		Port:       "50051",
		Timeout:    timeout,
		ClientUrls: []string{"http://localhost:3000"},
		Timezone:   "Asia/Jakarta",
	}

	authConfig := ezutil.Auth{
		SecretKey:      secretKey,
		TokenDuration:  tokenDuration,
		CookieDuration: cookieDuration,
		Issuer:         "aishiteru",
		URL:            "http://localhost:8000",
	}

	llmConfig := LLM{}

	return ezutil.Config{
		App:     &appConfig,
		Auth:    &authConfig,
		Generic: &llmConfig,
	}
}

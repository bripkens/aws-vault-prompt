package main

import (
	"fmt"
	"os"
	"time"
)

type Configuration struct {
	vault              string
	expiration         string
	region             string
	prompt             string
	promptExpiringSoon string
	promptExpired      string
}

func main() {
	configuration := getConfiguration()

	if len(configuration.vault) == 0 || len(configuration.expiration) == 0 {
		return
	}

	parsedExpiration, err := time.Parse(time.RFC3339, configuration.expiration)
	if err != nil {
		return
	}

	promptFormat := configuration.prompt
	remainingSessionDuration := parsedExpiration.Sub(time.Now())

	if remainingSessionDuration < 0 {
		promptFormat = configuration.promptExpired
	} else if remainingSessionDuration < (time.Minute * 15) {
		promptFormat = configuration.promptExpiringSoon
	}

	fmt.Printf(promptFormat, configuration.vault, configuration.region, int(remainingSessionDuration.Minutes()))
}

func getConfiguration() Configuration {
	return Configuration{
		vault:              os.Getenv("AWS_VAULT"),
		expiration:         os.Getenv("AWS_SESSION_EXPIRATION"),
		region:             os.Getenv("AWS_REGION"),
		prompt:             os.Getenv("AWS_VAULT_PROMPT"),
		promptExpiringSoon: os.Getenv("AWS_VAULT_PROMPT_EXPIRING_SOON"),
		promptExpired:      os.Getenv("AWS_VAULT_PROMPT_EXPIRED"),
	}
}

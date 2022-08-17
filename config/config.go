package config

import (
	"encoding/json"
	"os"
)

type configStruct struct {
	Token     string `json:"Token"`
	BotPrefix string `json:"BotPrefix"`
}

var (
	config    configStruct
	Token     string
	BotPrefix string
)

func ReadConfig() error {
	data, err := os.ReadFile("config.json")

	if err != nil {
		return err
	}

	err = json.Unmarshal(data, &config)

	if err != nil {
		return err
	}

	Token = config.Token
	BotPrefix = config.BotPrefix

	return nil
}

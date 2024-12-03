package config

import (
	"encoding/json"
	"os"
)

const configFileName = ".gatorconfig.json"

type Config struct {
	DBUrl           string `json:"db_url"`
	CurrentUserName string `json:"current_user_name "`
}

func Read() (Config, error) {
	filePath, err := getConfigFilePath()
	if err != nil {
		return Config{}, err
	}

	file, err := os.Open(filePath)
	if err != nil {
		return Config{}, err
	}

	var cfg Config
	decoder := json.NewDecoder(file)
	if err := decoder.Decode(&cfg); err != nil {
		return Config{}, err
	}
	return cfg, nil
}

func getConfigFilePath() (string, error) {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		return "", err
	}

	return homeDir + "/" + configFileName, nil
}

func (c *Config) SetUser(user string) error {
	c.CurrentUserName = user
	err := write(c)

	return err
}

func write(cfg *Config) error {
	configFilePath, err := getConfigFilePath()
	if err != nil {
		return err
	}

	file, err := os.Create(configFilePath)
	if err != nil {
		return err
	}
	defer file.Close()

	configJson, err := json.MarshalIndent(cfg, "", "  ") // Indentation for readability
	if err != nil {
		return err
	}

	// Write JSON to the file
	_, err = file.Write(configJson)
	return err
}

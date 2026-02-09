package config

import (
	"encoding/json"
	"os"
	"path/filepath"
)

const configFileName = "/.gatorconfig.json"

type Config struct {
	CurrentUserName string `json:"current_user_name"`
	DbUrl           string `json:"db_url"`
}

func getConfigFilePath() (string, error) {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		return "", err
	}
	configFile := filepath.Join(homeDir, configFileName)

	if _, err := os.Stat(configFile); os.IsNotExist(err) {
		return "", err
	}

	return configFile, nil
}

func Read() (Config, error) {

	configFile, err := getConfigFilePath()
	if err != nil {
		return Config{}, err
	}

	file, err := os.Open(configFile)
	defer file.Close()

	if err != nil {
		return Config{}, err
	}

	var cfg Config
	err = json.NewDecoder(file).Decode(&cfg)
	if err != nil {
		return Config{}, err
	}

	return cfg, nil
}

func write(cfg Config) error {
	configFile, err := getConfigFilePath()
	if err != nil {
		return err
	}

	file, err := os.Create(configFile)
	defer file.Close()

	if err != nil {
		return err
	}

	return json.NewEncoder(file).Encode(cfg)
}

func (c *Config) SetUser(user string) error {
	c.CurrentUserName = user
	return write(*c)
}

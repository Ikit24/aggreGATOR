package config

import (
	"encoding/json"
	"os"
	"path/filepath"
)

type Config struct {
	DBURL			string `json:"db_url"`
	CurrentUserName string `json:"current_user_name"`
}

func Read() (Config, error) {
	var cfg Config

	home, err := os.UserHomeDir()
	if err != nil {
		return cfg, err
	}

	path := filepath.Join(home, ".gatorconfig.json")

	data, err := os.ReadFile(path)
	if err != nil {
		return cfg, err
	}

	if err := json.Unmarshal(data, &cfg); err != nil {
		return cfg,  err
	}

	return cfg, nil
}

func (c *Config) SetUser(name string) error {
	c.CurrentUserName = name
	
	home, err := os.UserHomeDir()
	if err != nil {
		return err
	}

	path := filepath.Join(home, ".gatorconfig.json")

	data, err := json.Marshal(c)
	if err != nil {
		return err
	}

	if err := os.WriteFile(path, data, 0o600); err != nil {
		return err
	}
	
	return nil
}

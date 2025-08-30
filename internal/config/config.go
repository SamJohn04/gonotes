package config

import (
	"encoding/json"
	"log"
	"os"
	"path/filepath"
)

type StyleConfig struct {
	BackgroundColor string `json:"background,omitempty"`
	ForegroundColor string `json:"foreground,omitempty"`
}

func Load() StyleConfig {
	cfg := StyleConfig{
		BackgroundColor: "#232627",
		ForegroundColor: "#fefefe",
	}

	home, err := os.UserHomeDir()
	if err != nil {
		log.Println("Could not access home directory.")
		return cfg
	}

	loadStyles(filepath.Join(home, ".nate", "nate-config.json"), &cfg)
	return cfg
}

func loadStyles(path string, cfg *StyleConfig) {
	data, err := os.ReadFile(path)
	if err != nil {
		return
	}

	if err := json.Unmarshal(data, &cfg); err != nil {
		log.Println(err)
	}
}

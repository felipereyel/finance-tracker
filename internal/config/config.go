package config

import (
	"os"

	_ "github.com/joho/godotenv/autoload"
)

type ServerConfigs struct {
	AutoMigrate bool
}

func GetServerConfigs() ServerConfigs {
	config := ServerConfigs{}

	envAutoMigrate := os.Getenv("AUTOMIGRATE")
	if envAutoMigrate != "" {
		config.AutoMigrate = envAutoMigrate == "true"
	} else {
		config.AutoMigrate = true
	}

	return config
}

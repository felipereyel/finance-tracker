package config

import (
	"os"

	_ "github.com/joho/godotenv/autoload"
)

type envs struct {
	PublicDir   string
	Automigrate bool
}

var Envs = envs{
	PublicDir:   os.Getenv("PUBLIC_DIR"),
	Automigrate: os.Getenv("AUTOMIGRATE") == "true",
}

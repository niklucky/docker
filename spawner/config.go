package spawner

import (
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/niklucky/docker/pkg/spawner/api"
	"github.com/niklucky/docker/pkg/spawner/docker"
)

const DEFAULT_SERVER_ADDRESS = ":8080"

type Config struct {
	API    *api.Config
	Docker *docker.Config
}

func NewConfig(config *Config) *Config {
	if config == nil {
		config = &Config{}
	}
	err := godotenv.Load()
	if err != nil {
		log.Println("Error loading .env file")
	}
	config.API = checkAPIConfig(config.API)
	return config
}

func checkAPIConfig(config *api.Config) *api.Config {
	if config == nil {
		config = &api.Config{}
	}
	if config.ServerAddress == "" {
		config.ServerAddress = os.Getenv("SERVER_ADDRESS")
		if config.ServerAddress == "" {
			config.ServerAddress = DEFAULT_SERVER_ADDRESS
		}
	}
	if config.SecretKey == "" {
		config.SecretKey = os.Getenv("SECRET_KEY")
		if config.SecretKey == "" {
			log.Fatalln("You need to provide Secret key")
		}
	}
	return config
}

package config

import (
	"github.com/joho/godotenv"
	"log"
	"os"
	"strconv"
	"time"
)

type ServerConfig struct {
	ServerPort     int
	PostgresConfig *PostgresConfig
	YoutubeConfig  *YoutubeConfig
	ProbeDuration  time.Duration
}
type PostgresConfig struct {
	Host         string
	UserName     string
	Password     string
	DatabaseName string
	Port         string
}

var Configuration *ServerConfig

type YoutubeConfig struct {
	BaseUrl string
	ApiKey  string
}

func init() {

	var err = godotenv.Load("config/.env")
	if err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}

	port, err := strconv.Atoi(os.Getenv("SERVER_PORT"))

	if err != nil {
		log.Panic("Invalid Port", err)
	}

	duration, err := strconv.Atoi(os.Getenv("PROBE_DURATION_IN_SECONDS"))

	if err != nil {
		duration = 10
	}

	Configuration = &ServerConfig{
		ServerPort: port,
		PostgresConfig: &PostgresConfig{
			Host:         os.Getenv("POSTGRES_HOST"),
			UserName:     os.Getenv("POSTGRES_USER"),
			Password:     os.Getenv("POSTGRES_PASSWORD"),
			DatabaseName: os.Getenv("POSTGRES_DATABASE"),
			Port:         os.Getenv("POSTGRES_PORT"),
		},
		YoutubeConfig: &YoutubeConfig{
			BaseUrl: os.Getenv("YOUTUBE_URL"),
			ApiKey:  os.Getenv("YOUTUBE_API_KEY"),
		},
		ProbeDuration: time.Duration(duration) * time.Second,
	}

}

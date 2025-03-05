package config

type ServerConfig struct {
	ServerPort     int
	PostgresConfig *PostgresConfig
}
type PostgresConfig struct {
	Host         string
	UserName     string
	Password     string
	DatabaseName string
	Port         string
}

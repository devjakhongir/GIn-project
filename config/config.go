package config

type Config struct {
	HTTPPort string

	PostgresHost     string
	PostgresUser     string
	PostgresDatabase string
	PostgresPassword string
	PostgresPort     string
}

func Load() Config {

	var cfg Config

	cfg.HTTPPort = ":4000"

	cfg.PostgresHost = "localhost"
	cfg.PostgresUser = "ravshanbek"
	cfg.PostgresDatabase = "korzinka"
	cfg.PostgresPassword = "8990"
	cfg.PostgresPort = "5432"

	return cfg
}
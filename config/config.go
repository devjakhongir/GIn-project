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

	cfg.HTTPPort = ":5000"

	cfg.PostgresHost = "localhost"
	cfg.PostgresUser = "jahongir2"
	cfg.PostgresDatabase = "population"
	cfg.PostgresPassword = "1005"
	cfg.PostgresPort = "5432"

	return cfg
}

package config

type Config struct {
	PostgresHost     string
	PostgresPort 	 string
	PostgresUser 	 string
	PostgresPassword string
	PostgresDB 		 string
}

func Load() Config{
	cfg := Config{}

	cfg.PostgresHost = "localhost"
	cfg.PostgresPort = "5432"
	cfg.PostgresUser = "the_khoji"
	cfg.PostgresPassword = "546944"
	cfg.PostgresDB = "store_api"

	return cfg
}
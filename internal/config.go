package internal

import "flag"

type Config struct {
	Host        string
	Port        int
	Debug       bool
	PostgresDSN string
}

func ReadConfig() *Config {
	var cfg Config

	flag.StringVar(&cfg.Host, "host", "0.0.0.0", "flag for configure host")
	flag.IntVar(&cfg.Port, "port", 8080, "flag for configure port")
	flag.BoolVar(&cfg.Debug, "debug", false, "enabled debug logger level")
	flag.StringVar(&cfg.PostgresDSN, "dsn", "postgres://user_zero:password@localhost:5435/ToDoRestful?sslmode=disable", "PostgreSQL connection string")

	flag.Parse()

	return &cfg
}

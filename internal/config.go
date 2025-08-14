package internal

import (
	"cmp"
	"flag"
	"os"
	"strconv"
)

type Config struct {
	Host        string
	Port        int
	Debug       bool
	PostgresDSN string
}

const (
	defHost = "0.0.0.0"
	defPort = 8080
	defDB   = "postgres://user:password@localhost:5432/notes?sslmode=disable"
)

func ReadConfig() *Config {
	var cfg Config

	flag.StringVar(&cfg.Host, "host", "0.0.0.0", "flag for configure host")
	flag.IntVar(&cfg.Port, "port", 8080, "flag for configure port")
	flag.BoolVar(&cfg.Debug, "debug", false, "enabled debug logger level")
	flag.StringVar(&cfg.PostgresDSN, "dsn", defDB, "PostgreSQL connection string")

	flag.Parse()

	if cfg.Host == defHost {
		cfg.Host = cmp.Or(os.Getenv("TODO_HOST"), defHost)

	}

	if cfg.Port == defPort {
		port := cmp.Or(os.Getenv("TODO_PORT"), strconv.Itoa(defPort))
		portInt, err := strconv.Atoi(port)
		if err != nil {
			return nil
		}

		cfg.Port = portInt
	}

	if cfg.PostgresDSN == defDB {
		cfg.PostgresDSN = cmp.Or(os.Getenv("TODO_DB"), defDB)

	}

	return &cfg
}

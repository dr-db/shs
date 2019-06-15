package config

import (
	"flag"
	"fmt"
	"log"

	"github.com/pkg/errors"
)

type Config struct {
	Port int
	Root string
}

func MustParseConfig(args []string) *Config {
	cfg, err := parseConfig(args)
	if err != nil {
		log.Fatal(errors.Wrap(err, "loading config"))
	}
	return cfg
}

func parseConfig(args []string) (*Config, error) {
	cfg := &Config{}
	fs := flag.NewFlagSet("", flag.ContinueOnError)

	fs.IntVar(&cfg.Port, "p", 8080, "port to serve")
	fs.StringVar(&cfg.Root, "d", ".", "directory to serve")

	if err := fs.Parse(args); err != nil {
		return nil, errors.Wrap(err, "parsing arguments")
	}
	return cfg, nil
}

func (c Config) HostingAddr() string {
	return fmt.Sprintf(":%d", c.Port)
}

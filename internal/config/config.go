package config

import (
	"flag"
	"fmt"
	"log"
	"strings"

	"github.com/pkg/errors"
)

type Config struct {
	AllowedIPs []string
	Port       int
	Root       string
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
	var rawAllowedIPs string
	fs.StringVar(&rawAllowedIPs, "allowed-ips", "", "IP addresses to allow (comma-separated).  Default allows all.")

	if err := fs.Parse(args); err != nil {
		return nil, errors.Wrap(err, "parsing arguments")
	}

	if rawAllowedIPs != "" {
		cfg.AllowedIPs = strings.Split(rawAllowedIPs, ",")
	}
	return cfg, nil
}

func (c Config) AllowedIP(ip string) bool {
	if len(c.AllowedIPs) == 0 {
		return true
	}
	for _, allowedIP := range c.AllowedIPs {
		if ip == allowedIP {
			return true
		}
	}
	return false
}

func (c Config) HostingAddr() string {
	return fmt.Sprintf(":%d", c.Port)
}

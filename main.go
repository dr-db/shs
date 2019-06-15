package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/dr-db/shs/internal/config"
	"github.com/dr-db/shs/internal/server"
)

func main() {
	cfg := config.MustParseConfig(os.Args[1:])
	fmt.Printf("%#v\n", cfg)

	s := server.NewServer(cfg)
	http.Handle("/", s)

	if cfg.TLS {
		log.Fatal(http.ListenAndServeTLS(cfg.HostingAddr(), cfg.CertFile, cfg.KeyFile, nil))
	}
	log.Fatal(http.ListenAndServe(cfg.HostingAddr(), nil))
}

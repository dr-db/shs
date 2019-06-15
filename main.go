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

	log.Fatal(http.ListenAndServe(cfg.HostingAddr(), nil))
}

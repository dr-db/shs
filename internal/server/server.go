package server

import (
	"log"
	"net/http"
	"strings"

	"github.com/dr-db/shs/internal/config"
)

type Server interface {
	ServeHTTP(w http.ResponseWriter, r *http.Request)
}

type server struct {
	cfg *config.Config
	fs  http.Handler
}

func NewServer(cfg *config.Config) Server {
	return &server{
		cfg: cfg,
		fs:  http.FileServer(http.Dir(cfg.Root)),
	}
}

func (s *server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	addr := strings.SplitN(r.RemoteAddr, ":", 2)[0]
	if !s.cfg.AllowedIP(addr) {
		log.Printf("%q %q %q 403 - IP address not allowed", addr, r.Method, r.URL)
		http.Error(w, "permission denied", http.StatusForbidden)
		return
	}
	log.Printf("%q %q %q", addr, r.Method, r.URL)
	s.fs.ServeHTTP(w, r)
}

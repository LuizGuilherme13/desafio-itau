package server

import (
	"net/http"

	"github.com/LuizGuilherme13/desafio-itau/internal/models"
	"github.com/LuizGuilherme13/desafio-itau/internal/utils/clog"
)

type Server struct {
	Addr  string
	Store models.Storage
}

func New(addr string) *Server {
	return &Server{Addr: addr, Store: models.Storage{}}
}

func (s *Server) Start() {
	routes := s.MountRoutes()

	clog.Info("Server", "running on port "+s.Addr)
	if err := http.ListenAndServe(s.Addr, routes); err != nil {
		panic(err)
	}
}

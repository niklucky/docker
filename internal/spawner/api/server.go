package api

import (
	"fmt"
	"log"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

type Server struct {
	config *Config
	router *httprouter.Router
}

func New(config *Config) *Server {
	if config == nil {
		config = &Config{}
	}
	return &Server{
		config: config,
		router: httprouter.New(),
	}
}

func (s *Server) Start() error {
	s.configureRouter()
	log.Fatal(http.ListenAndServe(s.getAddress(), s.router))
	return nil
}

func (s *Server) getAddress() string {
	if s.config.Address == "" {
		return ":8080"
	}
	return s.config.Address
}

func (s *Server) configureRouter() {
	s.router.GET("/containers/:id", s.handleGetContainers)
}

func (s *Server) handleGetContainers(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	fmt.Fprintf(w, "hello, %s!\n", ps.ByName("id"))
}

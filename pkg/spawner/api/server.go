package api

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/niklucky/docker/pkg/spawner/docker"
)

type Server struct {
	config *Config
	router *httprouter.Router
	docker docker.Controller
	auth   *Auth
}

func New(config *Config, docker docker.Controller) *Server {
	if config == nil {
		config = &Config{}
	}
	return &Server{
		config: config,
		docker: docker,
		router: httprouter.New(),
		auth:   NewAuth(config.SecretKey),
	}
}

func (s *Server) Start() error {
	s.configureRouter()
	log.Println("Starting server on ", s.getAddress())
	log.Fatal(http.ListenAndServe(s.getAddress(), s.router))
	return nil
}

func (s *Server) getAddress() string {
	if s.config.ServerAddress == "" {
		return ":8080"
	}
	return s.config.ServerAddress
}

func (s *Server) configureRouter() {
	s.router.GET("/containers", s.auth.ByToken(s.handleGetContainers))
	s.router.GET("/containers/:id", s.auth.ByToken(s.handleGetContainers))
	s.router.POST("/containers", s.auth.ByToken(s.handleCreateContainer))
	s.router.PUT("/containers/:id", s.auth.ByToken(s.handleGetContainers))
	s.router.DELETE("/containers/:id", s.auth.ByToken(s.handleGetContainers))
	s.router.GET("/system/info", s.auth.ByToken(s.handleGetContainers))
}

func (s *Server) handleGetContainers(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	result, err := s.docker.ContainerList()
	buildResponse(w, result, err)
}
func (s *Server) handleCreateContainer(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	var options docker.CreateContainerOptions
	err := json.NewDecoder(r.Body).Decode(&options)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	result, err := s.docker.ContainerCreate(&options)
	buildResponse(w, result, err)
}

func buildResponse(w http.ResponseWriter, payload interface{}, err error) {
	response := make(map[string]interface{})
	response["data"] = payload
	response["error"] = err

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)

}

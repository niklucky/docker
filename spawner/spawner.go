// package spawner
package spawner

import (
	"github.com/niklucky/docker/pkg/spawner/api"
	"github.com/niklucky/docker/pkg/spawner/docker"
)

// Spawner - struct contains common info for spawning containers
type Spawner struct {
	config *Config
	api    *api.Server
	docker docker.Controller
}

func New(config *Config) *Spawner {
	s := &Spawner{
		config: NewConfig(config),
	}
	s.docker = docker.New(s.config.Docker)
	s.api = api.New(s.config.API, s.docker)
	return s
}

func (s *Spawner) Start() error {
	return s.api.Start()
}
